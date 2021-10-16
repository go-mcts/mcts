// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"runtime"
	"time"
)

// Options use functional-option to customize mcts
type Options struct {
	Goroutines    int
	MaxIterations int
	MaxTime       time.Duration
	Verbose       bool
}

var defaultOptions = Options{
	Goroutines:    runtime.NumCPU(),
	MaxIterations: 10000,
	MaxTime:       -1,
	Verbose:       false,
}

type Option func(*Options)

// Goroutines number of goroutines, default is runtime.NumCPU()
func Goroutines(number int) Option {
	return func(o *Options) {
		o.Goroutines = number
	}
}

// MaxIterations maximum number of iterations, default is 10000
//
// iter < 0: not limit
func MaxIterations(iter int) Option {
	return func(o *Options) {
		o.MaxIterations = iter
	}
}

// MaxTime search timeout, default is not limit
func MaxTime(d time.Duration) Option {
	return func(o *Options) {
		o.MaxTime = d
	}
}

// Verbose print details log, default is false
func Verbose(v bool) Option {
	return func(o *Options) {
		o.Verbose = v
	}
}

func newOptions(opts ...Option) Options {
	options := defaultOptions

	for _, o := range opts {
		o(&options)
	}

	return options
}
