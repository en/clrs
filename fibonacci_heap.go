package clrs

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

func (h *fibHeap) listConcatenate(h2 *fibHeap) {
	if h.min == nil {
		h.min = h2.min
	} else {
		h2.min.left.right = h.min
		h2.min.left, h.min.left = h.min.left, h2.min.left
		h2.min.left.right = h2.min
	}
}
