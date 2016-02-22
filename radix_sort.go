package clrs

func radixSort(a []int, d int) {
	for i := 1; i <= d; i++ {
		b := make([]int, len(a))
		countingSortOnDigit(a, b, 9, i)
		copy(a, b)
	}
}
