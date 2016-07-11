package clrs

import (
	"testing"
)

var (
	bTreeNodes []*bTreeNode
)

func buildTestBTree() *bTree {
	bt := new(bTree)
	bt._t = 3
	bt.bTreeCreate()
	bTreeNodes = append(bTreeNodes, bt.root)

	for _, keys := range [][]rune{
		{'A', 'C', 'D', 'E'},
		{'J', 'K'},
		{'N', 'O'},
		{'R', 'S', 'T', 'U', 'V'},
		{'Y', 'Z'},
	} {
		node := bt.allocateNode()
		node.n = len(keys)
		node.leaf = true
		for i, k := range keys {
			node.key[i] = k
		}
		bTreeNodes = append(bTreeNodes, node)
	}

	bt.root.n = 4
	bt.root.leaf = false
	for i, k := range []rune{'G', 'M', 'P', 'X'} {
		bt.root.key[i] = k
	}
	for i, n := range bTreeNodes[1:] {
		bt.root.c[i] = n
	}
	return bt
}

func TestBTreeSearch(t *testing.T) {
	bt := buildTestBTree()
	cases := []struct {
		k         rune
		wantNode  *bTreeNode
		wantIndex int
	}{
		{'B', nil, 0},
		{'G', bTreeNodes[0], 1},
		{'M', bTreeNodes[0], 2},
		{'P', bTreeNodes[0], 3},
		{'X', bTreeNodes[0], 4},
		{'A', bTreeNodes[1], 1},
		{'C', bTreeNodes[1], 2},
		{'K', bTreeNodes[2], 2},
		{'N', bTreeNodes[3], 1},
		{'V', bTreeNodes[4], 5},
		{'Y', bTreeNodes[5], 1},
	}
	for _, c := range cases {
		node, index := bt.bTreeSearch(bt.root, c.k)
		if node != c.wantNode || index != c.wantIndex {
			t.Errorf("search %c", c.k)
			t.Errorf(" got node %v[%v]", node, index)
			t.Errorf("want node %v[%v]", c.wantNode, c.wantIndex)
		}
	}
}
