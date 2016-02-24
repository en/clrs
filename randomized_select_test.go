package clrs

import (
	"testing"
)

func TestRandomizedSelect(t *testing.T) {
	cases := []struct {
		a    []int
		i    int
		want int
	}{
		{[]int{31, 41, 59, 26, 41, 58}, 1, 26},
		{[]int{31, 41, 59, 26, 41, 58}, 2, 31},
		{[]int{31, 41, 59, 26, 41, 58}, 3, 41},
		{[]int{31, 41, 59, 26, 41, 58}, 4, 41},
		{[]int{31, 41, 59, 26, 41, 58}, 5, 58},
		{[]int{31, 41, 59, 26, 41, 58}, 6, 59},
	}
	for _, c := range cases {
		got := randomizedSelect(c.a, 0, len(c.a)-1, c.i)
		if got != c.want {
			t.Errorf("search the %vth smallest element of", c.i)
			t.Errorf("     %v", c.a)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}
