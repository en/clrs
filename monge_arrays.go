package clrs

func findMinimum(a [][]int, row, left, right int) int {
	min := left
	for i := left + 1; i <= right; i++ {
		if a[row][i] < a[row][min] {
			min = i
		}
	}
	return min
}

func findMinimums(a [][]int, rows []int) []int {
	if len(rows) == 1 {
		min := findMinimum(a, rows[0], 0, len(a[rows[0]])-1)
		return []int{min}
	}
	var evenRows []int
	for i, v := range rows {
		if (i+1)%2 == 0 {
			evenRows = append(evenRows, v)
		}
	}
	evenMins := findMinimums(a, evenRows)
	left := 0
	var mins []int
	for i, v := range rows {
		if (i+1)%2 == 0 {
			mins = append(mins, evenMins[(i-1)/2])
			left = evenMins[(i-1)/2]
		} else {
			var right int
			if i+1 < len(rows) {
				right = evenMins[i/2]
			} else {
				right = len(a[i]) - 1
			}
			min := findMinimum(a, v, left, right)
			mins = append(mins, min)
		}
	}
	return mins
}

func isMongeArrays(a [][]int) bool {
	n := len(a)
	m := len(a[0])
	for i := 0; i < m-1; i++ {
		for j := 0; j < n-1; j++ {
			if (a[j][i] + a[j+1][i+1]) > (a[j+1][i] + a[j][i+1]) {
				return false
			}
		}
	}
	return true
}
