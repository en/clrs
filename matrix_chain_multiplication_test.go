package clrs

import (
	"reflect"
	"testing"
)

var (
	matrixChainP = []int{30, 35, 15, 5, 10, 20, 25}
	wantM        = [][]int{
		{0, 15750, 7875, 9375, 11875, 15125},
		{0, 0, 2625, 4375, 7125, 10500},
		{0, 0, 0, 750, 2500, 5375},
		{0, 0, 0, 0, 1000, 3500},
		{0, 0, 0, 0, 0, 5000},
		{0, 0, 0, 0, 0, 0},
	}
	wantS = [][]int{
		{0, 1, 1, 3, 3, 3},
		{0, 0, 2, 3, 3, 3},
		{0, 0, 0, 3, 3, 3},
		{0, 0, 0, 0, 4, 5},
		{0, 0, 0, 0, 0, 5},
		{0, 0, 0, 0, 0, 0},
	}
	wantOut = "((A1(A2A3))((A4A5)A6))"
)

func TestMatrixChainOrder(t *testing.T) {
	m, s := matrixChainOrder(matrixChainP)
	if !reflect.DeepEqual(m, wantM) {
		t.Errorf(" got %v", m)
		t.Errorf("want %v", wantM)
	}
	if !reflect.DeepEqual(s, wantS) {
		t.Errorf(" got %v", s)
		t.Errorf("want %v", wantS)
	}
	out := printOptimalParents(s, 1, 6)
	if !reflect.DeepEqual(out, wantOut) {
		t.Errorf(" got %v", out)
		t.Errorf("want %v", wantOut)
	}
}
