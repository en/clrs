package clrs

import (
	"testing"
)

func TestIterativeLinearSearch(t *testing.T) {
	for _, c := range searchCases {
		got := iterativeLinearSearch(c.a, c.v)
		if got != c.want {
			t.Errorf("search for %v in %v", c.v, c.a)
			t.Errorf("got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}

func TestRecursiveLinearSearch(t *testing.T) {
	for _, c := range searchCases {
		got := recursiveLinearSearch(c.a, c.v, len(c.a)-1)
		if got != c.want {
			t.Errorf("search for %v in %v", c.v, c.a)
			t.Errorf("got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}
