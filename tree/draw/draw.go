package draw

import (
	"fmt"
	"io"
	"math"
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/hummerd/algorithms/tree"
	"github.com/hummerd/algorithms/tree/rbtree"
	"golang.org/x/exp/constraints"
)

const (
	pad    = 4
	radius = 16
)

// DrawSVGFile generates svg for subtree n to file fileName.
func DrawSVGFile[T constraints.Ordered](fileName string, n tree.Node[T]) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer f.Close()

	DrawSVG(f, n)

	return nil
}

// DrawSVG generates svg for subtree n to out.
func DrawSVG[T constraints.Ordered](out io.Writer, n tree.Node[T]) {
	h := tree.Height(n)

	lastRowCount := math.Pow(2, float64(h-1))

	width := (2*radius+pad)*int(lastRowCount) + pad
	height := (2*radius+pad)*h + pad

	canvas := svg.New(out)
	canvas.Start(width, height)

	drawNode(canvas, n, 0, width, 0)

	canvas.End()
}

func drawNode[T constraints.Ordered](canvas *svg.SVG, n tree.Node[T], left, right, height int) {
	if n == nil {
		return
	}

	h := height + pad + radius
	m := (right + left) / 2

	fill := "gray"

	rbn, ok := n.(*rbtree.RBNode[T])
	if ok && rbn.Red {
		fill = "red"
	}

	if n.LeftNode() != nil {
		canvas.Line(m, h, (left+m)/2, h+2*radius, "stroke-width:2;stroke:black")
	}

	if n.RightNode() != nil {
		canvas.Line(m, h, (m+right)/2, h+2*radius, "stroke-width:2;stroke:black")
	}

	canvas.Circle(m, h, radius, "fill:"+fill)
	canvas.Text(m, h+radius/4, fmt.Sprintf("%v", n.NodeValue()), "text-anchor:middle;font-size:16px;fill:white")

	drawNode(canvas, n.LeftNode(), left, m, h+radius)
	drawNode(canvas, n.RightNode(), m, right, h+radius)
}
