package clrs

func randomizedSelect(a []int, p, r, i int) int {
	if p == r {
		return a[p]
	}
	q := randomizedPartition(a, p, r)
	k := q - p + 1
	if i == k {
		return a[q]
	} else if i < k {
		return randomizedSelect(a, p, q-1, i)
	} else {
		return randomizedSelect(a, q+1, r, i-k)
	}
}
