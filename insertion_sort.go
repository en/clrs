package clrs

func iterativeInsertionSort(a []int) {
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

func recursiveInsertionSort(a []int, n int) {
	if n > 0 {
		recursiveInsertionSort(a, n-1)
		key := a[n]
		i := n - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
}
