// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"math/rand"
	"runtime"
	"time"
)

type Options struct {
	NumberOfGoroutines int
	MaxIterations      int
	MaxTime            time.Duration

	// TODO: add verbose log
}

type Option func(*Options)

func NumberOfGoroutines(number int) Option {
	return func(o *Options) {
		o.NumberOfGoroutines = number
	}
}

func MaxIterations(iter int) Option {
	return func(o *Options) {
		o.MaxIterations = iter
	}
}

func MaxTime(d time.Duration) Option {
	return func(o *Options) {
		o.MaxTime = d
	}
}

func newOptions(opts ...Option) Options {
	options := Options{
		NumberOfGoroutines: runtime.NumCPU(),
		MaxIterations:      10000,
		MaxTime:            -1,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func ComputeTree(rootState State, rd *rand.Rand, opts ...Option) *Node {
	options := newOptions(opts...)

	if options.MaxIterations < 0 && options.MaxTime < 0 {
		panic("illegal options")
	}

	if rootState.PlayerToMove() != 1 && rootState.PlayerToMove() != 2 {
		panic("only support player1 and player2")
	}

	var deadline time.Time
	if options.MaxTime >= 0 {
		deadline = time.Now().Add(options.MaxTime)
	}

	root := NewNode(rootState)
	for i := 1; i <= options.MaxIterations || options.MaxIterations < 0; i++ {
		node := root
		state := rootState.Clone()

		for !node.HasUntriedMoves() && node.HasChildren() {
			node = node.SelectChildUCT()
			state.DoMove(node.Move)
		}

		if node.HasUntriedMoves() {
			move := node.GetUntriedMove(rd)
			state.DoMove(move)
			node = node.AddChild(move, state)
		}

		for state.HasMoves() {
			state.DoRandomMove(rd)
		}

		for node != nil {
			node.Update(state.GetResult(node.PlayerToMove))
			node = node.Parent
		}

		if !deadline.IsZero() && time.Now().Before(deadline) {
			break
		}
	}

	return root
}

func ComputeMove(rootState State, opts ...Option) Move {
	options := newOptions(opts...)

	if rootState.PlayerToMove() != 1 && rootState.PlayerToMove() != 2 {
		panic("only support player1 and player2")
	}

	moves := rootState.GetMoves()
	if len(moves) == 0 {
		panic("root moves is empty")
	}

	if len(moves) == 1 {
		return moves[0]
	}

	rootFutures := make(chan *Node, options.NumberOfGoroutines)
	for i := 0; i < options.NumberOfGoroutines; i++ {
		go func() {
			rd := rand.New(rand.NewSource(time.Now().UnixNano()))
			rootFutures <- ComputeTree(rootState, rd, opts...)
		}()
	}

	visits := make(map[Move]int)
	wins := make(map[Move]float64)
	gamePlayed := 0
	for i := 0; i < options.NumberOfGoroutines; i++ {
		root := <-rootFutures
		gamePlayed += root.Visits
		for _, c := range root.Children {
			visits[c.Move] += c.Visits
			wins[c.Move] += c.Wins
		}
	}

	bestScore := float64(-1)
	var bestMove Move
	for move, v := range visits {
		w := wins[move]
		expectedSuccessRate := (w + 1) / (float64(v) + 2)
		if expectedSuccessRate > bestScore {
			bestMove = move
			bestScore = expectedSuccessRate
		}
	}
	return bestMove
}
