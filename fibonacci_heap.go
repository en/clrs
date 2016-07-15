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
