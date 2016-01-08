package clrs

import (
	"math/rand"
	"testing"
	"time"
)

func TestChipTesting(t *testing.T) {
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 32; i++ {
		mid := i / 2
		good := rand.Intn(i-mid) + mid + 1
		c := make([]chip, i)
		for j := 0; j < good; j++ {
			c[j] = chip{good: true}
		}
		var chips []chip
		perm := rand.Perm(i)
		for i, v := range perm {
			c[v].id = i
			chips = append(chips, c[v])
		}
		if chip := diogenes(chips); !chip.good {
			t.Errorf("finding a single good chip from %v", chips)
			t.Errorf("got %v", chip)
		}
	}

	chips := []chip{
		{id: 0, good: false, conspire: true}, {id: 1, good: true}, {id: 2, good: true},
		{id: 3, good: false, conspire: true}, {id: 4, good: true}, {id: 5, good: true},
		{id: 6, good: false},
	}
	if chip := diogenes(chips); !chip.good {
		t.Errorf("finding a single good chip from %v", chips)
		t.Errorf("got %v", chip)
	}
}
