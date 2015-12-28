package clrs

import (
	"reflect"
	"testing"
)

func TestBinaryAddition(t *testing.T) {
	cases := []struct {
		a    []byte
		b    []byte
		want []byte
	}{
		{[]byte{0, 1, 1, 1, 0, 0, 0}, []byte{1, 0, 0, 1, 1, 0, 1}, []byte{1, 0, 0, 0, 0, 1, 0, 1}},
		{[]byte{1, 0, 1, 1, 0, 0, 0}, []byte{1, 0, 0, 1, 1, 0, 1}, []byte{1, 0, 1, 0, 0, 1, 0, 1}},
	}
	for _, c := range cases {
		s := make([]byte, len(c.a)+1)
		binaryAddition(c.a, c.b, s)
		if !reflect.DeepEqual(s, c.want) {
			t.Errorf("adding %v", c.a)
			t.Errorf("   and %v", c.b)
			t.Errorf("   got %v", s)
			t.Errorf("  want %v", c.want)
		}
	}
}
