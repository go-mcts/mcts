// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nim

import (
	"math/rand"

	"github.com/go-mcts/mcts"
)

var (
	_ mcts.Move  = (*Move)(nil)
	_ mcts.State = (*State)(nil)
)

type Move int

type State struct {
	playerToMove int
	chips        int
}

func (s *State) PlayerToMove() int {
	return s.playerToMove
}

func (s *State) HasMoves() bool {
	s.checkInvariant()
	return s.chips > 0
}

func (s *State) GetMoves() []mcts.Move {
	s.checkInvariant()

	var moves []mcts.Move
	for i := 1; i <= min(3, s.chips); i++ {
		moves = append(moves, i)
	}
	return moves
}

func (s *State) DoMove(move mcts.Move) {
	m := move.(int)
	if m < 1 || m > 3 {
		panic("illegal move")
	}
	s.checkInvariant()

	s.chips -= m
	s.playerToMove = 3 - s.playerToMove

	s.checkInvariant()
}

func (s *State) DoRandomMove(rd *rand.Rand) {
	if s.chips <= 0 {
		panic("invalid chips")
	}
	s.checkInvariant()

	max := min(3, s.chips)
	s.DoMove(rd.Intn(max) + 1)

	s.checkInvariant()
}

func (s *State) GetResult(currentPlayer int) float64 {
	if s.chips != 0 {
		panic("game not over")
	}
	s.checkInvariant()

	if s.playerToMove == currentPlayer {
		return 1.0
	}
	return 0.0
}

func (s *State) Clone() mcts.State {
	return &State{
		playerToMove: s.playerToMove,
		chips:        s.chips,
	}
}

func (s *State) checkInvariant() {
	if s.chips < 0 || (s.playerToMove != 1 && s.playerToMove != 2) {
		panic("illegal state")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
