package clrs

import (
	"math"
)

type fibNode struct {
	key    int
	p      *fibNode
	child  *fibNode
	left   *fibNode
	right  *fibNode
	degree int
	mark   bool
}

type fibHeap struct {
	min *fibNode
	n   int
}

func makeFibHeap() *fibHeap {
	return new(fibHeap)
}

func fibHeapUnion(h1, h2 *fibHeap) *fibHeap {
	h := makeFibHeap()
	h.min = h1.min
	h.listConcatenate(h2)
	if h1.min == nil || h2.min != nil && h2.min.key < h1.min.key {
		h.min = h2.min
	}
	h.n = h1.n + h2.n
	return h
}

func (h *fibHeap) fibHeapInsert(x *fibNode) {
	x.degree = 0
	x.p = nil
	x.child = nil
	x.mark = false
	if h.min == nil {
		x.left = x
		x.right = x
		h.min = x
	} else {
		h.listInsert(h.min, x)
		if x.key < h.min.key {
			h.min = x
		}
	}
	h.n = h.n + 1
}

func (h *fibHeap) listInsert(root, x *fibNode) {
	x.right = root
	x.left = root.left
	x.left.right = x
	root.left = x
}

func (h *fibHeap) listDelete(x *fibNode) {
	x.right.left = x.left
	x.left.right = x.right
}

func (h *fibHeap) listConcatenate(h2 *fibHeap) {
	if h.min == nil {
		h.min = h2.min
	} else {
		h2.min.left.right = h.min
		h2.min.left, h.min.left = h.min.left, h2.min.left
		h2.min.left.right = h2.min
	}
}

func (h *fibHeap) fibHeapExtractMin() *fibNode {
	z := h.min
	if z != nil {
		for _, x := range h.children(z.child) {
			h.listInsert(z, x)
			x.p = nil
		}
		h.listDelete(z)
		if z == z.right {
			h.min = nil
		} else {
			h.min = z.right
			h.consolidate()
		}
		h.n = h.n - 1
	}
	return z
}

func (h *fibHeap) children(x *fibNode) []*fibNode {
	var cs []*fibNode
	for c := x; c != nil; c = c.right {
		if len(cs) > 0 && c == cs[0] {
			break
		}
		cs = append(cs, c)
	}
	return cs
}

func (h *fibHeap) d() int {
	return int(math.Log2(float64(h.n)))
}

func (h *fibHeap) consolidate() {
	hd := h.d()
	a := make([]*fibNode, hd+1)
	for _, w := range h.children(h.min) {
		x := w
		d := x.degree
		for a[d] != nil {
			y := a[d]
			if x.key > y.key {
				x, y = y, x
			}
			h.fibHeapLink(y, x)
			a[d] = nil
			d = d + 1
		}
		a[d] = x
	}
	h.min = nil
	for i := 0; i <= hd; i++ {
		if a[i] != nil {
			if h.min == nil {
				a[i].left = a[i]
				a[i].right = a[i]
				h.min = a[i]
			} else {
				h.listInsert(h.min, a[i])
				a[i].p = nil
				if a[i].key < h.min.key {
					h.min = a[i]
				}
			}
		}
	}
}

func (h *fibHeap) fibHeapLink(y, x *fibNode) {
	h.listDelete(y)
	if x.child == nil {
		y.left = y
		y.right = y
		x.child = y
	} else {
		h.listInsert(x.child, y)
	}
	y.p = x
	x.degree = x.degree + 1
	y.mark = false
}

func (h *fibHeap) fibHeapWalk() []*fibNode {
	var (
		nodes   []*fibNode
		child   func(x *fibNode)
		sibling func(x *fibNode)
	)
	child = func(x *fibNode) {
		if x.child != nil {
			sibling(x.child)
		}
	}
	sibling = func(x *fibNode) {
		start := -1
		for s := x; s != nil; s = s.right {
			if start != -1 && nodes[start] == s {
				break
			}
			child(s)
			nodes = append(nodes, s)
			if start == -1 {
				start = len(nodes) - 1
			}
		}
	}

	if h.min != nil {
		sibling(h.min)
	}
	return nodes
}

func (h *fibHeap) fibHeapDecreaseKey(x *fibNode, k int) error {
	if k > x.key {
		return errGreater
	}
	x.key = k
	y := x.p
	if y != nil && x.key < y.key {
		h.cut(x, y)
		h.cascadingCut(y)
	}
	if x.key < h.min.key {
		h.min = x
	}
	return nil
}

func (h *fibHeap) cut(x, y *fibNode) {
	h.listDelete(x)
	if x == x.right {
		y.child = nil
	}
	y.degree = y.degree - 1
	h.listInsert(h.min, x)
	x.p = nil
	x.mark = false
}

func (h *fibHeap) cascadingCut(y *fibNode) {
	z := y.p
	if z != nil {
		if !y.mark {
			y.mark = true
		} else {
			h.cut(y, z)
			h.cascadingCut(z)
		}
	}
}

func (h *fibHeap) fibHeapDelete(x *fibNode) {
	if h.min != nil {
		h.fibHeapDecreaseKey(x, h.min.key-1)
		h.fibHeapExtractMin()
	}
}
