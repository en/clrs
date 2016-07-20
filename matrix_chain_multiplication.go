package clrs

import (
	"fmt"
)

func matrixChainOrder(p []int) ([][]int, [][]int) {
	n := len(p) - 1
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	s := make([][]int, n)
	for i := range s {
		s[i] = make([]int, n)
	}
	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i-1][j-1] = -1
			for k := i; k <= j-1; k++ {
				q := m[i-1][k-1] + m[k][j-1] + p[i-1]*p[k]*p[j]
				if m[i-1][j-1] == -1 || q < m[i-1][j-1] {
					m[i-1][j-1] = q
					s[i-1][j-1] = k
				}
			}
		}
	}
	return m, s
}

func printOptimalParents(s [][]int, _i, _j int) string {
	var (
		out string
		f   func(int, int)
	)
	f = func(i, j int) {
		if i == j {
			out = out + "A" + fmt.Sprint(i)
		} else {
			out = out + "("
			f(i, s[i-1][j-1])
			f(s[i-1][j-1]+1, j)
			out = out + ")"
		}
	}
	f(_i, _j)
	return out
}

func memoizedMatrixChain(p []int) (int, [][]int) {
	n := len(p) - 1
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			m[i-1][j-1] = -1
		}
	}
	return lookupChain(m, p, 1, n), m
}

func lookupChain(m [][]int, p []int, i, j int) int {
	if m[i-1][j-1] != -1 {
		return m[i-1][j-1]
	}
	if i == j {
		m[i-1][j-1] = 0
	} else {
		for k := i; k <= j-1; k++ {
			q := lookupChain(m, p, i, k) + lookupChain(m, p, k+1, j) + p[i-1]*p[k]*p[j]
			if m[i-1][j-1] == -1 || q < m[i-1][j-1] {
				m[i-1][j-1] = q
			}
		}
	}
	return m[i-1][j-1]
}
