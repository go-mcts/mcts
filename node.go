// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"math"
	"math/rand"
)

type Node struct {
	Move         Move
	Parent       *Node
	PlayerToMove int
	Wins         float64
	Visits       int
	Moves        []Move
	Children     []*Node
}

func NewNode(state State) *Node {
	return newNode(state, nil, nil)
}

func newNode(state State, move Move, parent *Node) *Node {
	return &Node{
		Move:         move,
		Parent:       parent,
		PlayerToMove: state.PlayerToMove(),
		Wins:         0,
		Visits:       0,
		Moves:        state.GetMoves(),
	}
}

func (p *Node) HasUntriedMoves() bool {
	return len(p.Moves) > 0
}

func (p *Node) HasChildren() bool {
	return len(p.Children) > 0
}

func (p *Node) GetUntriedMove(rd *rand.Rand) Move {
	l := len(p.Moves)
	if l == 0 {
		panic("untried moves is empty")
	}
	return p.Moves[rd.Intn(l)]
}

func (p *Node) BestChild() *Node {
	if len(p.Moves) > 0 {
		panic("not full expanded")
	}
	l := len(p.Children)
	if l == 0 {
		panic("children is empty")
	}

	best := p.Children[0]
	for i := 1; i < l; i++ {
		if p.Children[i].Visits > best.Visits {
			best = p.Children[i]
		}
	}
	return best
}

func (p *Node) SelectChildUCT() *Node {
	l := len(p.Children)
	if l == 0 {
		panic("children is empty")
	}

	best := p.Children[0]
	bestScore := best.Wins/float64(best.Visits) +
		math.Sqrt(2.0*math.Log(float64(p.Visits))/float64(best.Visits))

	for i := 1; i < l; i++ {
		c := p.Children[i]
		uctScore := c.Wins/float64(c.Visits) +
			math.Sqrt(2.0*math.Log(float64(p.Visits))/float64(c.Visits))

		if uctScore > bestScore {
			bestScore = uctScore
			best = c
		}
	}
	return best
}

func (p *Node) AddChild(move Move, state State) *Node {
	node := newNode(state, move, p)
	p.Children = append(p.Children, node)

	l := len(p.Moves)
	for i := 0; i < l-1; i++ {
		if p.Moves[i] == move {
			p.Moves[i] = p.Moves[l-1]
			break
		}
	}
	p.Moves = p.Moves[:l-1]
	return node
}

func (p *Node) Update(result float64) {
	p.Visits++
	p.Wins += result
}
