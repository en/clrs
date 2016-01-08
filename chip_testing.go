package clrs

import (
	"math/rand"
)

type chip struct {
	id       int
	good     bool
	conspire bool
}

func (c *chip) report(other chip) bool {
	if c.good {
		return other.good
	}
	if c.conspire && !other.good {
		return true
	}
	return []bool{true, false}[rand.Intn(2)]
}

func jig(a, b chip) (bool, bool) {
	return a.report(b), b.report(a)
}

func diogenes(chips []chip) chip {
	n := len(chips)
	if n <= 2 {
		return chips[0]
	}
	mid := n / 2
	var newChips []chip
	for i := 0; i <= mid-1; i++ {
		a := chips[i]
		b := chips[mid+i]
		if x, y := jig(a, b); x && y {
			newChips = append(newChips, []chip{a, b}[rand.Intn(2)])
		}
	}
	if n%2 != 0 && len(newChips)%2 == 0 {
		newChips = append(newChips, chips[n-1])
	}
	return diogenes(newChips)
}
