// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxElement(t *testing.T) {
	nodes := []*Node{
		{Visits: 3, uctScore: 6.2},
		{Visits: 5, uctScore: 3.5},
		{Visits: 1, uctScore: 5.12},
		{Visits: 4, uctScore: 7.1},
		{Visits: 2, uctScore: 1.23},
	}

	e := maxElement(nodes, func(i, j int) bool {
		return nodes[i].Visits > nodes[j].Visits
	}).(*Node)

	assert.Equal(t, nodes[1], e)

	e = maxElement(nodes, func(i, j int) bool {
		return nodes[i].uctScore > nodes[j].uctScore
	}).(*Node)

	assert.Equal(t, nodes[3], e)
}
