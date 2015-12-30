package clrs

func iterativeLinearSearch(a []int, v int) int {
	for i := 0; i <= len(a)-1; i++ {
		if a[i] == v {
			return i
		}
	}
	return -1
}

func recursiveLinearSearch(a []int, v, n int) int {
	if a[n] == v {
		return n
	}
	if n == 0 {
		return -1
	} else {
		return recursiveLinearSearch(a, v, n-1)
	}
}
