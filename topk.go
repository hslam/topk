// Copyright (c) 2020 Meng Huang (mhboy@outlook.com)
// This package is licensed under a MIT license that can be found in the LICENSE file.

// Package topk finds the top k elements in the collection.
package topk

// Interface is a type, typically a collection, that satisfies topk.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Top finds the top k elements in the collection.
// When the sortK is set,the top k elements will be sorted in descending order.
func Top(h Interface, k int, sortK bool) {
	n := h.Len()
	if k > n {
		k = n
	}
	for i := k/2 - 1; i >= 0; i-- {
		down(h, i, k)
	}
	if k < n {
		for i := k; i < n; i++ {
			if h.Less(0, i) {
				h.Swap(0, i)
				down(h, 0, k)
			}
		}
	}
	if sortK {
		for i := k; i > 0; i-- {
			h.Swap(0, i-1)
			down(h, 0, i-1)
		}
	}
}

func up(h Interface, j int) {
	child := j
	for {
		parent := (child - 1) / 2
		if parent == child || !h.Less(child, parent) {
			break
		}
		h.Swap(parent, child)
		child = parent
	}
}

func down(h Interface, i, n int) bool {
	parent := i
	for {
		leftChild := 2*parent + 1
		if leftChild >= n || leftChild < 0 { // leftChild < 0 after int overflow
			break
		}
		lessChild := leftChild
		if rightChild := leftChild + 1; rightChild < n && h.Less(rightChild, leftChild) {
			lessChild = rightChild
		}
		if !h.Less(lessChild, parent) {
			break
		}
		h.Swap(parent, lessChild)
		parent = lessChild
	}
	return parent > i
}
