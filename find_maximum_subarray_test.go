package clrs

import (
	"testing"
)

var cases = []struct {
	a     []int
	left  int
	right int
	sum   int
}{
	{[]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, -7}, 7, 10, 43},
	{[]int{13, -3, -25, 11, 1}, 0, 0, 13},
	{[]int{-13, -3, -25, -20, -3, -16, -23, -18, -20, -7, -12, -5, -22, -15, -4, -7}, 1, 1, -3},
	{[]int{13, 3, 25, 11, 1}, 0, 4, 53},
}

func TestBruteForceFindMaximumSubarray(t *testing.T) {
	for _, c := range cases {
		left, right, sum := bruteForceFindMaximumSubarray(c.a, 0, len(c.a)-1)
		if left != c.left || right != c.right || sum != c.sum {
			t.Errorf("array %v", c.a)
			t.Errorf("  got [%v, %v] = %v", left, right, sum)
			t.Errorf(" want [%v, %v] = %v", c.left, c.right, c.sum)
		}
	}
}

func TestFindMaximumSubarray(t *testing.T) {
	for _, c := range cases {
		left, right, sum := findMaximumSubarray(c.a, 0, len(c.a)-1)
		if left != c.left || right != c.right || sum != c.sum {
			t.Errorf("array %v", c.a)
			t.Errorf("  got [%v, %v] = %v", left, right, sum)
			t.Errorf(" want [%v, %v] = %v", c.left, c.right, c.sum)
		}
	}
}

func TestLinearTimeFindMaximumSubarray(t *testing.T) {
	for _, c := range cases {
		left, right, sum := linearTimeFindMaximumSubarray(c.a, 0, len(c.a)-1)
		if left != c.left || right != c.right || sum != c.sum {
			t.Errorf("array %v", c.a)
			t.Errorf("  got [%v, %v] = %v", left, right, sum)
			t.Errorf(" want [%v, %v] = %v", c.left, c.right, c.sum)
		}
	}
}
