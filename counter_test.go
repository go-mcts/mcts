// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type someStructPointer struct {
	s string
}

func newPointer(s string) *someStructPointer {
	return &someStructPointer{s}
}

func TestCounter(t *testing.T) {
	m := make(map[*someStructPointer]int)
	m[newPointer("abc")]++
	m[newPointer("abc")]++
	assert.Equal(t, 2, len(m))
	assert.Equal(t, 0, m[newPointer("abc")])

	c := newCounter()
	c.incr(newPointer("abc"), 1)
	c.incr(newPointer("abc"), 1)
	assert.Equal(t, float64(2), c.get(newPointer("abc")))
}
