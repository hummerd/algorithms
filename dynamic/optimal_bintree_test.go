package dynamic_test

import (
	"testing"

	"github.com/hummerd/algorithms/dynamic"
	"github.com/hummerd/algorithms/tree"
	"github.com/hummerd/algorithms/tree/bintree"
)

func TestBuildTree(t *testing.T) {
	_, _, r := dynamic.BuildOBT(
		[]float64{0.15, 0.1, 0.05, 0.1, 0.2},
		[]float64{0.05, 0.1, 0.05, 0.05, 0.05, 0.1})

	expected := &bintree.BinNode[int]{
		Value: 2,
		Left: &bintree.BinNode[int]{
			Value: 1,
		},
		Right: &bintree.BinNode[int]{
			Value: 5,
			Left: &bintree.BinNode[int]{
				Value: 4,
				Left: &bintree.BinNode[int]{
					Value: 3,
				},
			},
		},
	}

	root := dynamic.RestoreTree(1, 5, r)
	eq := tree.Equal[int](expected, root)
	if !eq {
		t.Fail()
	}
	// DrawSVGFile("tree.svg", root)
}
