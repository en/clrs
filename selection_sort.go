package clrs

func selectionSort(a []int) {
	for j := 0; j <= len(a)-2; j++ {
		smallest := j
		for i := j + 1; i <= len(a)-1; i++ {
			if a[i] < a[smallest] {
				smallest = i
			}
		}
		a[j], a[smallest] = a[smallest], a[j]
	}
}
