// Copyright 2021 go-mcts. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mcts

import (
	"fmt"
	"reflect"
)

// Struct pointers as map keys is not work correctly.
// see https://abhinavg.net/posts/pointers-as-map-keys/
//
// use fmt.Sprintf("%v", key) as map keys
type counter struct {
	m map[interface{}]*entry
}

type entry struct {
	key   interface{}
	count float64
}

func newCounter() *counter {
	return &counter{make(map[interface{}]*entry)}
}

func (c *counter) incr(key interface{}, count float64) {
	key = interface2key(key)
	if ent, ok := c.m[key]; ok {
		ent.count += count
	} else {
		c.m[key] = &entry{key, 1}
	}
}

func (c *counter) get(key interface{}) float64 {
	if ent, ok := c.m[interface2key(key)]; ok {
		return ent.count
	}
	return 0
}

func (c *counter) rng(f func(key interface{}, count float64)) {
	for _, ent := range c.m {
		f(ent.key, ent.count)
	}
}

func interface2key(i interface{}) interface{} {
	rt := reflect.TypeOf(i)
	if rt.Kind() == reflect.Ptr {
		return fmt.Sprintf("%#v", i)
	}
	return i
}
