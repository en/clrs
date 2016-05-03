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
	rbnodes[3].right = nilNode

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

func buildTestCase1() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.4 (a)
	for _, v := range []int{11, 2, 14, 1, 7, 15, 5, 8, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = rbnodes[4]
	rbnodes[1].color = RED

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = nilNode
	rbnodes[2].right = rbnodes[5]
	rbnodes[2].color = BLACK

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = BLACK

	rbnodes[4].p = rbnodes[1]
	rbnodes[4].left = rbnodes[6]
	rbnodes[4].right = rbnodes[7]
	rbnodes[4].color = BLACK

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = nilNode
	rbnodes[5].color = RED

	rbnodes[6].p = rbnodes[4]
	rbnodes[6].left = rbnodes[8]
	rbnodes[6].right = nilNode
	rbnodes[6].color = RED

	rbnodes[7].p = rbnodes[4]
	rbnodes[7].left = nilNode
	rbnodes[7].right = nilNode
	rbnodes[7].color = RED

	rbnodes[8].p = rbnodes[6]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode
	rbnodes[8].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func buildTestCase2() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.4 (b)
	for _, v := range []int{11, 2, 14, 1, 7, 15, 5, 8, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = rbnodes[4]
	rbnodes[1].color = RED

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = nilNode
	rbnodes[2].right = rbnodes[5]
	rbnodes[2].color = BLACK

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = BLACK

	rbnodes[4].p = rbnodes[1]
	rbnodes[4].left = rbnodes[6]
	rbnodes[4].right = rbnodes[7]
	rbnodes[4].color = RED

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = nilNode
	rbnodes[5].color = RED

	rbnodes[6].p = rbnodes[4]
	rbnodes[6].left = rbnodes[8]
	rbnodes[6].right = nilNode
	rbnodes[6].color = BLACK

	rbnodes[7].p = rbnodes[4]
	rbnodes[7].left = nilNode
	rbnodes[7].right = nilNode
	rbnodes[7].color = BLACK

	rbnodes[8].p = rbnodes[6]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode
	rbnodes[8].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func buildTestCase3() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.4 (c)
	for _, v := range []int{11, 7, 14, 2, 8, 15, 1, 5, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = rbnodes[4]
	rbnodes[1].color = RED

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = nilNode
	rbnodes[2].right = rbnodes[5]
	rbnodes[2].color = BLACK

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = rbnodes[6]
	rbnodes[3].right = rbnodes[7]
	rbnodes[3].color = RED

	rbnodes[4].p = rbnodes[1]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode
	rbnodes[4].color = BLACK

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = nilNode
	rbnodes[5].color = RED

	rbnodes[6].p = rbnodes[4]
	rbnodes[6].left = nilNode
	rbnodes[6].right = nilNode
	rbnodes[6].color = BLACK

	rbnodes[7].p = rbnodes[4]
	rbnodes[7].left = rbnodes[8]
	rbnodes[7].right = nilNode
	rbnodes[7].color = BLACK

	rbnodes[8].p = rbnodes[6]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode
	rbnodes[8].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func buildTestCase4() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// mirror of figure 13.4 (a)
	for _, v := range []int{11, 14, 2, 15, 7, 1, 8, 5, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[4]
	rbnodes[2].right = rbnodes[5]
	rbnodes[2].color = RED

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = RED

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = rbnodes[6]
	rbnodes[4].right = rbnodes[7]
	rbnodes[4].color = BLACK

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = nilNode
	rbnodes[5].color = BLACK

	rbnodes[6].p = rbnodes[4]
	rbnodes[6].left = nilNode
	rbnodes[6].right = nilNode
	rbnodes[6].color = RED

	rbnodes[7].p = rbnodes[4]
	rbnodes[7].left = nilNode
	rbnodes[7].right = rbnodes[8]
	rbnodes[7].color = RED

	rbnodes[8].p = rbnodes[7]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode
	rbnodes[8].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func buildTestCase5() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// mirror of figure 13.4 (b)
	for _, v := range []int{11, 14, 2, 15, 7, 1, 8, 5, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[4]
	rbnodes[2].right = rbnodes[5]
	rbnodes[2].color = RED

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = RED

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = rbnodes[6]
	rbnodes[4].right = rbnodes[7]
	rbnodes[4].color = RED

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = nilNode
	rbnodes[5].color = BLACK

	rbnodes[6].p = rbnodes[4]
	rbnodes[6].left = nilNode
	rbnodes[6].right = nilNode
	rbnodes[6].color = BLACK

	rbnodes[7].p = rbnodes[4]
	rbnodes[7].left = nilNode
	rbnodes[7].right = rbnodes[8]
	rbnodes[7].color = BLACK

	rbnodes[8].p = rbnodes[7]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode
	rbnodes[8].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func buildTestCase6() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// mirror of figure 13.4 (c)
	for _, v := range []int{11, 14, 7, 15, 8, 2, 5, 1, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = rbnodes[3]
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[4]
	rbnodes[2].right = rbnodes[5]
	rbnodes[2].color = RED

	rbnodes[3].p = rbnodes[1]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = RED

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode
	rbnodes[4].color = BLACK

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = rbnodes[6]
	rbnodes[5].right = rbnodes[7]
	rbnodes[5].color = RED

	rbnodes[6].p = rbnodes[5]
	rbnodes[6].left = nilNode
	rbnodes[6].right = rbnodes[8]
	rbnodes[6].color = BLACK

	rbnodes[7].p = rbnodes[5]
	rbnodes[7].left = nilNode
	rbnodes[7].right = nilNode
	rbnodes[7].color = BLACK

	rbnodes[8].p = rbnodes[6]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode
	rbnodes[8].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func TestRedBlackTreeToArray(t *testing.T) {
	rbt := buildTestRBT()
	want := []rbdata{
		{7, BLACK},
		{4, BLACK}, {11, BLACK},
		{3, BLACK}, {6, BLACK}, {9, BLACK}, {18, BLACK},
		{2, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {14, BLACK}, {19, BLACK},
		{-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {12, BLACK}, {17, BLACK}, {-1, BLACK}, {22, BLACK},
		{-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {20, BLACK}, {-1, BLACK},
	}
	got := rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestRedBlackTreeRotate(t *testing.T) {
	rbt := buildTestRBT()
	want := []rbdata{
		{7, BLACK},
		{4, BLACK}, {18, BLACK},
		{3, BLACK}, {6, BLACK}, {11, BLACK}, {19, BLACK},
		{2, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {9, BLACK}, {14, BLACK}, {-1, BLACK}, {22, BLACK},
		{-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {12, BLACK}, {17, BLACK}, {-1, BLACK}, {-1, BLACK}, {20, BLACK}, {-1, BLACK},
	}
	rbt.leftRotate(rbnodes[2])
	got := rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	want = []rbdata{
		{7, BLACK},
		{4, BLACK}, {11, BLACK},
		{3, BLACK}, {6, BLACK}, {9, BLACK}, {18, BLACK},
		{2, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {14, BLACK}, {19, BLACK},
		{-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {12, BLACK}, {17, BLACK}, {-1, BLACK}, {22, BLACK},
		{-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {20, BLACK}, {-1, BLACK},
	}
	rbt.rightRotate(rbnodes[6])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestRedBlackInsertFixup(t *testing.T) {
	want := []rbdata{
		{7, BLACK},
		{2, RED}, {11, RED},
		{1, BLACK}, {5, BLACK}, {8, BLACK}, {14, BLACK},
		{-1, BLACK}, {-1, BLACK}, {4, RED}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {15, RED},
	}
	// case 1
	rbt := buildTestCase1()
	rbt.rbInsertFixup(rbnodes[8])
	got := rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 1 got %v", got)
		t.Errorf("case 1 want %v", want)
	}
	// case 2
	rbt = buildTestCase2()
	rbt.rbInsertFixup(rbnodes[4])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 2 got %v", got)
		t.Errorf("case 2 want %v", want)
	}
	// case 3
	rbt = buildTestCase3()
	rbt.rbInsertFixup(rbnodes[3])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 2 got %v", got)
		t.Errorf("case 2 want %v", want)
	}
	want = []rbdata{
		{7, BLACK},
		{11, RED}, {2, RED},
		{14, BLACK}, {8, BLACK}, {5, BLACK}, {1, BLACK},
		{15, RED}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {-1, BLACK}, {4, RED}, {-1, BLACK}, {-1, BLACK},
	}
	// case 4
	rbt = buildTestCase4()
	rbt.rbInsertFixup(rbnodes[8])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 4 got %v", got)
		t.Errorf("case 4 want %v", want)
	}
	// case 5
	rbt = buildTestCase5()
	rbt.rbInsertFixup(rbnodes[4])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 5 got %v", got)
		t.Errorf("case 5 want %v", want)
	}
	// case 6
	rbt = buildTestCase6()
	rbt.rbInsertFixup(rbnodes[5])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 6 got %v", got)
		t.Errorf("case 6 want %v", want)
	}
}
