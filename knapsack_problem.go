package clrs

func dynamic01Knapsack(v, w []int, n, m int) [][]int {
	c := make([][]int, n+1)
	for i := range c {
		c[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if w[i-1] <= j {
				if v[i-1]+c[i-1][j-w[i-1]] > c[i-1][j] {
					c[i][j] = v[i-1] + c[i-1][j-w[i-1]]
				} else {
					c[i][j] = c[i-1][j]
				}
			} else {
				c[i][j] = c[i-1][j]
			}
		}
	}
	return c
}

func printKnapsackSolution(c [][]int, w []int, _m int) []int {
	n := len(w)
	var (
		out []int
		f   func(int, int)
	)
	f = func(i, m int) {
		if i == 0 || m == 0 {
			return
		}
		if c[i][m] == c[i-1][m] {
			f(i-1, m)
		} else {
			out = append(out, i)
			f(i-1, m-w[i-1])
		}
	}
	f(n, _m)
	return out
}
