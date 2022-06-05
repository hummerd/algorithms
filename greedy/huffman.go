package greedy

import (
	"bufio"
	"container/heap"
	"encoding/binary"
	"io"
)

type BitSeq struct {
	Bits uint64
	Len  int
}

func BuildHuffCodes(r io.Reader) map[rune]BitSeq {
	sf := symbolFreq(r)
	return BuildHuffCodesByFreq(sf)
}

func BuildHuffCodesByFreq(freq map[rune]int) map[rune]BitSeq {
	ht := huffCodes(freq)

	result := map[rune]BitSeq{}
	huffCodesMap(*ht.l, result, true, BitSeq{})
	huffCodesMap(*ht.r, result, false, BitSeq{})
	return result
}

func symbolFreq(r io.Reader) map[rune]int {
	f := make(map[rune]int)

	s := bufio.NewScanner(r)
	for s.Scan() {
		for _, r := range s.Text() {
			f[r] += 1
		}
	}

	return f
}

type huffNode struct {
	f int
	s rune
	l *huffNode
	r *huffNode
}

type huffNodeHeap []huffNode

func (h *huffNodeHeap) Len() int {
	return len(*h)
}

func (h *huffNodeHeap) Less(i, j int) bool {
	return (*h)[i].f < (*h)[j].f
}

func (h *huffNodeHeap) Swap(i, j int) {
	hv := *h
	hv[i], hv[j] = hv[j], hv[i]
}

func (h *huffNodeHeap) Push(x any) {
	*h = append(*h, x.(huffNode))
}

func (h *huffNodeHeap) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

type bitWriter struct {
	tail  int
	buff  uint64
	bytes []byte
	w     io.Writer
}

func (bw *bitWriter) Write(v uint64, l int) {
	if l > bw.tail {
		bw.buff |= v << (64 - bw.tail)
		binary.BigEndian.PutUint64(bw.bytes, bw.buff)
		bw.w.Write(bw.bytes)

		v = v >> uint64(bw.tail)
		l = l - bw.tail

		bw.buff = 0
		bw.tail = 64
	}

	bw.buff |= v << (64 - bw.tail)
	bw.tail = bw.tail - l

	if bw.tail == 0 {
		binary.BigEndian.PutUint64(bw.bytes, bw.buff)
		bw.w.Write(bw.bytes)
		bw.buff = 0
		bw.tail = 64
	}
}

func (bw *bitWriter) Flush() {
	if bw.tail == 64 {
		return
	}

	binary.BigEndian.PutUint64(bw.bytes, bw.buff)
	bw.w.Write(bw.bytes)
}

func huffCodesMap(n huffNode, r map[rune]BitSeq, left bool, pbs BitSeq) {
	var a uint64
	if !left {
		a = 1
	}

	bs := BitSeq{
		Bits: (pbs.Bits << 1) | a,
		Len:  pbs.Len + 1,
	}

	if n.l == nil {
		r[n.s] = bs
		return
	}

	huffCodesMap(*n.l, r, true, bs)
	huffCodesMap(*n.r, r, false, bs)
}

func huffCodes(freq map[rune]int) huffNode {
	h := make([]huffNode, 0, len(freq))
	for k, v := range freq {
		h = append(h, huffNode{
			f: v,
			s: k,
		})
	}

	if len(h) == 0 {
		panic("empty input")
	}

	if len(h) == 1 {
		return h[0]
	}

	hh := huffNodeHeap(h)
	heap.Init(&hh)

	var root huffNode

	for {
		f := heap.Pop(&hh).(huffNode)
		s := heap.Pop(&hh).(huffNode)

		n := huffNode{
			f: f.f + s.f,
			l: &f,
			r: &s,
		}

		if hh.Len() == 0 {
			root = n
			break
		}

		heap.Push(&hh, n)
	}

	return root
}
