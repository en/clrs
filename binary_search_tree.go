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

func (t *BST) preorderTreeWalkIter() chan int {
	c := make(chan int)
	go func() {
		t.preorderTreeWalk(t.root, c)
		close(c)
	}()
	return c
}

func (t *BST) preorderTreeWalk(x *BSTNode, c chan int) {
	if x != nil {
		c <- x.key
		t.preorderTreeWalk(x.left, c)
		t.preorderTreeWalk(x.right, c)
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

func (t *BST) constantExtraSpacePreorderTreeWalkIter() chan int {
	c := make(chan int)
	go func() {
		t.constantExtraSpacePreorderTreeWalk(c)
		close(c)
	}()
	return c
}

func (t *BST) constantExtraSpacePreorderTreeWalk(c chan int) {
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
