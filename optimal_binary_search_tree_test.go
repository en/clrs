package clrs

import (
	"reflect"
	"testing"
)

var (
	optimalBstP        = []float32{0.15, 0.10, 0.05, 0.10, 0.20}
	optimalBstQ        = []float32{0.05, 0.10, 0.05, 0.05, 0.05, 0.10}
	wantOptimalBstRoot = [][]int{
		{1, 1, 2, 2, 2},
		{0, 2, 2, 2, 4},
		{0, 0, 3, 4, 5},
		{0, 0, 0, 4, 5},
		{0, 0, 0, 0, 5},
	}
	wantOptimalBstOut = []string{
		"k2 is the root",
		"k1 is the left child of k2",
		"d0 is the left child of k1",
		"d1 is the right child of k1",
		"k5 is the right child of k2",
		"k4 is the left child of k5",
		"k3 is the left child of k4",
		"d2 is the left child of k3",
		"d3 is the right child of k3",
		"d4 is the right child of k4",
		"d5 is the right child of k5",
	}
)

func TestOptimalBst(t *testing.T) {
	_, root := optimalBst(optimalBstP, optimalBstQ, len(optimalBstP))
	if !reflect.DeepEqual(root, wantOptimalBstRoot) {
		t.Errorf(" got %v", root)
		t.Errorf("want %v", wantOptimalBstRoot)
	}
	out := constructOptimalBst(root)
	if !reflect.DeepEqual(out, wantOptimalBstOut) {
		t.Errorf(" got %v", out)
		t.Errorf("want %v", wantOptimalBstOut)
	}
}
