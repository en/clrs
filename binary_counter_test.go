package clrs

import (
	"reflect"
	"testing"
)

func TestIncrement(t *testing.T) {
	want := [][]byte{
		[]byte{1, 0, 0, 0, 0, 0, 0, 0},
		[]byte{0, 1, 0, 0, 0, 0, 0, 0},
		[]byte{1, 1, 0, 0, 0, 0, 0, 0},
		[]byte{0, 0, 1, 0, 0, 0, 0, 0},
		[]byte{1, 0, 1, 0, 0, 0, 0, 0},
		[]byte{0, 1, 1, 0, 0, 0, 0, 0},
		[]byte{1, 1, 1, 0, 0, 0, 0, 0},
		[]byte{0, 0, 0, 1, 0, 0, 0, 0},
		[]byte{1, 0, 0, 1, 0, 0, 0, 0},
		[]byte{0, 1, 0, 1, 0, 0, 0, 0},
		[]byte{1, 1, 0, 1, 0, 0, 0, 0},
		[]byte{0, 0, 1, 1, 0, 0, 0, 0},
		[]byte{1, 0, 1, 1, 0, 0, 0, 0},
		[]byte{0, 1, 1, 1, 0, 0, 0, 0},
		[]byte{1, 1, 1, 1, 0, 0, 0, 0},
		[]byte{0, 0, 0, 0, 1, 0, 0, 0},
	}
	a := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for _, w := range want {
		increment(a)
		if !reflect.DeepEqual(a, w) {
			t.Errorf(" got %v", a)
			t.Errorf("want %v", w)
		}
	}
}
