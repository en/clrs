package clrs

import (
	"testing"
)

func TestFindMinimumAndMaximum(t *testing.T) {
	cases := []struct {
		a        []int
		min, max int
	}{
		{[]int{31, 41, 59, 26, 41, 58}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 0, 4},
		{[]int{6, 5, 4, 3, 2, 1}, 5, 0},
		{[]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}, 11, 4},
	}
	for _, c := range cases {
		min, max := findMinimumAndMaximum(c.a)
		if min != c.min || max != c.max {
			t.Errorf("  search %v", c.a)
			t.Errorf(" got min %v", min)
			t.Errorf("want min %v", c.min)
			t.Errorf(" got max %v", max)
			t.Errorf("want max %v", c.max)

		}
	}
}
