package clrs

import (
	"testing"
)

func buildTestRootedTree() *rootedTree {
	rt := new(rootedTree)
	var nodes []*rootedTreeNode
	for i := 0; i <= 13; i++ {
		n := new(rootedTreeNode)
		n.key = i
		nodes = append(nodes, n)
	}
	nodes[0].leftChild = nodes[1]
	nodes[1].p = nodes[0]
	nodes[1].leftChild = nodes[4]
	nodes[1].rightSibling = nodes[2]
	nodes[2].p = nodes[0]
	nodes[2].leftChild = nodes[6]
	nodes[2].rightSibling = nodes[3]
	nodes[3].p = nodes[0]
	nodes[3].leftChild = nodes[10]
	nodes[4].p = nodes[1]
	nodes[4].rightSibling = nodes[5]
	nodes[5].p = nodes[1]
	nodes[5].leftChild = nodes[11]
	nodes[6].p = nodes[2]
	nodes[6].rightSibling = nodes[7]
	nodes[7].p = nodes[2]
	nodes[7].rightSibling = nodes[8]
	nodes[8].p = nodes[2]
	nodes[8].leftChild = nodes[12]
	nodes[8].rightSibling = nodes[9]
	nodes[9].p = nodes[2]
	nodes[10].p = nodes[3]
	nodes[11].p = nodes[5]
	nodes[12].p = nodes[8]
	nodes[12].rightSibling = nodes[13]
	nodes[13].p = nodes[8]
	rt.root = nodes[0]
	return rt
}

func TestRootedTreeTraversal(t *testing.T) {
	rt := buildTestRootedTree()
	want := []int{0, 1, 4, 5, 11, 2, 6, 7, 8, 12, 13, 9, 3, 10}
	i := 0
	for key := range rt.traversalIter() {
		if key != want[i] {
			t.Errorf(" got %v", key)
			t.Errorf("want %v", want[i])
		}
		i = i + 1
	}
}
