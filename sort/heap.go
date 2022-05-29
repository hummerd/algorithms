// Heap sort - chapter 6.
package sort

import (
	"golang.org/x/exp/constraints"
)

func BuildMaxHeap[T constraints.Ordered](a []T) *MaxHeap[T] {
	buildMaxHeap(a)
	return &MaxHeap[T]{a}
}

func buildMaxHeap[T constraints.Ordered](a []T) {
	n := len(a) / 2
	for i := n; i >= 0; i-- {
		maxHeapify(a, i)
	}
}

type MaxHeap[T constraints.Ordered] struct {
	a []T
}

func (h *MaxHeap[T]) Add(i T) {
	h.a = append(h.a, i)
	maxIncreaseKey(h.a, len(h.a)-1)
}

func maxIncreaseKey[T constraints.Ordered](a []T, i int) {
	p := parent(i)
	if a[i] > a[p] {
		a[p], a[i] = a[i], a[p]
		maxIncreaseKey(a, p)
	}
}

func maxHeapify[T constraints.Ordered](a []T, i int) {
	l := left(i)
	r := right(i)

	m := i
	if l < len(a) && a[l] > a[m] {
		m = l
	}
	if r < len(a) && a[r] > a[m] {
		m = r
	}

	if m != i {
		a[m], a[i] = a[i], a[m]
		maxHeapify(a, m)
	}
}

// returns left node for zero based indexing
func left(i int) int {
	return i*2 + 1
}

// returns right node for zero based indexing
func right(i int) int {
	return i*2 + 2
}

// returns parent node for zero based indexing
func parent(i int) int {
	return (i - 1) / 2
}

// Heap a using heap sort algorithm.
// complexity: O(n*log(n))
func Heap[T constraints.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}

	buildMaxHeap(a)
	for i := len(a) - 1; i > 0; i-- {
		a[0], a[i] = a[i], a[0]
		maxHeapify(a[:i], 0)
	}
}
