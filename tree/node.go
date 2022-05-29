package tree

import "golang.org/x/exp/constraints"

type Node[T constraints.Ordered] interface {
	LeftNode() Node[T]
	RightNode() Node[T]
	NodeValue() T
}

func Height[T constraints.Ordered](n Node[T]) int {
	if n == nil {
		return 0
	}

	lh := Height(n.LeftNode())
	rh := Height(n.RightNode())

	h := lh
	if rh > h {
		h = rh
	}

	return h + 1
}

func Count[T constraints.Ordered](n Node[T]) int {
	if n == nil {
		return 0
	}

	lc := Count(n.LeftNode())
	rc := Count(n.RightNode())

	return lc + rc + 1
}

func Equal[T constraints.Ordered](n1, n2 Node[T]) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.NodeValue() != n2.NodeValue() {
		return false
	}

	return Equal(n1.LeftNode(), n2.LeftNode()) &&
		Equal(n1.RightNode(), n2.RightNode())
}
