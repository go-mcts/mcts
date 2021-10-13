// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

type Move interface{}

type State interface {
	PlayerToMove() int
	HasMoves() bool
	GetMoves() []Move
	DoMove()
	DoRandomMove()
	GetResult() float64
}
