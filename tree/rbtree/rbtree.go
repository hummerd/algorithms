package rbtree

import (
	"fmt"

	"github.com/hummerd/algorithms/tree"
	"golang.org/x/exp/constraints"
)

type RBTree[T constraints.Ordered] struct {
	Root *RBNode[T]
}

func (t *RBTree[T]) Insert(v T) {
	if t.Root == nil {
		t.Root = &RBNode[T]{
			Value: v,
		}
		return
	}

	top := t.Root.insert(v)
	if top.Parent == nil {
		t.Root = top
	} else if top.Parent.Parent == nil {
		t.Root = top.Parent
	}
}

func (t *RBTree[T]) Delete(v T) {
	n := t.Root.find(v)
	if n == nil {
		return
	}

	c := n.delete()
	if c == nil || c.Parent == nil {
		t.Root = c
	} else if c.Parent.Parent == nil {
		t.Root = c.Parent
	} else if c.Parent.Parent.Parent == nil {
		t.Root = c.Parent.Parent
	}
}

func (t *RBTree[T]) Height() int {
	if t.Root == nil {
		return 0
	}

	return t.Root.Height()
}

type RBNode[T constraints.Ordered] struct {
	Left   *RBNode[T]
	Right  *RBNode[T]
	Parent *RBNode[T]
	Red    bool
	Value  T
}

func (n *RBNode[T]) LeftNode() tree.Node[T] {
	if n.Left == nil {
		return nil
	}
	return n.Left
}

func (n *RBNode[T]) RightNode() tree.Node[T] {
	if n.Right == nil {
		return nil
	}
	return n.Right
}

func (n *RBNode[T]) NodeValue() T {
	return n.Value
}

func (n *RBNode[T]) Black() bool {
	if n == nil {
		return true
	}

	return !n.Red
}

func (n *RBNode[T]) find(v T) *RBNode[T] {
	for n != nil {
		if v > n.Value {
			n = n.Right
		} else if v < n.Value {
			n = n.Left
		} else {
			return n
		}
	}

	return n
}

func (n *RBNode[T]) successor() *RBNode[T] {
	if n == nil {
		return nil
	}

	if n.Right != nil {
		return n.Right.min()
	}

	p := n.Parent
	for p != nil && n == p.Right {
		n = p
		p = p.Parent
	}

	return p
}

func (n *RBNode[T]) min() *RBNode[T] {
	if n == nil {
		return nil
	}

	for n.Left != nil {
		n = n.Left
	}

	return n
}

func (n *RBNode[T]) max() *RBNode[T] {
	if n == nil {
		return nil
	}

	for n.Right != nil {
		n = n.Right
	}

	return n
}

func (n *RBNode[T]) delete() *RBNode[T] {
	if n == nil {
		panic("can not delete nil node")
	}

	var d *RBNode[T] // node that will be physically deleted
	if n.Left == nil || n.Right == nil {
		d = n
	} else {
		d = n.successor()
	}

	var c *RBNode[T] // child node that will replace deleted
	if d.Left != nil {
		c = d.Left
	} else {
		c = d.Right
	}

	cfake := c == nil
	if !cfake {
		c.Parent = d.Parent
	} else {
		c = &RBNode[T]{
			Red:    false,
			Parent: d.Parent,
		}
	}

	if d.Parent != nil {
		if d.Parent.Left == d {
			d.Parent.Left = c
		} else {
			d.Parent.Right = c
		}
	}

	if d != n {
		n.Value = d.Value
	}

	pp := c
	if !d.Red {
		pp = c.deleteFixup()
	}

	if cfake {
		if c.Parent != nil {
			if c.Parent.Left == c {
				c.Parent.Left = nil
			} else {
				c.Parent.Right = nil
			}
		} else {
			return nil
		}
	}

	return pp
}

func (n *RBNode[T]) deleteFixup() *RBNode[T] {
	fmt.Println("fixup", n)
	for n.Parent != nil && !n.Red {
		if n == n.Parent.Left {
			fmt.Println("fixup loop left", n)
			// case 1 - transform it to case 2, 3 or 4
			r := n.Parent.Right
			if r.Red {
				r.Red = false
				r.Parent.Red = true
				n.Parent.RotateLeft()
				r = n.Parent.Right
			}

			fmt.Println(r.Value, r.Right, r.Left)
			if r.Right.Black() && r.Left.Black() {
				r.Red = true
				n = n.Parent
			} else {
				if r.Right.Black() {
					r.Left.Red = false
					r.Red = true
					r.RotateRight()
					r = n.Parent.Right
				}

				r.Red = n.Parent.Red
				n.Parent.Red = false
				r.Right.Red = false
				n.Parent.RotateLeft()
				break
			}
		} else {
			fmt.Println("fixup loop right", n)
			l := n.Parent.Left
			fmt.Println(n.Parent)
			if l.Red {
				l.Red = false
				l.Parent.Red = true
				n.Parent.RotateRight()
				l = n.Parent.Left
			}

			if l.Left.Black() && l.Right.Black() {
				l.Red = true
				n = n.Parent
			} else {
				if l.Left.Black() {
					l.Right.Red = false
					l.Red = true
					l.RotateLeft()
					l = n.Parent.Left
				}

				l.Red = n.Parent.Red
				n.Parent.Red = false
				l.Left.Red = false
				n.Parent.RotateRight()
				break
			}
		}
	}

	n.Red = false
	return n
}

func (n *RBNode[T]) insert(v T) *RBNode[T] {
	if n == nil {
		panic("can not insert into nil node")
	}

	var p *RBNode[T]

	for n != nil {
		p = n

		if v > p.Value {
			n = n.Right
		} else {
			n = n.Left
		}
	}

	nn := &RBNode[T]{
		Value:  v,
		Red:    true,
		Parent: p,
	}

	if v > p.Value {
		p.Right = nn
	} else {
		p.Left = nn
	}

	return nn.insertFixup()
}

func (n *RBNode[T]) insertFixup() *RBNode[T] {
	for n.Parent != nil && n.Parent.Red {
		parentLeft := n.Parent.Parent.Left == n.Parent

		var uncle *RBNode[T]
		if parentLeft {
			uncle = n.Parent.Parent.Right
		} else {
			uncle = n.Parent.Parent.Left
		}

		if uncle != nil && uncle.Red {
			uncle.Red = false
			n.Parent.Red = false
			n.Parent.Parent.Red = true
			n = n.Parent.Parent

			if n.Parent == nil {
				n.Red = false
			}

			continue
		}

		if parentLeft {
			if n.Parent.Right == n {
				n = n.Parent
				n.RotateLeft()
			}

			n.Parent.Parent.RotateRight()
			n.Parent.Red = false
			n.Parent.Right.Red = true
		} else {
			if n.Parent.Left == n {
				n = n.Parent
				n.RotateRight()
			}

			n.Parent.Parent.RotateLeft()
			n.Parent.Red = false
			n.Parent.Left.Red = true
		}

		if n.Parent.Parent == nil {
			n.Parent.Red = false
		}
	}

	return n
}

func (n *RBNode[T]) FixParents(p *RBNode[T]) {
	if n == nil {
		return
	}

	n.Parent = p

	n.Left.FixParents(n)
	n.Right.FixParents(n)
}

func (n *RBNode[T]) Height() int {
	if n == nil {
		return 0
	}

	lh := n.Left.Height()
	rh := n.Right.Height()

	if lh > rh {
		return lh + 1
	}

	return rh + 1
}

func (n *RBNode[T]) String() string {
	if n == nil {
		return "<nil>"
	}

	p := ""
	if n.Parent != nil {
		p = "; Parent " + fmt.Sprint(n.Parent.Value)
	}

	c := "b"
	if n.Red {
		c = "r"
	}

	return "Node " + fmt.Sprint(n.Value) + c + p
}

//     N    <-
//   B   C
//      D E
// ------------
//    C
//   N  E
//  B D
func (n *RBNode[T]) RotateLeft() {
	if n == nil {
		return
	}

	c := n.Right
	if c == nil {
		return
	}

	p := n.Parent
	if p != nil {
		p.ReplaceChild(n, c)
	} else {
		c.Parent = nil
	}

	d := c.Left

	c.SetLeft(n)
	n.SetRight(d)
}

//     N    ->
//   B   C
//  D E
// ------------
//      B
//   D     N
//        E C
func (n *RBNode[T]) RotateRight() {
	if n == nil {
		return
	}

	b := n.Left
	if b == nil {
		return
	}

	p := n.Parent
	if p != nil {
		p.ReplaceChild(n, b)
	} else {
		b.Parent = nil
	}

	e := b.Right

	b.SetRight(n)
	n.SetLeft(e)
}

func (n *RBNode[T]) ReplaceChild(old, new *RBNode[T]) {
	if n == nil {
		return
	}

	if n.Left == old {
		n.Left = new
	} else {
		n.Right = new
	}

	if new != nil {
		new.Parent = n
	}
}

func (n *RBNode[T]) SetLeft(l *RBNode[T]) {
	if n == nil {
		return
	}

	n.Left = l
	if l != nil {
		l.Parent = n
	}
}

func (n *RBNode[T]) SetRight(r *RBNode[T]) {
	if n == nil {
		return
	}

	n.Right = r
	if r != nil {
		r.Parent = n
	}
}
