package clrs

import (
	"math/rand"
)

func quicksort(a []int, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quicksort(a, p, q-1)
		quicksort(a, q+1, r)
	}
}

func partition(a []int, p, r int) int {
	x := a[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if a[j] <= x {
			i = i + 1
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

func randomizedPartition(a []int, p, r int) int {
	i := rand.Intn(r-p+1) + p
	a[r], a[i] = a[i], a[r]
	return partition(a, p, r)
}

func randomizedQuicksort(a []int, p, r int) {
	if p < r {
		q := randomizedPartition(a, p, r)
		randomizedQuicksort(a, p, q-1)
		randomizedQuicksort(a, q+1, r)
	}
}

func hoarePartition(a []int, p, r int) int {
	x := a[p]
	i := p - 1
	j := r + 1
	for {
		for {
			j = j - 1
			if a[j] <= x {
				break
			}
		}
		for {
			i = i + 1
			if a[i] >= x {
				break
			}
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
		} else {
			return j
		}
	}
}

func hoareQuicksort(a []int, p, r int) {
	if p <= r-1 {
		q := hoarePartition(a, p, r)
		hoareQuicksort(a, p, q)
		hoareQuicksort(a, q+1, r)
	}
}

func tailRecursiveQuicksort(a []int, p, r int) {
	for p < r {
		q := partition(a, p, r)
		tailRecursiveQuicksort(a, p, q-1)
		p = q + 1
	}
}
