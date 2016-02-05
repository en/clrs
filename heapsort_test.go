package clrs

import (
	"reflect"
	"testing"
)

func TestHeapsort(t *testing.T) {
	for _, c := range sortCases {
		in := &heap{}
		in.a = make([]int, len(c.in))
		copy(in.a, c.in)
		heapsort(in)
		if !reflect.DeepEqual(in.a, c.want) {
			t.Errorf("sorted %v", c.in)
			t.Errorf("   got %v", in.a)
			t.Errorf("  want %v", c.want)
		}
	}
}
