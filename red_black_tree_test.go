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

func buildDeleteTestCase1() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.7 (a)
	for _, v := range []int{1, 0, 3, 2, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = BLACK

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = nilNode
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[3]
	rbnodes[2].right = rbnodes[4]
	rbnodes[2].color = RED

	rbnodes[3].p = rbnodes[2]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = BLACK

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode
	rbnodes[4].color = BLACK

	rbt.root = rbnodes[0]
	return rbt
}

func buildDeleteTestCase2() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.7 (b)
	for _, v := range []int{1, 0, 3, 2, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = RED

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = nilNode
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[3]
	rbnodes[2].right = rbnodes[4]
	rbnodes[2].color = BLACK

	rbnodes[3].p = rbnodes[2]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = BLACK

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode
	rbnodes[4].color = BLACK

	rbt.root = rbnodes[0]
	return rbt
}

func buildDeleteTestCase3() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.7 (c)
	for _, v := range []int{1, 0, 3, 2, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = RED

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = nilNode
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[3]
	rbnodes[2].right = rbnodes[4]
	rbnodes[2].color = BLACK

	rbnodes[3].p = rbnodes[2]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = RED

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode
	rbnodes[4].color = BLACK

	rbt.root = rbnodes[0]
	return rbt
}

func buildDeleteTestCase4() *rbtree {
	rbt := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 13.7 (d)
	for _, v := range []int{1, 0, 3, 2, 4} {
		n := new(rbnode)
		n.key = v
		rbnodes = append(rbnodes, n)
	}
	rbnodes[0].p = nilNode
	rbnodes[0].left = rbnodes[1]
	rbnodes[0].right = rbnodes[2]
	rbnodes[0].color = RED

	rbnodes[1].p = rbnodes[0]
	rbnodes[1].left = nilNode
	rbnodes[1].right = nilNode
	rbnodes[1].color = BLACK

	rbnodes[2].p = rbnodes[0]
	rbnodes[2].left = rbnodes[3]
	rbnodes[2].right = rbnodes[4]
	rbnodes[2].color = BLACK

	rbnodes[3].p = rbnodes[2]
	rbnodes[3].left = nilNode
	rbnodes[3].right = nilNode
	rbnodes[3].color = RED

	rbnodes[4].p = rbnodes[2]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode
	rbnodes[4].color = RED

	rbt.root = rbnodes[0]
	return rbt
}

func buildTestOsTree() *rbtree {
	ost := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 14.1
	data := []struct {
		key   int
		size  int
		color bool
	}{
		{26, 20, BLACK},
		{17, 12, RED}, {41, 7, BLACK},
		{14, 7, BLACK}, {21, 4, BLACK}, {30, 5, RED}, {47, 1, BLACK},
		{10, 4, RED}, {16, 2, BLACK}, {19, 2, BLACK}, {21, 1, BLACK}, {28, 1, BLACK}, {38, 3, BLACK},
		{7, 2, BLACK}, {12, 1, BLACK}, {14, 1, RED}, {20, 1, RED}, {35, 1, RED}, {39, 1, RED},
		{3, 1, RED},
	}
	for _, v := range data {
		n := new(rbnode)
		n.key = v.key
		n.size = v.size
		n.color = v.color
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
	rbnodes[3].right = rbnodes[8]

	rbnodes[4].p = rbnodes[1]
	rbnodes[4].left = rbnodes[9]
	rbnodes[4].right = rbnodes[10]

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = rbnodes[11]
	rbnodes[5].right = rbnodes[12]

	rbnodes[6].p = rbnodes[2]
	rbnodes[6].left = nilNode
	rbnodes[6].right = nilNode

	rbnodes[7].p = rbnodes[3]
	rbnodes[7].left = rbnodes[13]
	rbnodes[7].right = rbnodes[14]

	rbnodes[8].p = rbnodes[3]
	rbnodes[8].left = rbnodes[15]
	rbnodes[8].right = nilNode

	rbnodes[9].p = rbnodes[4]
	rbnodes[9].left = nilNode
	rbnodes[9].right = rbnodes[16]

	rbnodes[10].p = rbnodes[4]
	rbnodes[10].left = nilNode
	rbnodes[10].right = nilNode

	rbnodes[11].p = rbnodes[5]
	rbnodes[11].left = nilNode
	rbnodes[11].right = nilNode

	rbnodes[12].p = rbnodes[5]
	rbnodes[12].left = rbnodes[17]
	rbnodes[12].right = rbnodes[18]

	rbnodes[13].p = rbnodes[7]
	rbnodes[13].left = rbnodes[19]
	rbnodes[13].right = nilNode

	rbnodes[14].p = rbnodes[7]
	rbnodes[14].left = nilNode
	rbnodes[14].right = nilNode

	rbnodes[15].p = rbnodes[8]
	rbnodes[15].left = nilNode
	rbnodes[15].right = nilNode

	rbnodes[16].p = rbnodes[9]
	rbnodes[16].left = nilNode
	rbnodes[16].right = nilNode

	rbnodes[17].p = rbnodes[12]
	rbnodes[17].left = nilNode
	rbnodes[17].right = nilNode

	rbnodes[18].p = rbnodes[12]
	rbnodes[18].left = nilNode
	rbnodes[18].right = nilNode

	rbnodes[19].p = rbnodes[13]
	rbnodes[19].left = nilNode
	rbnodes[19].right = nilNode

	ost.root = rbnodes[0]
	return ost
}

func buildTestIntervalTree() *rbtree {
	it := new(rbtree)
	rbnodes = make([]*rbnode, 0)
	// figure 14.4
	data := []struct {
		low   int
		high  int
		max   int
		color bool
	}{
		{16, 21, 30, BLACK},
		{8, 9, 23, RED}, {25, 30, 30, RED},
		{5, 8, 10, BLACK}, {15, 23, 23, BLACK}, {17, 19, 20, BLACK}, {26, 26, 26, BLACK},
		{0, 3, 3, RED}, {6, 10, 10, RED}, {19, 20, 20, RED},
	}
	for _, v := range data {
		n := new(rbnode)
		n.i.low = v.low
		n.i.high = v.high
		n.max = v.max
		n.color = v.color
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
	rbnodes[3].right = rbnodes[8]

	rbnodes[4].p = rbnodes[1]
	rbnodes[4].left = nilNode
	rbnodes[4].right = nilNode

	rbnodes[5].p = rbnodes[2]
	rbnodes[5].left = nilNode
	rbnodes[5].right = rbnodes[9]

	rbnodes[6].p = rbnodes[2]
	rbnodes[6].left = nilNode
	rbnodes[6].right = nilNode

	rbnodes[7].p = rbnodes[3]
	rbnodes[7].left = nilNode
	rbnodes[7].right = nilNode

	rbnodes[8].p = rbnodes[3]
	rbnodes[8].left = nilNode
	rbnodes[8].right = nilNode

	rbnodes[9].p = rbnodes[5]
	rbnodes[9].left = nilNode
	rbnodes[9].right = nilNode

	it.root = rbnodes[0]
	return it
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

func TestRedBlackDeleteFixup(t *testing.T) {
	want := []rbdata{
		{3, BLACK},
		{1, BLACK}, {4, RED},
		{0, BLACK}, {2, RED}, {-1, BLACK}, {-1, BLACK},
	}
	//case 1
	rbt := buildDeleteTestCase1()
	rbt.rbDeleteFixup(rbnodes[1])
	got := rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 1 got %v", got)
		t.Errorf("case 1 want %v", want)
	}
	want = []rbdata{
		{1, BLACK},
		{0, BLACK}, {3, RED},
		{-1, BLACK}, {-1, BLACK}, {2, BLACK}, {4, BLACK},
	}
	//case 2
	rbt = buildDeleteTestCase2()
	rbt.rbDeleteFixup(rbnodes[1])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 2 got %v", got)
		t.Errorf("case 2 want %v", want)
	}
	want = []rbdata{
		{2, BLACK},
		{1, BLACK}, {3, BLACK},
		{0, BLACK}, {-1, BLACK}, {-1, BLACK}, {4, BLACK},
	}
	//case 3
	rbt = buildDeleteTestCase3()
	rbt.rbDeleteFixup(rbnodes[1])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 3 got %v", got)
		t.Errorf("case 3 want %v", want)
	}
	want = []rbdata{
		{2, BLACK},
		{1, BLACK}, {3, BLACK},
		{0, BLACK}, {-1, BLACK}, {-1, BLACK}, {4, BLACK},
	}
	//case 4
	rbt = buildDeleteTestCase3()
	rbt.rbDeleteFixup(rbnodes[1])
	got = rbt.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" case 4 got %v", got)
		t.Errorf("case 4 want %v", want)
	}
}

func TestOsSelect(t *testing.T) {
	ost := buildTestOsTree()
	want := rbnodes[12]
	got := ost.osSelect(ost.root, 17)
	if got != want {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestOsRank(t *testing.T) {
	ost := buildTestOsTree()
	want := 17
	got := ost.osRank(rbnodes[12])
	if got != want {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestIntervalSearch(t *testing.T) {
	it := buildTestIntervalTree()
	want := rbnodes[4]
	x := interval{22, 25}
	got := it.intervalSearch(x)
	if got != want {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	// reset
	it = buildTestIntervalTree()
	want = nilNode
	x = interval{11, 14}
	got = it.intervalSearch(x)
	if got != want {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
