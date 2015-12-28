package clrs

import (
	"testing"
)

func TestCountInversions(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{0}, 0},
		{[]int{2, 3, 8, 6, 1}, 5},
		{[]int{6, 5, 4, 3, 2, 1}, 15},
		{[]int{1, 2, 3, 4, 5, 6}, 0},
	}
	for _, c := range cases {
		got := countInversions(c.in, 0, len(c.in)-1)
		if got != c.want {
			t.Errorf("count %v", c.in)
			t.Errorf("  got %v", got)
			t.Errorf(" want %v", c.want)
		}
	}
}
