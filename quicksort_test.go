package clrs

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		quicksort(in, 0, len(in)-1)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}

func TestRandomizedQuickSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		randomizedQuicksort(in, 0, len(in)-1)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}

func TestHoareQuickSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		hoareQuicksort(in, 0, len(in)-1)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}

func TestTailRecursiveQuickSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		tailRecursiveQuicksort(in, 0, len(in)-1)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}
