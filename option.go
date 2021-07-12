package mcts

import "time"

var ZeroDuration time.Duration

type computeOptions struct {
	numberOfGoroutines int
	maxIterations      int
	maxTime            time.Duration
	verbose            bool
}

var defaultComputeOptions = computeOptions{
	numberOfGoroutines: 5,
	maxIterations:      10000,
	maxTime:            ZeroDuration, // default is no time limit.
	verbose:            false,
}

type ComputeOption func(*computeOptions)

func NumberOfGoroutines(numberOfGoroutines int) ComputeOption {
	return func(opts *computeOptions) {
		opts.numberOfGoroutines = numberOfGoroutines
	}
}

func MaxIterations(maxIterations int) ComputeOption {
	return func(opts *computeOptions) {
		opts.maxIterations = maxIterations
	}
}

func MaxTime(maxTime time.Duration) ComputeOption {
	return func(opts *computeOptions) {
		opts.maxTime = maxTime
	}
}

func Verbose(verbose bool) ComputeOption {
	return func(opts *computeOptions) {
		opts.verbose = verbose
	}
}
