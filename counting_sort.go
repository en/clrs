package clrs

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
