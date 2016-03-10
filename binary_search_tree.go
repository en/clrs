package clrs

// BSTNode ...
type BSTNode struct {
	key   int
	p     *BSTNode
	left  *BSTNode
	right *BSTNode
}

// BST ...
type BST struct {
	root *BSTNode
}

func (t *BST) toArray() []int {
	a := []int{}
	for key := range t.inorderTreeWalkIter() {
		a = append(a, key)
	}
	return a
}

func (t *BST) inorderTreeWalkIter() chan int {
	c := make(chan int)
	go func() {
		t.inorderTreeWalk(t.root, c)
		close(c)
	}()
	return c
}

func (t *BST) inorderTreeWalk(x *BSTNode, c chan int) {
	if x != nil {
		t.inorderTreeWalk(x.left, c)
		c <- x.key
		t.inorderTreeWalk(x.right, c)
	}
}

func (t *BST) iterativePreorderTreeWalkIter() chan int {
	c := make(chan int)
	go func() {
		t.iterativePreorderTreeWalk(c)
		close(c)
	}()
	return c
}

func (t *BST) iterativePreorderTreeWalk(c chan int) {
	s := stack{}
	s.New()
	if t.root != nil {
		s.push(t.root)
	}
	for !s.empty() {
		d, _ := s.pop()
		top := d.(*BSTNode)
		c <- top.key
		if top.right != nil {
			s.push(top.right)
		}
		if top.left != nil {
			s.push(top.left)
		}
	}
}

func (t *BST) constantExtraSpaceTreeWalkIter() chan int {
	c := make(chan int)
	go func() {
		t.constantExtraSpaceTreeWalk(c)
		close(c)
	}()
	return c
}

func (t *BST) constantExtraSpaceTreeWalk(c chan int) {
	if t.root == nil {
		return
	}
	node := t.root
	var prev *BSTNode
	for node != nil {
		if prev == node.p {
			c <- node.key
			prev = node
			if node.left != nil {
				node = node.left
			} else if node.right != nil {
				node = node.right
			} else {
				node = node.p
			}
		} else if (prev == node.left) && (node.right != nil) {
			prev = node
			node = node.right
		} else {
			prev = node
			node = node.p
		}
	}
}

func (t *BST) treeSearch(x *BSTNode, k int) *BSTNode {
	if x == nil || k == x.key {
		return x
	}
	if k < x.key {
		return t.treeSearch(x.left, k)
	}
	return t.treeSearch(x.right, k)
}

func (t *BST) iterativeTreeSearch(x *BSTNode, k int) *BSTNode {
	for x != nil && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (t *BST) treeMinimum(x *BSTNode) *BSTNode {
	for x.left != nil {
		x = x.left
	}
	return x
}

func (t *BST) treeMaximum(x *BSTNode) *BSTNode {
	for x.right != nil {
		x = x.right
	}
	return x
}

func (t *BST) treeSuccessor(x *BSTNode) *BSTNode {
	if x.right != nil {
		return t.treeMinimum(x.right)
	}
	y := x.p
	for y != nil && x == y.right {
		x = y
		y = y.p
	}
	return y
}

func (t *BST) treePredecessor(x *BSTNode) *BSTNode {
	if x.left != nil {
		return t.treeMaximum(x.left)
	}
	y := x.p
	for y != nil && x == y.left {
		x = y
		y = y.p
	}
	return y
}

func (t *BST) treeInsert(z *BSTNode) {
	var y *BSTNode
	x := t.root
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nil {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func (t *BST) transplant(u, v *BSTNode) {
	if u.p == nil {
		t.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	if v != nil {
		v.p = u.p
	}
}

func (t *BST) treeDelete(z *BSTNode) {
	if z.left == nil {
		t.transplant(z, z.right)
	} else if z.right == nil {
		t.transplant(z, z.left)
	} else {
		y := t.treeMinimum(z.right)
		if y.p != z {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.p = y
	}
}
