package clrs

type bTreeNode struct {
	n    int
	key  []rune
	leaf bool
	c    []*bTreeNode
}

type bTree struct {
	root *bTreeNode
	_t   int
}

func diskRead(x *bTreeNode) {
}

func diskWrite(x *bTreeNode) {
}

func (t bTree) allocateNode() *bTreeNode {
	x := new(bTreeNode)
	x.key = make([]rune, 2*t._t-1)
	x.c = make([]*bTreeNode, 2*t._t)
	return x
}

func (t bTree) bTreeSearch(x *bTreeNode, k rune) (*bTreeNode, int) {
	i := 1
	for i <= x.n && k > x.key[i-1] {
		i = i + 1
	}
	if i <= x.n && k == x.key[i-1] {
		return x, i
	} else if x.leaf {
		return nil, 0
	} else {
		diskRead(x.c[i])
		return t.bTreeSearch(x.c[i-1], k)
	}
}

func (t *bTree) bTreeCreate() {
	x := t.allocateNode()
	x.leaf = true
	x.n = 0
	diskWrite(x)
	t.root = x
}

func (t *bTree) bTreeSplitChild(x *bTreeNode, i int) {
	z := t.allocateNode()
	y := x.c[i-1]
	z.leaf = y.leaf
	z.n = t._t - 1
	for j := 1; j <= z.n; j++ {
		z.key[j-1] = y.key[j+t._t-1]
	}
	if !y.leaf {
		for j := 1; j <= t._t; j++ {
			z.c[j-1] = y.c[j+t._t-1]
		}
	}
	y.n = t._t - 1
	for j := x.n + 1; j >= i+1; j-- {
		x.c[j] = x.c[j-1]
	}
	x.c[i] = z
	for j := x.n; j >= i; j-- {
		x.key[j] = x.key[j-1]
	}
	x.key[i-1] = y.key[t._t-1]
	x.n = x.n + 1
	diskWrite(y)
	diskWrite(z)
	diskWrite(x)
}

func (t *bTree) bTreeInsert(k rune) {
	r := t.root
	if r.n == 2*t._t-1 {
		s := t.allocateNode()
		t.root = s
		s.leaf = false
		s.n = 0
		s.c[0] = r
		t.bTreeSplitChild(s, 1)
		t.bTreeInsertNonfull(s, k)
	} else {
		t.bTreeInsertNonfull(r, k)
	}
}

func (t *bTree) bTreeInsertNonfull(x *bTreeNode, k rune) {
	i := x.n
	if x.leaf {
		for i >= 1 && k < x.key[i-1] {
			x.key[i] = x.key[i-1]
			i = i - 1
		}
		x.key[i] = k
		x.n = x.n + 1
		diskWrite(x)
	} else {
		for i >= 1 && k < x.key[i-1] {
			i = i - 1
		}
		i = i + 1
		diskRead(x.c[i-1])
		if x.c[i-1].n == 2*t._t-1 {
			t.bTreeSplitChild(x, i)
			if k > x.key[i-1] {
				i = i + 1
			}
		}
		t.bTreeInsertNonfull(x.c[i-1], k)
	}
}
