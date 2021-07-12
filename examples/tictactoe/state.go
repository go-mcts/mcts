package main

import (
	"github.com/go-mcts/mcts"
)

var _ mcts.State = (*State)(nil)

const (
	ChessX = 1
	ChessO = -1
)

type State struct {
	playerJustMoved mcts.Player
	board           [3][3]int
}

func (s *State) PlayerJustMoved() mcts.Player {
	return s.playerJustMoved
}

func (s *State) Clone() mcts.State {
	return &State{
		playerJustMoved: s.playerJustMoved,
		board:           s.board,
	}
}

func (s *State) DoMove(move mcts.Move) {
	m := move.(*Move)
	if m.x < 0 || m.y < 0 || m.x > 2 || m.y > 2 || s.board[m.x][m.y] != 0 {
		panic("illegal move")
	}
	s.board[m.x][m.y] = m.v
	s.playerJustMoved = mcts.ChangePlayer(s.playerJustMoved)
}

func (s *State) GetMoves() []mcts.Move {
	var v int
	if s.playerJustMoved == mcts.Player1 {
		v = ChessO
	} else {
		v = ChessX
	}
	moves := make([]mcts.Move, 0)
	if s.getResult(s.playerJustMoved) == mcts.ResultNone {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if s.board[i][j] == 0 {
					moves = append(moves, &Move{
						x: i,
						y: j,
						v: v,
					})
				}
			}
		}
	}
	return moves
}

// getResult without panic
func (s *State) getResult(playerJustMoved mcts.Player) mcts.Result {
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

		if row == 3 || col == 3 {
			if playerJustMoved == mcts.Player1 {
				return mcts.ResultWin
			} else {
				return mcts.ResultLose
			}
		}

		if row == -3 || col == -3 {
			if playerJustMoved == mcts.Player2 {
				return mcts.ResultWin
			} else {
				return mcts.ResultLose
			}
		}
	}

	tl := s.board[0][0] + s.board[1][1] + s.board[2][2]
	tr := s.board[0][2] + s.board[1][1] + s.board[2][0]

	if tl == 3 || tr == 3 {
		if playerJustMoved == mcts.Player1 {
			return mcts.ResultWin
		} else {
			return mcts.ResultLose
		}
	}

	if tl == -3 || tr == -3 {
		if playerJustMoved == mcts.Player2 {
			return mcts.ResultWin
		} else {
			return mcts.ResultLose
		}
	}

	if zero == 0 {
		return mcts.ResultDraw
	}

	return mcts.ResultNone
}

func (s *State) GetResult(playerJustMoved mcts.Player) mcts.Result {
	if r := s.getResult(playerJustMoved); r != mcts.ResultNone {
		return r
	}
	panic("game is not over")
}
