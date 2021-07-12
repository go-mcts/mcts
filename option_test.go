package mcts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroDuration(t *testing.T) {
	assert.Equal(t, ZeroDuration, defaultComputeOptions.maxTime)
}
