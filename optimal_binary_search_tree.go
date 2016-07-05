package clrs

import (
	"fmt"
)

func optimalBst(p, q []float32, n int) ([][]float32, [][]int) {
	e := make([][]float32, n+1)
	for i := range e {
		e[i] = make([]float32, n+1)
	}
	w := make([][]float32, n+1)
	for i := range w {
		w[i] = make([]float32, n+1)
	}
	root := make([][]int, n)
	for i := range root {
		root[i] = make([]int, n)
	}
	for i := 1; i <= n+1; i++ {
		e[i-1][i-1] = q[i-1]
		w[i-1][i-1] = q[i-1]
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			e[i-1][j] = -1
			w[i-1][j] = w[i-1][j-1] + p[j-1] + q[j]
			for r := i; r <= j; r++ {
				t := e[i-1][r-1] + e[r][j] + w[i-1][j]
				if e[i-1][j] == -1 || t < e[i-1][j] {
					e[i-1][j] = t
					root[i-1][j-1] = r
				}
			}
		}
	}
	return e, root
}

// see exercise 15.5-1
func constructOptimalBst(root [][]int) []string {
	var (
		out []string
		f   func(int, int)
	)
	n := len(root)
	out = append(out, fmt.Sprintf("k%v is the root", root[0][n-1]))
	f = func(i, j int) {
		r := root[i-1][j-1]
		if i == r {
			out = append(out, fmt.Sprintf("d%v is the left child of k%v", r-1, r))
		} else {
			out = append(out, fmt.Sprintf("k%v is the left child of k%v", root[i-1][r-2], r))
			f(i, r-1)
		}
		if j == r {
			out = append(out, fmt.Sprintf("d%v is the right child of k%v", r, r))
		} else {
			out = append(out, fmt.Sprintf("k%v is the right child of k%v", root[r][j-1], r))
			f(r+1, j)
		}
	}
	f(1, n)
	return out
}
