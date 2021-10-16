// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

func Test_newOptions(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want Options
	}{
		{
			"default options",
			args{},
			defaultOptions,
		},
		{
			"override some options",
			args{
				opts: []Option{
					MaxIterations(100000),
					Verbose(true),
				},
			},
			Options{
				Groutines:     runtime.NumCPU(),
				MaxIterations: 100000,
				MaxTime:       -1,
				Verbose:       true,
			},
		},
		{
			"override all options",
			args{
				opts: []Option{
					Goroutines(1),
					MaxIterations(100000),
					MaxTime(5 * time.Second),
					Verbose(true),
				},
			},
			Options{
				Groutines:     1,
				MaxIterations: 100000,
				MaxTime:       5 * time.Second,
				Verbose:       true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
