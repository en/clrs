package clrs

import (
	"math"
)

type protoVEB struct {
	u       int
	a       []int
	summary *protoVEB
	cluster []*protoVEB
}

func (v protoVEB) high(x int) int {
	return int(math.Floor(float64(x) / math.Sqrt(float64(v.u))))
}

func (v protoVEB) low(x int) int {
	return x % int(math.Sqrt(float64(v.u)))
}

func (v protoVEB) index(x, y int) int {
	return x*int(math.Sqrt(float64(v.u))) + y
}

func (v protoVEB) member(x int) int {
	if v.u == 2 {
		return v.a[x]
	}
	return v.cluster[v.high(x)].member(v.low(x))
}
