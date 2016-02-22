package clrs

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	a := []int{329, 457, 657, 839, 436, 720, 355}
	want := []int{329, 355, 436, 457, 657, 720, 839}
	in := make([]int, len(a))
	copy(in, a)
	radixSort(in, 3)
	if !reflect.DeepEqual(in, want) {
		t.Errorf("sorted %v", a)
		t.Errorf("   got %v", in)
		t.Errorf("  want %v", want)
	}
}
