package main

import "github.com/go-mcts/mcts"

var _ mcts.Move = (*Move)(nil)

type Move struct {
	x int
	y int
	v int
}
