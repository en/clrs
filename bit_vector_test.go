package clrs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBitVector(t *testing.T) {
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

	var b bitVector
	if b != 0x0 {
		t.Errorf(" got %v", fmt.Sprintf("%b", b))
		t.Errorf("want %v", fmt.Sprintf("%b", 0x0))
	}
	for i := 0; i < 16; i++ {
		err := b.Insert(i)
		if err != nil {
			t.Errorf("insert %v", i)
			t.Errorf("   err %v", err)
		}
	}
	if b != 0xffff {
		t.Errorf(" got %v", fmt.Sprintf("%b", b))
		t.Errorf("want %v", fmt.Sprintf("%b", 0xffff))
	}
	for i := 1; i < 16; i += 2 {
		b.Delete(i)
		if i == 7 {
			i = 6
		}
	}
	if b != 0xaa55 {
		t.Errorf(" got %v", fmt.Sprintf("%b", b))
		t.Errorf("want %v", fmt.Sprintf("%b", 0xaa55))
	}
}
