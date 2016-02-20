package clrs

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	cases := []struct {
		a    []int
		k    int
		want []int
	}{
		{[]int{2, 5, 3, 0, 2, 3, 0, 3}, 5, []int{0, 0, 2, 2, 3, 3, 3, 5}},
		{[]int{6, 0, 2, 0, 1, 3, 4, 6, 1, 3, 2}, 6, []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 6, 6}},
	}

	for _, c := range cases {
		b := make([]int, len(c.a))
		countingSort(c.a, b, c.k)
		if !reflect.DeepEqual(b, c.want) {
			t.Errorf("sorted %v", c.a)
			t.Errorf("   got %v", b)
			t.Errorf("  want %v", c.want)
		}
	}
}
