package clrs

func countInversions(a []int, p, r int) int {
	inversions := 0
	if p < r {
		q := (p + r) / 2
		inversions = inversions + countInversions(a, p, q)
		inversions = inversions + countInversions(a, q+1, r)
		inversions = inversions + mergeInversions(a, p, q, r)
	}
	return inversions
}

func mergeInversions(a []int, p, q, r int) int {
	n1 := q - p + 1
	n2 := r - q
	left := make([]int, n1)
	right := make([]int, n2)
	for i := 0; i <= n1-1; i++ {
		left[i] = a[p+i]
	}
	for j := 0; j <= n2-1; j++ {
		right[j] = a[q+j+1]
	}
	i := 0
	j := 0
	inversions := 0
	counted := false
	k := p
	for ; k <= r; k++ {
		if i == n1 || j == n2 {
			break
		}
		if !counted && right[j] < left[i] {
			inversions = inversions + n1 - i
			counted = true
		}
		if left[i] <= right[j] {
			a[k] = left[i]
			i = i + 1
		} else {
			a[k] = right[j]
			j = j + 1
			counted = false
		}
	}
	for ; i <= n1-1; i++ {
		a[k] = left[i]
		k = k + 1
	}
	for ; j <= n2-1; j++ {
		a[k] = right[j]
		k = k + 1
	}
	return inversions
}
