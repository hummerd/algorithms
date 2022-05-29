package sort

import (
	"math/rand"
	"testing"
)

func TestHeapProperties(t *testing.T) {
	h := MaxHeap[int]{}
	for i := 0; i < 24; i++ {
		h.Add(rand.Intn(10))
	}

	// check heap properties
	n := len(h.a) / 2
	for i := n; i >= 0; i-- {
		l := left(i)
		r := right(i)

		if l < len(h.a) && h.a[i] < h.a[l] {
			t.Log(h.a)
			t.Fatalf("left %d is bigger then parent [%d]=%d", h.a[l], i, h.a[i])
		}

		if r < len(h.a) && h.a[i] < h.a[r] {
			t.Log(h.a)
			t.Fatalf("right %d is bigger then parent [%d]=%d", h.a[r], i, h.a[i])
		}
	}
}
