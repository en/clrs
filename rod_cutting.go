package clrs

import (
	"math"
)

func cutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1
	for i := 1; i <= n; i++ {
		q = int(math.Max(float64(q), float64(p[i-1]+cutRod(p, n-i))))
	}
	return q
}

func memoizedCutRod(p []int, n int) int {
	r := make(map[int]int)
	return memoizedCutRodAux(p, n, r)
}

func memoizedCutRodAux(p []int, n int, r map[int]int) int {
	if v, ok := r[n]; ok {
		return v
	}
	var q int
	if n == 0 {
		q = 0
	} else {
		q = -1
		for i := 1; i <= n; i++ {
			q = int(math.Max(float64(q), float64(p[i-1]+memoizedCutRodAux(p, n-i, r))))
		}
	}
	r[n] = q
	return q
}

func bottomUpCutRod(p []int, n int) int {
	r := make(map[int]int)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			q = int(math.Max(float64(q), float64(p[i-1]+r[j-i])))
		}
		r[j] = q
	}
	return r[n]
}

func extendedBottomUpCutRod(p []int, n int) (int, map[int]int) {
	r := make(map[int]int)
	s := make(map[int]int)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := -1
		for i := 1; i <= j; i++ {
			if q < p[i-1]+r[j-i] {
				q = p[i-1] + r[j-i]
				s[j] = i
			}
		}
		r[j] = q
	}
	return r[n], s
}
