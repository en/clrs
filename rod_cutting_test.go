package clrs

import (
	"testing"
)

var (
	rodPrice        = []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	rodCuttingCases = []struct {
		i int
		r int
		s int
	}{
		{0, 0, 0},
		{1, 1, 1},
		{2, 5, 2},
		{3, 8, 3},
		{4, 10, 2},
		{5, 13, 2},
		{6, 17, 6},
		{7, 18, 1},
		{8, 22, 2},
		{9, 25, 3},
		{10, 30, 10},
	}
)

func TestCutRod(t *testing.T) {
	for _, c := range rodCuttingCases {
		got := cutRod(rodPrice, c.i)
		if got != c.r {
			t.Errorf("cutting up a rod of %v inches", c.i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", c.r)
		}
	}
}

func TestMemoizedCutRod(t *testing.T) {
	for _, c := range rodCuttingCases {
		got := memoizedCutRod(rodPrice, c.i)
		if got != c.r {
			t.Errorf("cutting up a rod of %v inches", c.i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", c.r)
		}
	}
}

func TestBottomUpCutRod(t *testing.T) {
	for _, c := range rodCuttingCases {
		got := bottomUpCutRod(rodPrice, c.i)
		if got != c.r {
			t.Errorf("cutting up a rod of %v inches", c.i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", c.r)
		}
	}
}

func TestExtendedBottomUpCutRod(t *testing.T) {
	for _, c := range rodCuttingCases {
		r, s := extendedBottomUpCutRod(rodPrice, c.i)
		if r != c.r {
			t.Errorf("cutting up a rod of %v inches", c.i)
			t.Errorf(" got %v", r)
			t.Errorf("want %v", c.r)
		}
		n := c.i
		for n > 0 {
			if s[n] != rodCuttingCases[n].s {
				t.Errorf(" got %v", s[n])
				t.Errorf("want %v", rodCuttingCases[n].s)
			}
			n = n - s[n]
		}
	}
}
