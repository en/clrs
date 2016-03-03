package clrs

type binaryTreeNode struct {
	key   int
	p     *binaryTreeNode
	left  *binaryTreeNode
	right *binaryTreeNode
}

type binaryTree struct {
	root *binaryTreeNode
}

func (t *binaryTree) recursiveTraversalIter() chan int {
	c := make(chan int)
	go func() {
		t.recursiveTraversal(t.root, c)
		close(c)
	}()
	return c
}

func (t *binaryTree) recursiveTraversal(n *binaryTreeNode, c chan int) {
	if n == nil {
		return
	}
	c <- n.key
	if n.left != nil {
		t.recursiveTraversal(n.left, c)
	}
	if n.right != nil {
		t.recursiveTraversal(n.right, c)
	}
}

func (t *binaryTree) nonrecursiveTraversalIter() chan int {
	c := make(chan int)
	go func() {
		t.nonrecursiveTraversal(c)
		close(c)
	}()
	return c
}

func (t *binaryTree) nonrecursiveTraversal(c chan int) {
	s := stack{}
	s.New()
	if t.root != nil {
		s.push(t.root)
	}
	for !s.empty() {
		d, _ := s.pop()
		top := d.(*binaryTreeNode)
		c <- top.key
		if top.right != nil {
			s.push(top.right)
		}
		if top.left != nil {
			s.push(top.left)
		}
	}
}
