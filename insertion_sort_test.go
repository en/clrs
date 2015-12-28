package clrs

import (
	"reflect"
	"testing"
)

func TestIterativeInsertionSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		iterativeInsertionSort(in)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}

func TestRecursiveInsertionSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		recursiveInsertionSort(in, len(in)-1)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}
