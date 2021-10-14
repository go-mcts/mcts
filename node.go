// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"math"
	"math/rand"
	"reflect"
)

type Node struct {
	Move         Move
	Parent       *Node
	PlayerToMove int
	Wins         float64
	Visits       int
	Moves        []Move
	Children     []*Node
	uctScore     float64
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
		uctScore:     0,
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
	if len(p.Children) == 0 {
		panic("children is empty")
	}

	return maxElement(p.Children, func(i, j int) bool {
		return p.Children[i].Visits > p.Children[j].Visits
	}).(*Node)
}

func (p *Node) SelectChildUCT() *Node {
	if len(p.Children) == 0 {
		panic("children is empty")
	}
	for _, c := range p.Children {
		c.uctScore = c.Wins/float64(c.Visits) +
			math.Sqrt(2.0*math.Log(float64(p.Visits))/float64(c.Visits))
	}

	return maxElement(p.Children, func(i, j int) bool {
		return p.Children[i].uctScore > p.Children[j].uctScore
	}).(*Node)
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

func maxElement(a interface{}, greater func(i, j int) bool) interface{} {
	rv := reflect.ValueOf(a)
	l := rv.Len()
	if l == 0 {
		return nil
	}
	j := 0
	for i := 1; i < l; i++ {
		if greater(i, j) {
			j = i
		}
	}
	return rv.Index(j).Interface()
}
