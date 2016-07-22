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

func (v protoVEB) minimum() int {
	if v.u == 2 {
		if v.a[0] == 1 {
			return 0
		} else if v.a[1] == 1 {
			return 1
		}
		return -1
	}
	minCluster := v.summary.minimum()
	if minCluster == -1 {
		return -1
	}
	offset := v.cluster[minCluster].minimum()
	return v.index(minCluster, offset)
}

func (v protoVEB) maximum() int {
	// TODO
	return 0
}

func (v protoVEB) successor(x int) int {
	if v.u == 2 {
		if x == 0 && v.a[1] == 1 {
			return 1
		}
		return -1
	}
	offset := v.cluster[v.high(x)].successor(v.low(x))
	if offset != -1 {
		return v.index(v.high(x), offset)
	}
	succCluster := v.summary.successor(v.high(x))
	if succCluster == -1 {
		return -1
	}
	offset = v.cluster[succCluster].minimum()
	return v.index(succCluster, offset)
}

func (v protoVEB) predecessor(x int) int {
	// TODO
	return 0
}

func (v *protoVEB) insert(x int) {
	if v.u == 2 {
		v.a[x] = 1
	} else {
		v.cluster[v.high(x)].insert(v.low(x))
		v.summary.insert(v.high(x))
	}
}

func (v *protoVEB) delete(x int) {
	// TODO
}
