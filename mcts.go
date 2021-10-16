// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"math/rand"
)

// Move must be implemented for different games
type Move interface{}

// State must be implemented for different games
type State interface {
	// PlayerToMove is who next to play
	PlayerToMove() int
	// HasMoves return whether the game is over
	HasMoves() bool
	// GetMoves get all legal moves
	GetMoves() []Move
	// DoMove modify state with the given move
	DoMove(move Move)
	// DoRandomMove do random move with the given random engine
	DoRandomMove(rd *rand.Rand)
	// GetResult return game result
	GetResult(currentPlayerToMove int) float64
	// Clone is deep copy
	Clone() State
}
