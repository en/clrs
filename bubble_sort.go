package clrs

func bubbleSort(a []int) {
	for i := 0; i <= len(a)-2; i++ {
		for j := len(a) - 1; j >= i+1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
