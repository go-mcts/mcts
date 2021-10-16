// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ Move  = (*testGameMove)(nil)
	_ State = (*testGameState)(nil)
)

// Player 1 has two options:
//		1: Draw.
//		2: Nothing happens (Player 2's turn).
//
// Player 2 has five options:
//		  1: Player 1 wins.
//		2-5: Player X wins. (default: 2)
//
// If X == 1, player 1 should play 2 for a guaranteed win after the next move by player 2.
// If X == 2, player 1 should play 1 for an immediate draw.
type testGameState struct {
	playerToMove int
	winner       int
	x            int
}

func newTestGameState(x int) State {
	return &testGameState{
		playerToMove: 1,
		winner:       -1,
		x:            x,
	}
}

func (s *testGameState) PlayerToMove() int {
	return s.playerToMove
}

func (s *testGameState) HasMoves() bool {
	return s.winner < 0
}

func (s *testGameState) GetMoves() []Move {
	var moves []Move
	if !s.HasMoves() {
		return moves
	}

	if s.playerToMove == 1 {
		moves = append(moves, 1, 2)
	} else if s.playerToMove == 2 {
		moves = append(moves, 1, 2, 3, 4, 5)
	}
	return moves
}

func (s *testGameState) DoMove(move Move) {
	m := move.(int)
	if s.playerToMove == 1 {
		if m != 1 && m != 2 {
			panic("illegal move")
		}

		if move == 1 {
			s.winner = 0
		}
	} else if s.playerToMove == 2 {
		if m < 1 || m > 5 {
			panic("illegal move")
		}

		if move == 1 {
			s.winner = 1
		} else {
			s.winner = s.x
		}
	}

	s.playerToMove = 3 - s.playerToMove
}

func (s *testGameState) DoRandomMove(rd *rand.Rand) {
	if s.playerToMove == 1 {
		s.DoMove(rd.Intn(2) + 1)
	} else if s.playerToMove == 2 {
		s.DoMove(rd.Intn(5) + 1)
	}
}

func (s *testGameState) GetResult(currentPlayer int) float64 {
	if s.winner < 0 {
		panic("game not over")
	}

	if s.winner == 0 {
		return 0.5
	}

	if s.winner == currentPlayer {
		return 0.0
	}
	return 1.0
}

func (s *testGameState) Clone() State {
	return &testGameState{
		playerToMove: s.playerToMove,
		winner:       s.winner,
		x:            s.x,
	}
}

type testGameMove int

func TestMCTS(t *testing.T) {
	state := newTestGameState(1)
	move := ComputeMove(state, Verbose(true))
	assert.Equal(t, 2, move)

	state = newTestGameState(2)
	move = ComputeMove(state, Verbose(true))
	assert.Equal(t, 1, move)
}
