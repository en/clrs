package clrs

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	for _, c := range sortCases {
		in := make([]int, len(c.in))
		copy(in, c.in)
		bubbleSort(in)
		if !reflect.DeepEqual(in, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in)
			t.Errorf("  want %v", c.want)
		}
	}
}
