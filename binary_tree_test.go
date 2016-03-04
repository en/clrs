package clrs

import (
	"testing"
)

func buildTestBinaryTree() *binaryTree {
	bt := new(binaryTree)
	var nodes []*binaryTreeNode
	for i := 1; i <= 11; i++ {
		n := new(binaryTreeNode)
		n.key = i
		nodes = append(nodes, n)
	}
	nodes[0].left = nodes[1]
	nodes[0].right = nodes[2]
	nodes[1].p = nodes[0]
	nodes[1].right = nodes[3]
	nodes[2].p = nodes[0]
	nodes[2].left = nodes[4]
	nodes[2].right = nodes[5]
	nodes[3].p = nodes[1]
	nodes[3].left = nodes[6]
	nodes[3].right = nodes[7]
	nodes[4].p = nodes[2]
	nodes[4].right = nodes[8]
	nodes[5].p = nodes[2]
	nodes[5].left = nodes[9]
	nodes[6].p = nodes[3]
	nodes[7].p = nodes[3]
	nodes[8].p = nodes[4]
	nodes[8].left = nodes[10]
	nodes[9].p = nodes[5]
	nodes[10].p = nodes[8]
	bt.root = nodes[0]
	return bt
}

func TestBinaryTreeRecursiveTraversal(t *testing.T) {
	bt := buildTestBinaryTree()
	want := []int{1, 2, 4, 7, 8, 3, 5, 9, 11, 6, 10}
	i := 0
	for key := range bt.recursiveTraversalIter() {
		if key != want[i] {
			t.Errorf(" got %v", key)
			t.Errorf("want %v", want[i])
		}
		i = i + 1
	}
}

func TestBinaryTreeNonrecursiveTraversal(t *testing.T) {
	bt := buildTestBinaryTree()
	want := []int{1, 2, 4, 7, 8, 3, 5, 9, 11, 6, 10}
	i := 0
	for key := range bt.nonrecursiveTraversalIter() {
		if key != want[i] {
			t.Errorf(" got %v", key)
			t.Errorf("want %v", want[i])
		}
		i = i + 1
	}
}

func TestBinaryTreeNonrecursiveConstantExtraSpaceTraversal(t *testing.T) {
	bt := buildTestBinaryTree()
	want := []int{1, 2, 4, 7, 8, 3, 5, 9, 11, 6, 10}
	i := 0
	for key := range bt.nonrecursiveConstantExtraSpaceTraversalIter() {
		if key != want[i] {
			t.Errorf(" got %v", key)
			t.Errorf("want %v", want[i])
		}
		i = i + 1
	}
}
