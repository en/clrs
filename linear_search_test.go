package clrs

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	for _, c := range searchCases {
		got := linearSearch(c.a, c.v)
		if got != c.want {
			t.Errorf("search for %v in %v", c.v, c.a)
			t.Errorf("got %v", got)
			t.Errorf("want %v", c.want)
		}
	}
}
