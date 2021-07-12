package main

import (
	"fmt"
	"github.com/go-mcts/mcts"
)

var _ mcts.State = (*NimState)(nil)

type NimState struct {
	playerJustMoved mcts.Player
	chips           int
}

func (s *NimState) PlayerJustMoved() mcts.Player {
	return s.playerJustMoved
}

func (s *NimState) Clone() mcts.State {
	return &NimState{
		playerJustMoved: s.playerJustMoved,
		chips:           s.chips,
	}
}

func (s *NimState) DoMove(move mcts.Move) {
	m := move.(NimMove)
	if m < 1 || m > 3 {
		panic(fmt.Errorf("illegal move: %v", m))
	}
	s.chips -= int(m)
	s.playerJustMoved = mcts.ChangePlayer(s.playerJustMoved)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (s *NimState) GetMoves() []mcts.Move {
	limit := min(3, s.chips)
	moves := make([]mcts.Move, limit)
	for i := 0; i < limit; i++ {
		moves[i] = NimMove(i + 1)
	}
	return moves
}

func (s *NimState) GetResult(playerJustMoved mcts.Player) mcts.Result {
	if s.chips != 0 {
		panic(fmt.Errorf("illegal chips: %v", s.chips))
	}
	if s.playerJustMoved == playerJustMoved {
		return mcts.ResultWin
	} else {
		return mcts.ResultDraw
	}
}
