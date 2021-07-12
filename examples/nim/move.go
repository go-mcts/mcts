package main

import "github.com/go-mcts/mcts"

var _ mcts.Move = (*NimMove)(nil)

type NimMove int
