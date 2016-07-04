package clrs

func lcsLength(x, y []rune) ([][]int, [][]rune) {
	m := len(x)
	n := len(y)
	b := make([][]rune, m+1)
	for i := range b {
		b[i] = make([]rune, n+1)
	}
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				b[i][j] = '↖'
			} else if c[i-1][j] >= c[i][j-1] {
				c[i][j] = c[i-1][j]
				b[i][j] = '↑'
			} else {
				c[i][j] = c[i][j-1]
				b[i][j] = '←'
			}
		}
	}
	return c, b
}

func printLcs(_b [][]rune, _x []rune, _i, _j int) []rune {
	var (
		out []rune
		f   func([][]rune, []rune, int, int)
	)
	f = func(b [][]rune, x []rune, i, j int) {
		if i == 0 || j == 0 {
			return
		}
		if b[i][j] == '↖' {
			f(b, x, i-1, j-1)
			out = append(out, x[i-1])
		} else if b[i][j] == '↑' {
			f(b, x, i-1, j)
		} else {
			f(b, x, i, j-1)
		}
	}
	f(_b, _x, _i, _j)
	return out
}
