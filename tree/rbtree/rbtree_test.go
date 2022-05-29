package rbtree

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestTreeHeight(t *testing.T) {
	n := 144

	tree := &RBTree[int]{}

	for i := 0; i < n; i++ {
		v := rand.Intn(64)
		t.Log("insert", v)
		tree.Insert(v)

		err := checkTree(tree.Root)
		if err != nil {
			t.Fatal(err)
		}
	}

	// rbt.DrawSVG("test.svg", tree.Root())
	h := tree.Height()

	if float64(h) > 2*math.Log2(float64(n+1)) {
		t.Fatal("height is too big: ", h)
	}
}

func TestTreeDelete(t *testing.T) {
	n := 144

	tree := &RBTree[int]{}
	vs := []int{}

	for i := 0; i < n; i++ {
		v := rand.Intn(64)
		vs = append(vs, v)
		tree.Insert(v)
	}

	for i, v := range vs {
		t.Log("delete", v)
		tree.Delete(v)

		r := tree.Root
		if i == n-1 && r != nil {
			t.Fatal("non empty tree after all")
		}

		if i < n-1 && r == nil {
			t.Fatal("unexpected empty tree", i)
		}

		err := checkTree(r)
		if err != nil {
			// rbt.DrawSVG("delerr.svg", tree.Root())
			t.Fatal(err)
		}
	}

	// rbt.DrawSVG("del.svg", tree.Root)
}

func BenchmarkTreeInsert(b *testing.B) {
	tree := &RBTree[int]{}

	for i := 0; i < b.N; i++ {
		// v := rand.Intn(64)
		tree.Insert(7)
	}
}

func checkTree[T constraints.Ordered](n *RBNode[T]) error {
	if n == nil {
		return nil
	}

	if n.Red {
		return errors.New("root not black")
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

func checkNode[T constraints.Ordered](n *RBNode[T]) error {
	if n == nil {
		return nil
	}

	if n.Parent != nil && n.Parent.Left != n && n.Parent.Right != n {
		return errors.New("wrong parent")
	}

	if n.Red {
		if n.Left != nil && n.Left.Red {
			return errors.New("left child not black")
		}

		if n.Right != nil && n.Right.Red {
			return errors.New("right child not black")
		}
	}

	_, err := blackHeight(n)
	if err != nil {
		return err
	}

	return nil
}

func blackHeight[T constraints.Ordered](n *RBNode[T]) (int, error) {
	if n == nil {
		return 0, nil
	}

	bl, err := blackHeight(n.Left)
	if err != nil {
		return 0, err
	}

	br, err := blackHeight(n.Right)
	if err != nil {
		return 0, err
	}

	if bl != br {
		return 0, fmt.Errorf("black height differs for %v; %d != %d", n.Value, bl, br)
	}

	if n.Black() {
		return bl + 1, nil
	}

	return bl, nil
}
