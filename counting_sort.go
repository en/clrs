package clrs

import (
	"math"
)

func countingSort(a, b []int, k int) {
	c := make([]int, k+1)
	for i := 0; i <= k; i++ {
		c[i] = 0
	}
	for j := 0; j < len(a); j++ {
		c[a[j]] = c[a[j]] + 1
	}
	for i := 1; i <= k; i++ {
		c[i] = c[i] + c[i-1]
	}
	for j := len(a) - 1; j >= 0; j-- {
		b[c[a[j]]-1] = a[j]
		c[a[j]] = c[a[j]] - 1
	}
}

func digit(v, d int) int {
	// 753
	//   ^
	//   digit 1
	return v / int(math.Pow(float64(10), float64(d-1))) % 10
}

func countingSortOnDigit(a, b []int, k, d int) {
	c := make([]int, k+1)
	for i := 0; i <= k; i++ {
		c[i] = 0
	}
	for j := 0; j < len(a); j++ {
		c[digit(a[j], d)] = c[digit(a[j], d)] + 1
	}
	for i := 1; i <= k; i++ {
		c[i] = c[i] + c[i-1]
	}
	for j := len(a) - 1; j >= 0; j-- {
		n := digit(a[j], d)
		b[c[n]-1] = a[j]
		c[n] = c[n] - 1
	}
}
