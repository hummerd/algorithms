package dynamic

import (
	"fmt"

	"github.com/hummerd/algorithms/tree/bintree"
)

func RestoreTree(i, j int, restore [][]int) *bintree.BinNode[int] {
	if i > j {
		return nil
	}

	n := restore[i][j]

	l := RestoreTree(i, n-1, restore)
	r := RestoreTree(n+1, j, restore)

	return &bintree.BinNode[int]{
		Value: n,
		Left:  l,
		Right: r,
	}
}

// optimal binary tree
func BuildOBT(p, q []float64) ([][]float64, [][]float64, [][]int) {
	n := len(p)

	e := make([][]float64, n+2) // first index is from zero, second is from 1 (0 reserved for empty subtree, with only q)
	w := make([][]float64, n+2)
	root := make([][]int, n+1)

	for i := 1; i <= n+1; i++ {
		e[i] = make([]float64, n+1)
		w[i] = make([]float64, n+1)

		e[i][i-1] = q[i-1]
		w[i][i-1] = q[i-1]
	}

	for i := 0; i < n+1; i++ {
		root[i] = make([]int, n+1)
	}

	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			min := -1.0
			fmt.Println(l, i, j)
			w[i][j] = w[i][j-1] + p[j-1] + q[j]
			for k := i; k <= j; k++ {
				t := e[i][k-1] + e[k+1][j] + w[i][j]
				fmt.Println(l, i, k, j, "t", t)
				if min == -1.0 || t < min {
					min = t
					root[i][j] = k
				}
			}
			e[i][j] = min
		}
	}

	return e, w, root
}
