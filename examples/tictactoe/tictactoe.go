// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tictactoe

import (
	"math/rand"

	"github.com/go-mcts/mcts"
)

var (
	_ mcts.Move  = (*move)(nil)
	_ mcts.State = (*state)(nil)
)

type move struct {
	x int
	y int
	v int
}

type state struct {
	playerToMove int
	board        [3][3]int
}

func (s *state) PlayerToMove() int {
	return s.playerToMove
}

func (s *state) HasMoves() bool {
	return s.getResult(s.playerToMove) == -1
}

func (s *state) GetMoves() []mcts.Move {
	moves := make([]mcts.Move, 0)
	if s.getResult(s.playerToMove) == -1 {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if s.board[i][j] == 0 {
					m := &move{
						x: i,
						y: j,
						v: s.playerToMove,
					}
					if s.playerToMove == 1 {
						m.v = 1
					} else {
						m.v = -1
					}
					moves = append(moves, m)
				}
			}
		}
	}
	return moves
}

func (s *state) DoMove(mctsMove mcts.Move) {
	m := mctsMove.(*move)
	if m.x < 0 || m.y < 0 || m.x > 2 || m.y > 2 || s.board[m.x][m.y] != 0 {
		panic("illegal move")
	}
	s.board[m.x][m.y] = m.v
	s.playerToMove = 3 - s.playerToMove
}

func (s *state) DoRandomMove(rd *rand.Rand) {
	moves := s.GetMoves()
	s.DoMove(moves[rd.Intn(len(moves))])
}

func (s *state) GetResult(currentPlayerToMove int) float64 {
	if result := s.getResult(currentPlayerToMove); result == -1 {
		panic("game is not over")
	} else {
		return result
	}
}

func (s *state) getResult(currentPlayerToMove int) float64 {
	zero := 0

	for i := 0; i < 3; i++ {
		row, col := 0, 0
		for j := 0; j < 3; j++ {
			if s.board[i][j] == 0 {
				zero++
			}
			row += s.board[i][j]
			col += s.board[j][i]
		}

		if row == 3 || row == -3 || col == 3 || col == -3 {
			if s.playerToMove == currentPlayerToMove {
				return 1
			}
			return 0
		}
	}

	tl := s.board[0][0] + s.board[1][1] + s.board[2][2]
	tr := s.board[0][2] + s.board[1][1] + s.board[2][0]

	if tl == 3 || tr == 3 || tl == -3 || tr == -3 {
		if s.playerToMove == currentPlayerToMove {
			return 1
		}
		return 0
	}

	if zero == 0 {
		return 0.5
	}

	return -1
}

func (s *state) Clone() mcts.State {
	return &state{
		playerToMove: s.playerToMove,
		board:        s.board,
	}
}
