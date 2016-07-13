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

func (t *bTree) bTreeDelete(x *bTreeNode, k rune) {
	i := 1
	for i <= x.n && k > x.key[i-1] {
		i = i + 1
	}
	if i <= x.n && k == x.key[i-1] {
		if x.leaf { // case 1
			for j := i - 1; j <= x.n-2; j++ {
				x.key[j] = x.key[j+1]
			}
			x.n = x.n - 1
		} else { // case 2
			y := x.c[i-1]
			z := x.c[i]
			if y.n >= t._t { // case 2a
				node, kp := t.bTreePredecessor(y, k)
				t.bTreeDelete(node, kp)
				x.key[i-1] = kp
			} else if z.n >= t._t { // case 2b
				node, kp := t.bTreeSuccessor(z, k)
				t.bTreeDelete(node, kp)
				x.key[i-1] = kp
			} else { // case 2c
				t.bTreeMerge(x, i)
				t.bTreeDelete(y, k)
			}
		}
	} else { // case 3
		t.bTreeDelete(x.c[i-1], k)
	}
}

func (t *bTree) bTreePredecessor(x *bTreeNode, k rune) (*bTreeNode, rune) {
	if x.leaf {
		return x, x.key[x.n-1]
	}
	return t.bTreePredecessor(x.c[x.n], k)
}

func (t *bTree) bTreeSuccessor(x *bTreeNode, k rune) (*bTreeNode, rune) {
	if x.leaf {
		return x, x.key[0]
	}
	return t.bTreeSuccessor(x.c[0], k)
}

func (t *bTree) bTreeMerge(x *bTreeNode, i int) {
	k := x.key[i-1]
	y := x.c[i-1]
	z := x.c[i]

	y.key[y.n] = k
	y.n = y.n + 1

	var (
		j int
	)
	for j = range z.key[:z.n] {
		y.key[y.n+j] = z.key[j]
		y.c[y.n+j] = z.c[j]
	}
	y.c[y.n+j+1] = z.c[j+1]
	y.n = y.n + z.n

	for j = i - 1; j <= x.n-2; j++ {
		x.key[j] = x.key[j+1]
		x.c[j+1] = x.c[j+2]
	}
	x.n = x.n - 1
}
