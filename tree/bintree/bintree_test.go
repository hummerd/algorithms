package bintree_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/hummerd/algorithms/tree"
	"github.com/hummerd/algorithms/tree/bintree"
	"golang.org/x/exp/constraints"
)

func TestTreeDelete(t *testing.T) {
	n := 144

	bt := &bintree.BinTree[int]{}
	vs := []int{}

	for i := 0; i < n; i++ {
		v := rand.Intn(64)
		vs = append(vs, v)
		bt.Insert(v)
	}

	c := tree.Count[int](bt.Root)
	if c != n {
		t.Fatal("wrong count")
	}

	for i, v := range vs {
		t.Log("delete", v)
		bt.Delete(v)

		r := bt.Root
		if i == n-1 && r != nil {
			t.Fatal("non empty tree after all")
		}

		if i < n-1 && r == nil {
			t.Fatal("unexpected empty tree", i)
		}

		err := checkNode(r)
		if err != nil {
			// rbt.DrawSVG("delerr.svg", tree.Root())
			t.Fatal(err)
		}

		c := tree.Count(bt.RootNode())
		if c != n-i-1 {
			t.Fatal("wrong count")
		}
	}

	// rbt.DrawSVG("del.svg", tree.Root)
}

func checkNode[T constraints.Ordered](n *bintree.BinNode[T]) error {
	if n == nil {
		return nil
	}

	if n.Left != nil && n.Left.Value > n.Value {
		return errors.New("left node greater than parent")
	}

	if n.Right != nil && n.Right.Value < n.Value {
		return errors.New("right node fewer than parent")
	}

	err := checkNode(n.Left)
	if err != nil {
		return err
	}

	err = checkNode(n.Right)
	if err != nil {
		return err
	}

	return nil
}
