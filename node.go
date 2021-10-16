// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"math"
	"math/rand"
)

type node struct {
	move         Move
	parent       *node
	playerToMove int
	wins         float64
	visits       int
	moves        []Move
	children     []*node
}

func newNode(state State, move Move, parent *node) *node {
	return &node{
		move:         move,
		parent:       parent,
		playerToMove: state.PlayerToMove(),
		wins:         0,
		visits:       0,
		moves:        state.GetMoves(),
	}
}

func (p *node) hasUntriedMoves() bool {
	return len(p.moves) > 0
}

func (p *node) hasChildren() bool {
	return len(p.children) > 0
}

func (p *node) getUntriedMove(rd *rand.Rand) Move {
	l := len(p.moves)
	if l == 0 {
		panic("untried moves is empty")
	}
	return p.moves[rd.Intn(l)]
}

func (p *node) selectChildUCT() *node {
	l := len(p.children)
	if l == 0 {
		panic("children is empty")
	}

	best := p.children[0]
	bestScore := best.wins/float64(best.visits) +
		math.Sqrt(2.0*math.Log(float64(p.visits))/float64(best.visits))

	for i := 1; i < l; i++ {
		c := p.children[i]
		uctScore := c.wins/float64(c.visits) +
			math.Sqrt(2.0*math.Log(float64(p.visits))/float64(c.visits))

		if uctScore > bestScore {
			bestScore = uctScore
			best = c
		}
	}
	return best
}

func (p *node) addChild(move Move, state State) *node {
	node := newNode(state, move, p)
	p.children = append(p.children, node)

	l := len(p.moves)
	for i := 0; i < l-1; i++ {
		if p.moves[i] == move {
			p.moves[i] = p.moves[l-1]
			break
		}
	}
	p.moves = p.moves[:l-1]
	return node
}

func (p *node) update(result float64) {
	p.visits++
	p.wins += result
}
