package clrs

import (
	"reflect"
	"testing"
)

func TestSameComponent(t *testing.T) {
	g := new(component)
	g.v = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	g.e = []struct {
		e1, e2 rune
	}{
		{'b', 'd'},
		{'e', 'g'},
		{'a', 'c'},
		{'h', 'i'},
		{'a', 'b'},
		{'e', 'f'},
		{'b', 'c'},
	}
	g.connectedComponents()
	want := []map[rune]bool{
		map[rune]bool{'j': true},
		map[rune]bool{'h': true, 'i': true},
		map[rune]bool{'a': true, 'b': true, 'c': true, 'd': true},
		map[rune]bool{'e': true, 'f': true, 'g': true},
	}
	if !reflect.DeepEqual(g.s, want) {
		t.Errorf(" got %v", g.s)
		t.Errorf("want %v", want)
	}
	cases := []struct {
		u, v rune
		want bool
	}{
		{'a', 'd', true},
		{'a', 'e', false},
	}
	for _, c := range cases {
		got := g.sameComponent(c.u, c.v)
		if got != c.want {
			t.Errorf("same-component(%c, %c)", c.u, c.v)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}
