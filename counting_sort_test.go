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

func TestCountingSortOnDigit(t *testing.T) {
	a := []int{329, 457, 657, 839, 436, 720, 355}
	k := 9
	cases := []struct {
		d    int
		want []int
	}{
		{1, []int{720, 355, 436, 457, 657, 329, 839}},
		{2, []int{329, 720, 839, 436, 457, 657, 355}},
		{3, []int{329, 355, 457, 436, 657, 720, 839}},
	}

	for _, c := range cases {
		b := make([]int, len(a))
		countingSortOnDigit(a, b, k, c.d)
		if !reflect.DeepEqual(b, c.want) {
			t.Errorf("sorted %v on digit %v", a, c.d)
			t.Errorf("   got %v", b)
			t.Errorf("  want %v", c.want)
		}
	}
}
