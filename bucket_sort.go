package clrs

import (
	"math"
)

func insertionSortF(a []float64) {
	for j := 1; j <= len(a)-1; j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
}

func bucketSort(a []float64) {
	n := len(a)
	// TODO: replace b with an array of linked list
	b := make([][]float64, n)
	for i := 0; i < n; i++ {
		b[i] = make([]float64, 0)
	}
	for i := 0; i < n; i++ {
		index := int(math.Floor(float64(n) * a[i]))
		b[index] = append(b[index], a[i])
	}
	for i := 0; i < n; i++ {
		insertionSortF(b[i])
	}
	i := 0
	for j := 0; j < n; j++ {
		for _, v := range b[j] {
			a[i] = v
			i = i + 1
		}
	}
}
