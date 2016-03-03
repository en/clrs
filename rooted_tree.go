package clrs

type rootedTreeNode struct {
	key          int
	p            *rootedTreeNode
	leftChild    *rootedTreeNode
	rightSibling *rootedTreeNode
}

type rootedTree struct {
	root *rootedTreeNode
}

func (t *rootedTree) traversalIter() chan int {
	c := make(chan int)
	go func() {
		t.traversal(c)
		close(c)
	}()
	return c
}

func (t *rootedTree) traversal(c chan int) {
	node := t.root
	up := false
	for !(up && (node == t.root)) {
		if up {
			if node.rightSibling != nil {
				node = node.rightSibling
				up = false
			} else {
				node = node.p
			}
		} else {
			c <- node.key
			if node.leftChild != nil {
				node = node.leftChild
				continue
			}
			if node.rightSibling != nil {
				node = node.rightSibling
				continue
			} else {
				node = node.p
				up = true
			}
		}
	}
}
