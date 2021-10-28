// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tictactoe

import (
	"testing"

	"github.com/go-mcts/mcts"
	"github.com/stretchr/testify/assert"
)

func TestTicTacToe(t *testing.T) {
	rootState := &state{
		playerToMove: 1,
		board: [3][3]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}
	mctsMove := mcts.ComputeMove(rootState, mcts.MaxIterations(20000), mcts.Verbose(true))
	m := mctsMove.(move)
	assert.Equal(t, 1, m.x)
	assert.Equal(t, 1, m.y)
	assert.Equal(t, 1, m.v)

	rootState = &state{
		playerToMove: 1,
		board: [3][3]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, -1, 0},
		},
	}
	mctsMove = mcts.ComputeMove(rootState, mcts.Verbose(true))
	m = mctsMove.(move)
	assert.Equal(t, 1, m.v)

	assert.True(t, m.x == 0 && (m.y == 0 || m.y == 2) ||
		m.x == 2 && (m.y == 0 || m.y == 2))
}
