package clrs

import (
	"reflect"
	"testing"
)

var (
	bTreeNodes []*bTreeNode
)

func buildTestBTree() *bTree {
	bTreeNodes = make([]*bTreeNode, 0)
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
			t.Errorf("   search %c", c.k)
			t.Errorf(" got node %v[%v]", node, index)
			t.Errorf("want node %v[%v]", c.wantNode, c.wantIndex)
		}
	}
}

func TestBTreeInsert(t *testing.T) {
	bt := buildTestBTree()
	bt.bTreeInsert('B')
	wantRoot := []rune{'G', 'M', 'P', 'X', 0}
	if !reflect.DeepEqual(bt.root.key, wantRoot) {
		t.Errorf(" got keys %v", bt.root.key)
		t.Errorf("want keys %v", wantRoot)
	}
	wantChildren := [][]rune{
		[]rune{'A', 'B', 'C', 'D', 'E'},
		[]rune{'J', 'K'},
		[]rune{'N', 'O'},
		[]rune{'R', 'S', 'T', 'U', 'V'},
		[]rune{'Y', 'Z'},
	}
	for i, c := range wantChildren {
		node := bt.root.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert B:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}

	bt.bTreeInsert('Q')
	wantRoot = []rune{'G', 'M', 'P', 'T', 'X'}
	if !reflect.DeepEqual(bt.root.key, wantRoot) {
		t.Errorf("after insert Q:")
		t.Errorf(" got keys %v", bt.root.key)
		t.Errorf("want keys %v", wantRoot)
	}
	wantChildren = [][]rune{
		[]rune{'A', 'B', 'C', 'D', 'E'},
		[]rune{'J', 'K'},
		[]rune{'N', 'O'},
		[]rune{'Q', 'R', 'S'},
		[]rune{'U', 'V'},
		[]rune{'Y', 'Z'},
	}
	for i, c := range wantChildren {
		node := bt.root.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert Q:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}

	bt.bTreeInsert('L')
	wantRoot = []rune{'P'}
	if !reflect.DeepEqual(bt.root.key[:bt.root.n], wantRoot) {
		t.Errorf("after insert L:")
		t.Errorf(" got keys %v", bt.root.key[:bt.root.n])
		t.Errorf("want keys %v", wantRoot)
	}
	wantChildren = [][]rune{
		[]rune{'G', 'M'},
		[]rune{'T', 'X'},
	}
	for i, c := range wantChildren {
		node := bt.root.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert L:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}
	lnode := bt.root.c[0]
	rnode := bt.root.c[1]
	wantChildren = [][]rune{
		[]rune{'A', 'B', 'C', 'D', 'E'},
		[]rune{'J', 'K', 'L'},
		[]rune{'N', 'O'},
	}
	for i, c := range wantChildren {
		node := lnode.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert L:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}
	wantChildren = [][]rune{
		[]rune{'Q', 'R', 'S'},
		[]rune{'U', 'V'},
		[]rune{'Y', 'Z'},
	}
	for i, c := range wantChildren {
		node := rnode.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert L:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}

	bt.bTreeInsert('F')
	wantRoot = []rune{'P'}
	if !reflect.DeepEqual(bt.root.key[:bt.root.n], wantRoot) {
		t.Errorf("after insert F:")
		t.Errorf(" got keys %v", bt.root.key[:bt.root.n])
		t.Errorf("want keys %v", wantRoot)
	}
	wantChildren = [][]rune{
		[]rune{'C', 'G', 'M'},
		[]rune{'T', 'X'},
	}
	for i, c := range wantChildren {
		node := bt.root.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert F:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}
	lnode = bt.root.c[0]
	rnode = bt.root.c[1]
	wantChildren = [][]rune{
		[]rune{'A', 'B'},
		[]rune{'D', 'E', 'F'},
		[]rune{'J', 'K', 'L'},
		[]rune{'N', 'O'},
	}
	for i, c := range wantChildren {
		node := lnode.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert F:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}
	wantChildren = [][]rune{
		[]rune{'Q', 'R', 'S'},
		[]rune{'U', 'V'},
		[]rune{'Y', 'Z'},
	}
	for i, c := range wantChildren {
		node := rnode.c[i]
		if !reflect.DeepEqual(node.key[:node.n], c) {
			t.Errorf("after insert F:")
			t.Errorf(" got keys %v", node.key[:node.n])
			t.Errorf("want keys %v", c)
		}
	}
}
