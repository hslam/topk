// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

package topk

import (
	"math/rand"
	"testing"
)

type list []int

func (h list) Len() int           { return len(h) }
func (h list) Less(i, j int) bool { return h[i] < h[j] }
func (h list) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func TestTopK(t *testing.T) {
	testTopK(3, false, t)
	testTopK(3, true, t)
	testTopK(9, false, t)
}

func testTopK(k int, sortK bool, t *testing.T) {
	l := list{10, 7, 2, 5, 1}
	Top(l, k, sortK)
	n := l.Len()
	if k > n {
		k = n
	}
	for j := k; j < l.Len(); j++ {
		for i := 0; i < k-1; i++ {
			if l.Less(i, j) {
				t.Error("top error")
			}
		}
	}
	if sortK {
		for i := 0; i < k-1; i++ {
			if l.Less(i, i+1) {
				t.Error("out of order")
			}
		}
	}
	t.Log(l)
}

func TestHeapInsert(t *testing.T) {
	l := list{10, 7, 2, 5, 1}
	n := l.Len()
	for i := 1; i < n; i++ {
		up(l, i)
	}
	for i := 1; i < n-1; i++ {
		if l.Less(1, 0) {
			t.Error("HeapInsert error")
		}
	}
	t.Log(l)
}

func TestHeapify(t *testing.T) {
	l := list{10, 7, 2, 5, 1}
	n := l.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(l, i, n)
	}
	for i := 1; i < n-1; i++ {
		if l.Less(i, 0) {
			t.Error("Heapify error")
		}
	}
	t.Log(l)
}

func BenchmarkHeapInsert(b *testing.B) {
	var l list
	num := 1000
	for i := 0; i < num; i++ {
		l = append(l, rand.Intn(1000000))
	}
	cp := make(list, num)
	for i := 0; i < b.N; i++ {
		copy(cp, l)
		n := cp.Len()
		for j := 1; j < n; j++ {
			up(cp, j)
		}
	}
}

func BenchmarkHeapify(b *testing.B) {
	var l list
	num := 1000
	for i := 0; i < num; i++ {
		l = append(l, rand.Intn(1000000))
	}
	cp := make(list, num)
	for i := 0; i < b.N; i++ {
		copy(cp, l)
		n := cp.Len()
		for j := n/2 - 1; j >= 0; j-- {
			down(cp, j, n)
		}
	}
}

func BenchmarkTopk(b *testing.B) {
	var l list
	num := 1000
	for i := 0; i < num; i++ {
		l = append(l, rand.Intn(1000000))
	}
	cp := make(list, num)
	k := 10
	for i := 0; i < b.N; i++ {
		copy(cp, l)
		Top(cp, k, false)
	}
}
