package clrs

var (
	sortCases = []struct {
		in, want []int
	}{
		{[]int{0}, []int{0}},
		{[]int{31, 41, 59, 26, 41, 58}, []int{26, 31, 41, 41, 58, 59}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}, []int{-5467984, -784, 0, 0, 42, 59, 74, 238, 905, 959, 7586, 7586, 9845}},
	}
	searchCases = []struct {
		a    []int
		v    int
		want int
	}{
		{[]int{0}, 5, -1},
		{[]int{1, 2, 3, 4, 5}, 0, -1},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 2, 1},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 4, 3},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
	}
)
