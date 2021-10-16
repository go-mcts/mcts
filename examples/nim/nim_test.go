// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package nim

import (
	"testing"

	"github.com/go-mcts/mcts"
	"github.com/stretchr/testify/assert"
)

func TestNim(t *testing.T) {
	for chips := 4; chips <= 21; chips++ {
		if chips%4 != 0 {
			state := &State{
				playerToMove: 1,
				chips:        chips,
			}
			move := mcts.ComputeMove(state, mcts.MaxIterations(100000), mcts.Verbose(true))
			assert.Equal(t, chips%4, move)
		}
	}
}
