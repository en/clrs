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
	if v.u == 2 {
		if v.a[1] == 1 {
			return 1
		} else if v.a[0] == 1 {
			return 0
		}
		return -1
	}
	maxCluster := v.summary.maximum()
	if maxCluster == -1 {
		return -1
	}
	offset := v.cluster[maxCluster].maximum()
	return v.index(maxCluster, offset)
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
	if v.u == 2 {
		if x == 1 && v.a[0] == 1 {
			return 0
		}
		return -1
	}
	offset := v.cluster[v.high(x)].predecessor(v.low(x))
	if offset != -1 {
		return v.index(v.high(x), offset)
	}
	predecCluster := v.summary.predecessor(v.high(x))
	if predecCluster == -1 {
		return -1
	}
	offset = v.cluster[predecCluster].maximum()
	return v.index(predecCluster, offset)
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
	if v.u == 2 {
		v.a[x] = 0
	} else {
		h := v.high(x)
		v.cluster[h].delete(v.low(x))
		if v.cluster[h].count() == 0 {
			v.summary.delete(h)
		}
	}
}

func (v protoVEB) count() int {
	if v.u == 2 {
		return v.a[0] + v.a[1]
	}
	count := 0
	for i := 0; i < int(math.Sqrt(float64(v.u))); i++ {
		count = count + v.cluster[i].count()
	}
	return count
}

type vEB struct {
	u       int
	min     int
	max     int
	summary *vEB
	cluster []*vEB
}

func (v vEB) high(x int) int {
	lsr := math.Pow(2, math.Floor(math.Log2(float64(v.u))/2))
	return int(math.Floor(float64(x) / lsr))
}

func (v vEB) low(x int) int {
	lsr := math.Pow(2, math.Floor(math.Log2(float64(v.u))/2))
	return x % int(lsr)
}

func (v vEB) index(x, y int) int {
	lsr := math.Pow(2, math.Floor(math.Log2(float64(v.u))/2))
	return x*int(lsr) + y
}

func (v vEB) minimum() int {
	return v.min
}

func (v vEB) maximum() int {
	return v.max
}

func (v vEB) member(x int) bool {
	if x == v.min || x == v.max {
		return true
	} else if v.u == 2 {
		return false
	} else {
		return v.cluster[v.high(x)].member(v.low(x))
	}
}
