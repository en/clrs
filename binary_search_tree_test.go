package clrs

import (
	"reflect"
	"testing"
)

var (
	nodes []*BSTNode
)

func buildTestBST() *BST {
	bst := new(BST)
	nodes = make([]*BSTNode, 0)
	// figure 12.2
	for _, v := range []int{15, 6, 18, 3, 7, 17, 20, 2, 4, 13, 9} {
		n := new(BSTNode)
		n.key = v
		nodes = append(nodes, n)
	}
	nodes[0].left = nodes[1]
	nodes[0].right = nodes[2]

	nodes[1].p = nodes[0]
	nodes[1].left = nodes[3]
	nodes[1].right = nodes[4]

	nodes[2].p = nodes[0]
	nodes[2].left = nodes[5]
	nodes[2].right = nodes[6]

	nodes[3].p = nodes[1]
	nodes[3].left = nodes[7]
	nodes[3].right = nodes[8]

	nodes[4].p = nodes[1]
	nodes[4].right = nodes[9]

	nodes[5].p = nodes[2]

	nodes[6].p = nodes[2]

	nodes[7].p = nodes[3]

	nodes[8].p = nodes[3]

	nodes[9].p = nodes[4]
	nodes[9].left = nodes[10]

	nodes[10].p = nodes[9]

	bst.root = nodes[0]
	return bst
}

func TestBSTToArray(t *testing.T) {
	bst := buildTestBST()
	want := []int{
		15,
		6, 18,
		3, 7, 17, 20,
		2, 4, -1, 13, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}
	got := bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestBSTIterativePreorderTreeWalk(t *testing.T) {
	bst := buildTestBST()
	want := []int{15, 6, 3, 2, 4, 7, 13, 9, 18, 17, 20}
	i := 0
	for key := range bst.iterativePreorderTreeWalkIter() {
		if key != want[i] {
			t.Errorf(" got %v", key)
			t.Errorf("want %v", want[i])
		}
		i = i + 1
	}
}

func TestBSTConstantExtraSpaceTreeWalk(t *testing.T) {
	bst := buildTestBST()
	want := []int{15, 6, 3, 2, 4, 7, 13, 9, 18, 17, 20}
	i := 0
	for key := range bst.constantExtraSpaceTreeWalkIter() {
		if key != want[i] {
			t.Errorf(" got %v", key)
			t.Errorf("want %v", want[i])
		}
		i = i + 1
	}
}

func TestBSTTreeSearch(t *testing.T) {
	bst := buildTestBST()
	if n := bst.treeSearch(bst.root, 3); n == nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[3])
	}
	if n := bst.treeSearch(bst.root, 19); n != nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nil)
	}
	if n := bst.treeSearch(nodes[1], 13); n == nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[9])
	}
	if n := bst.treeSearch(nodes[1], 15); n != nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nil)
	}
}

func TestBSTIterativeTreeSearch(t *testing.T) {
	bst := buildTestBST()
	if n := bst.iterativeTreeSearch(bst.root, 3); n == nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[3])
	}
	if n := bst.iterativeTreeSearch(bst.root, 19); n != nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nil)
	}
	if n := bst.iterativeTreeSearch(nodes[1], 13); n == nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[9])
	}
	if n := bst.iterativeTreeSearch(nodes[1], 15); n != nil {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nil)
	}
}

func TestBSTTreeMinimum(t *testing.T) {
	bst := buildTestBST()
	if n := bst.treeMinimum(bst.root); n != nodes[7] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[7])
	}
	if n := bst.treeMinimum(nodes[4]); n != nodes[4] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[4])
	}
}

func TestBSTTreeMaximum(t *testing.T) {
	bst := buildTestBST()
	if n := bst.treeMaximum(bst.root); n != nodes[6] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[6])
	}
	if n := bst.treeMaximum(nodes[4]); n != nodes[9] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[9])
	}
}

func TestBSTTreeSuccessor(t *testing.T) {
	bst := buildTestBST()
	if n := bst.treeSuccessor(nodes[0]); n != nodes[5] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[5])
	}
	if n := bst.treeSuccessor(nodes[9]); n != nodes[0] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[0])
	}
}

func TestBSTTreePredecessor(t *testing.T) {
	bst := buildTestBST()
	if n := bst.treePredecessor(nodes[0]); n != nodes[9] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[9])
	}
	if n := bst.treePredecessor(nodes[10]); n != nodes[4] {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", nodes[4])
	}
}

func TestBSTTreeInsert(t *testing.T) {
	bst := new(BST)
	for _, v := range []int{12, 5, 18, 2, 9, 15, 19, 17, 13} {
		n := new(BSTNode)
		n.key = v
		bst.treeInsert(n)
	}
	want := []int{
		12,
		5, 18,
		2, 9, 15, 19,
		-1, -1, -1, -1, 13, 17, -1, -1,
	}
	got := bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestBSTTreeDelete(t *testing.T) {
	bst := buildTestBST()
	bst.treeDelete(nodes[8])
	want := []int{
		15,
		6, 18,
		3, 7, 17, 20,
		2, -1, -1, 13, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}
	got := bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	bst = buildTestBST()
	bst.treeDelete(nodes[9])
	want = []int{
		15,
		6, 18,
		3, 7, 17, 20,
		2, 4, -1, 9, -1, -1, -1, -1,
	}
	got = bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	bst = buildTestBST()
	bst.treeDelete(nodes[4])
	want = []int{
		15,
		6, 18,
		3, 13, 17, 20,
		2, 4, 9, -1, -1, -1, -1, -1,
	}
	got = bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	bst = buildTestBST()
	bst.treeDelete(nodes[2])
	want = []int{
		15,
		6, 20,
		3, 7, 17, -1,
		2, 4, -1, 13, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}
	got = bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
	bst = buildTestBST()
	bst.treeDelete(nodes[0])
	want = []int{
		17,
		6, 18,
		3, 7, -1, 20,
		2, 4, -1, 13, -1, -1, -1, -1,
		-1, -1, -1, -1, -1, -1, 9, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	}
	got = bst.toArray()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
