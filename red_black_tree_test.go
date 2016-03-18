package clrs

import (
	"reflect"
	"testing"
)

var (
	rbnodes []*rbnode
)

func buildTestRBT() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.3
	for _, v := range []int{7, 4, 11, 3, 6, 9, 18, 2, 14, 19, 12, 17, 22, 20} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = rbnodes[4]

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[5]
	rbnodes[2].right = rbnodes[6]

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = rbnodes[7]

	rbnodes[4].p = rbnodes[1]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = nilNode

	rbnodes[6].p = rbnodes[2]
	rbnodes[6].left = rbnodes[8]
	rbnodes[6].right = rbnodes[9]

	rbnodes[7].p = rbnodes[3]
	rbnodes[7].left = nilNode
	rbnodes[7].right = nilNode

	rbnodes[8].p = rbnodes[6]
	rbnodes[8].left = rbnodes[10]
	rbnodes[8].right = rbnodes[11]

	rbnodes[9].p = rbnodes[6]
	rbnodes[9].left = nilNode
	rbnodes[9].right = rbnodes[12]

	rbnodes[10].p = rbnodes[8]
	rbnodes[10].left = nilNode
	rbnodes[10].right = nilNode

	rbnodes[11].p = rbnodes[8]
	rbnodes[11].left = nilNode
	rbnodes[11].right = nilNode

	rbnodes[12].p = rbnodes[9]
	rbnodes[12].left = rbnodes[13]
	rbnodes[12].right = nilNode

	rbnodes[13].p = rbnodes[12]
	rbnodes[13].left = nilNode
	rbnodes[13].right = nilNode

	rbt.root = rbnodes[0]
	return rbt
}

func TestRedBlackTreeToArray(t *testing.T) {
	rbt := buildTestRBT()
	want := []int{
		7,
		4, 11,
		3, 6, 9, 18,
		2, -1, -1, -1, -1, -1, 14, 19,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 12, 17, -1, 22,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 20, -1,
	}
	got := rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestRedBlackTreeRotate(t *testing.T) {
	rbt := buildTestRBT()
	want := []int{
		7,
		4, 18,
		3, 6, 11, 19,
		2, -1, -1, -1, 9, 14, -1, 22,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 12, 17, -1, -1, 20, -1,
	}
	rbt.leftRotate(rbnodes[2])
	got := rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	want = []int{
		7,
		4, 11,
		3, 6, 9, 18,
		2, -1, -1, -1, -1, -1, 14, 19,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 12, 17, -1, 22,
		-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 20, -1,
	}
	rbt.rightRotate(rbnodes[6])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
