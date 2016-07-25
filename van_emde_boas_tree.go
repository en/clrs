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
	predCluster := v.summary.predecessor(v.high(x))
	if predCluster == -1 {
		return -1
	}
	offset = v.cluster[predCluster].maximum()
	return v.index(predCluster, offset)
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

func (v vEB) successor(x int) int {
	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		}
		return -1
	} else if v.min != -1 && x < v.min {
		return v.min
	} else {
		maxLow := v.cluster[v.high(x)].maximum()
		if maxLow != -1 && v.low(x) < maxLow {
			offset := v.cluster[v.high(x)].successor(v.low(x))
			return v.index(v.high(x), offset)
		}
		succCluster := v.summary.successor(v.high(x))
		if succCluster == -1 {
			return -1
		}
		offset := v.cluster[succCluster].minimum()
		return v.index(succCluster, offset)
	}
}

func (v vEB) predecessor(x int) int {
	if v.u == 2 {
		if x == 1 && v.min == 0 {
			return 0
		}
		return -1
	} else if v.max != -1 && x > v.max {
		return v.max
	} else {
		minLow := v.cluster[v.high(x)].minimum()
		if minLow != -1 && v.low(x) > minLow {
			offset := v.cluster[v.high(x)].predecessor(v.low(x))
			return v.index(v.high(x), offset)
		}
		predCluster := v.summary.predecessor(v.high(x))
		if predCluster == -1 {
			if v.min != -1 && x > v.min {
				return v.min
			}
			return -1
		}
		offset := v.cluster[predCluster].maximum()
		return v.index(predCluster, offset)
	}
}

func (v *vEB) emptyInsert(x int) {
	v.min = x
	v.max = x
}

func (v *vEB) insert(x int) {
	if v.min == -1 {
		v.emptyInsert(x)
	} else {
		if x < v.min {
			x, v.min = v.min, x
		}
		if v.u > 2 {
			if v.cluster[v.high(x)].minimum() == -1 {
				v.summary.insert(v.high(x))
				v.cluster[v.high(x)].emptyInsert(v.low(x))
			} else {
				v.cluster[v.high(x)].insert(v.low(x))
			}
		}
		if x > v.max {
			v.max = x
		}
	}
}

func (v *vEB) delete(x int) {
	if v.min == v.max {
		v.min = -1
		v.max = -1
	} else if v.u == 2 {
		if x == 0 {
			v.min = 1
		} else {
			v.min = 0
		}
		v.max = v.min
	} else {
		if x == v.min {
			firstCluster := v.summary.minimum()
			x := v.index(firstCluster, v.cluster[firstCluster].minimum())
			v.min = x
		}
		v.cluster[v.high(x)].delete(v.low(x))
		if v.cluster[v.high(x)].minimum() == -1 {
			v.summary.delete(v.high(x))
			if x == v.max {
				summaryMax := v.summary.maximum()
				if summaryMax == -1 {
					v.max = v.min
				} else {
					v.max = v.index(summaryMax, v.cluster[summaryMax].maximum())
				}
			}
		} else if x == v.max {
			v.max = v.index(v.high(x), v.cluster[v.high(x)].maximum())
		}
	}
}
