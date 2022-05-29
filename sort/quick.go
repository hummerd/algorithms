// Quick sort - chapter 7.
package sort

import (
	"golang.org/x/exp/constraints"
)

func Quick[T constraints.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}

	ni := len(a) - 1
	n := a[ni]

	i, j := -1, 0
	for ; j < len(a)-1; j++ {
		if a[j] <= n {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[ni], a[i+1] = a[i+1], a[ni]

	Quick(a[:i+1])
	Quick(a[i+1:])
}

func Quick2[T constraints.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}

	ni := len(a) / 2
	n := a[ni]

	i, j := 0, len(a)-1
	for i < j {
		if a[i] <= n {
			i++
			continue
		}

		if a[j] > n {
			j--
			continue
		}

		a[i], a[j] = a[j], a[i]
	}

	// everything that is <i is less or equal to n
	// everything that is >i is greater then n
	// we are going to spit task into two subtasks [0, i) and [i, len(a)-1]
	// let's check that a[i] will go to correct subtask
	if a[i] < n {
		a[ni], a[i] = a[i], a[ni]
	}

	Quick(a[:i])
	Quick(a[i:])
}
