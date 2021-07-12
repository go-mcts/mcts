package main

import (
	"fmt"
	"github.com/go-mcts/mcts"
)

var (
	_ mcts.Move  = (*NimMove)(nil)
	_ mcts.State = (*NimState)(nil)
)

type NimMove int

type NimState struct {
	playerJustMoved int
	chips           int
}

func (s *NimState) PlayerJustMoved() int {
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
	s.playerJustMoved = 3 - s.playerJustMoved
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

func (s *NimState) GetResult(playerJustMoved int) float64 {
	if s.chips != 0 {
		panic(fmt.Errorf("illegal chips: %v", s.chips))
	}
	if s.playerJustMoved == playerJustMoved {
		return 1.0
	} else {
		return 0.0
	}
}
