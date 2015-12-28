package clrs

func linearSearch(a []int, v int) int {
	for i := 0; i <= len(a)-1; i++ {
		if a[i] == v {
			return i
		}
	}
	return -1
}
