package clrs

import (
	"testing"
)

func TestIterativeBinarySearch(t *testing.T) {
	for _, c := range searchCases {
		got := iterativeBinarySearch(c.a, c.v, 0, len(c.a)-1)
		if got != c.want {
			t.Errorf("search for %v in %v", c.v, c.a)
			t.Errorf("got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}

func TestRecursiveBinarySearch(t *testing.T) {
	for _, c := range searchCases {
		got := recursiveBinarySearch(c.a, c.v, 0, len(c.a)-1)
		if got != c.want {
			t.Errorf("search for %v in %v", c.v, c.a)
			t.Errorf("got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}
