package clrs

func binaryAddition(a, b, c []byte) {
	flag := byte(0)
	for i := len(a) - 1; i >= 0; i-- {
		sum := a[i] + b[i] + flag
		c[i+1] = sum % 2
		if sum >= 2 {
			flag = 1
		} else {
			flag = 0
		}
	}
	c[0] = flag
}
