package clrs

func increment(a []byte) {
	n := len(a)
	i := 0
	for i < n && a[i] == 1 {
		a[i] = 0
		i = i + 1
	}
	if i < n {
		a[i] = 1
	}
}
