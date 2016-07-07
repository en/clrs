package clrs

func recursiveActivitySelector(s, f []int, k, n int) []int {
	var a []int
	m := k + 1
	for m <= n && s[m] < f[k] {
		m = m + 1
	}
	if m <= n {
		return append(recursiveActivitySelector(s, f, m, n), m)
	}
	return a
}

func greedyActivitySelector(s, f []int) []int {
	n := len(s) - 1
	a := []int{1}
	k := 1
	for m := 2; m <= n; m++ {
		if s[m] >= f[k] {
			a = append(a, m)
			k = m
		}
	}
	return a
}
