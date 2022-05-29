package bintree

import (
	"github.com/hummerd/algorithms/tree"
	"golang.org/x/exp/constraints"
)

type BinNode[T constraints.Ordered] struct {
	Left  *BinNode[T]
	Right *BinNode[T]
	Value T
}

func (n *BinNode[T]) LeftNode() tree.Node[T] {
	if n.Left == nil {
		return nil
	}
	return n.Left
}

func (n *BinNode[T]) RightNode() tree.Node[T] {
	if n.Right == nil {
		return nil
	}
	return n.Right
}

func (n *BinNode[T]) NodeValue() T {
	return n.Value
}

type BinTree[T constraints.Ordered] struct {
	Root *BinNode[T]
}

func (r *BinTree[T]) RootNode() tree.Node[T] {
	if r.Root == nil {
		return nil
	}
	return r.Root
}

func (r *BinTree[T]) Insert(v T) *BinNode[T] {
	if r.Root == nil {
		r.Root = &BinNode[T]{
			Value: v,
		}
		return r.Root
	}

	p := r.Root
	for p != nil {
		if p.Value >= v {
			if p.Left == nil {
				p.Left = &BinNode[T]{
					Value: v,
				}
				return p.Left
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &BinNode[T]{
					Value: v,
				}
				return p.Right
			}
			p = p.Right
		}
	}

	return nil
}

func (r *BinTree[T]) Find(v T) *BinNode[T] {
	p := r.Root
	for p != nil {
		if p.Value == v {
			return p
		}

		if p.Value >= v {
			p = p.Left
		} else {
			p = p.Right
		}
	}

	return p
}

func (r *BinTree[T]) Delete(v T) *BinNode[T] {
	var pp *BinNode[T]
	p := r.Root
	for p != nil {
		if p.Value == v {
			break
		}

		pp = p
		if p.Value >= v {
			p = p.Left
		} else {
			p = p.Right
		}
	}

	if p == nil {
		return nil
	}

	if p.Left == nil && p.Right == nil {
		if pp != nil {
			if pp.Left == p {
				pp.Left = nil
			} else {
				pp.Right = nil
			}
		} else {
			r.Root = nil
		}
		return p
	}

	if p.Left == nil || p.Right == nil {
		c := p.Left
		if c == nil {
			c = p.Right
		}

		if pp != nil {
			if pp.Left == p {
				pp.Left = c
			} else {
				pp.Right = c
			}
		} else {
			r.Root = c
		}

		return p
	}

	var llp *BinNode[T]
	ll := p.Right
	for ll.Left != nil {
		llp = ll
		ll = ll.Left
	}

	if llp != nil {
		llp.Left = ll.Right
	}

	if pp != nil {
		if pp.Left == p {
			pp.Left = ll
		} else {
			pp.Right = ll
		}
	} else {
		r.Root = ll
	}

	ll.Left = p.Left
	if p.Right != ll {
		ll.Right = p.Right
	}

	return p
}
