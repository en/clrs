package clrs

import (
	"reflect"
	"testing"
)

var (
	lcsX  = []rune{'A', 'B', 'C', 'B', 'D', 'A', 'B'}
	lcsY  = []rune{'B', 'D', 'C', 'A', 'B', 'A'}
	wantC = [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1},
		{0, 1, 1, 1, 1, 2, 2},
		{0, 1, 1, 2, 2, 2, 2},
		{0, 1, 1, 2, 2, 3, 3},
		{0, 1, 2, 2, 2, 3, 3},
		{0, 1, 2, 2, 3, 3, 4},
		{0, 1, 2, 2, 3, 4, 4},
	}
	wantB = [][]rune{
		{0, 0, 0, 0, 0, 0, 0},
		{0, '↑', '↑', '↑', '↖', '←', '↖'},
		{0, '↖', '←', '←', '↑', '↖', '←'},
		{0, '↑', '↑', '↖', '←', '↑', '↑'},
		{0, '↖', '↑', '↑', '↑', '↖', '←'},
		{0, '↑', '↖', '↑', '↑', '↑', '↑'},
		{0, '↑', '↑', '↑', '↖', '↑', '↖'},
		{0, '↖', '↑', '↑', '↑', '↖', '↑'},
	}
	wantLcsOut = []rune{'B', 'C', 'B', 'A'}
	lcsS1      = []rune{'A', 'C', 'C', 'G', 'G', 'T', 'C', 'G', 'A', 'G', 'T', 'G', 'C', 'G', 'C', 'G', 'G', 'A', 'A', 'G', 'C', 'C', 'G', 'G', 'C', 'C', 'G', 'A', 'A'}
	lcsS2      = []rune{'G', 'T', 'C', 'G', 'T', 'T', 'C', 'G', 'G', 'A', 'A', 'T', 'G', 'C', 'C', 'G', 'T', 'T', 'G', 'C', 'T', 'C', 'T', 'G', 'T', 'A', 'A', 'A'}
	lcsS3      = []rune{'G', 'T', 'C', 'G', 'T', 'C', 'G', 'G', 'A', 'A', 'G', 'C', 'C', 'G', 'G', 'C', 'C', 'G', 'A', 'A'}
)

func TestLcs(t *testing.T) {
	c, b := lcsLength(lcsX, lcsY)
	if !reflect.DeepEqual(c, wantC) {
		t.Errorf(" got %v", c)
		t.Errorf("want %v", wantC)
	}
	if !reflect.DeepEqual(b, wantB) {
		t.Errorf(" got %v", b)
		t.Errorf("want %v", wantB)
	}
	out := printLcs(b, lcsX, len(lcsX), len(lcsY))
	if !reflect.DeepEqual(out, wantLcsOut) {
		t.Errorf(" got %v", out)
		t.Errorf("want %v", wantLcsOut)
	}
}

func TestLcsS(t *testing.T) {
	_, b := lcsLength(lcsS1, lcsS2)
	out := printLcs(b, lcsS1, len(lcsS1), len(lcsS2))
	if !reflect.DeepEqual(out, lcsS3) {
		t.Errorf(" got %v", out)
		t.Errorf("want %v", lcsS3)
	}
}
