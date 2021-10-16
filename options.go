// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"runtime"
	"time"
)

type Options struct {
	Groutines     int
	MaxIterations int
	MaxTime       time.Duration
	Verbose       bool
}

type Option func(*Options)

func Goroutines(number int) Option {
	return func(o *Options) {
		o.Groutines = number
	}
}

func MaxIterations(iter int) Option {
	return func(o *Options) {
		o.MaxIterations = iter
	}
}

func MaxTime(d time.Duration) Option {
	return func(o *Options) {
		o.MaxTime = d
	}
}

func Verbose(v bool) Option {
	return func(o *Options) {
		o.Verbose = v
	}
}

func newOptions(opts ...Option) Options {
	options := Options{
		Groutines:     runtime.NumCPU(),
		MaxIterations: 10000,
		MaxTime:       -1,
		Verbose:       false,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}
