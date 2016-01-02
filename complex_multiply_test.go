package clrs

import (
	"testing"
)

func TestComplexMultiply(t *testing.T) {
	cases := []struct {
		a, b, c, d int
		r, i       int
	}{
		{1, 2, 3, 4, -5, 10},
		{5, -6, -7, 8, 13, 82},
	}
	for _, c := range cases {
		r, i := complexMultiply(c.a, c.b, c.c, c.d)
		if r != c.r || i != c.i {
			t.Errorf("%v*%v=?", complex(float32(c.a), float32(c.b)), complex(float32(c.c), float32(c.d)))
			t.Errorf(" got %v", complex(float32(r), float32(i)))
			t.Errorf("want %v", complex(float32(c.r), float32(c.i)))
		}
	}
}
