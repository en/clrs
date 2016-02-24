package clrs

func findMinimumAndMaximum(a []int) (int, int) {
	var min, max, start int
	if len(a)%2 != 0 {
		start = 1
		min = 0
		max = 0
	} else {
		start = 2
		if a[0] > a[1] {
			max = 0
			min = 1
		} else {
			max = 1
			min = 0
		}
	}

	for i := start; i <= len(a)-1; i += 2 {
		var smaller, larger int
		if a[i] > a[i+1] {
			larger = i
			smaller = i + 1
		} else {
			larger = i + 1
			smaller = i
		}
		if a[larger] > a[max] {
			max = larger
		}
		if a[smaller] < a[min] {
			min = smaller
		}
	}
	return min, max
}
