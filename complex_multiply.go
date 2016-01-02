package clrs

func complexMultiply(a, b, c, d int) (int, int) {
	x := b * (c + d)
	y := c * (a + b)
	z := a * (c - d)

	// real
	// y - x = ac + bc - bc - bd = ac - bd
	r := y - x
	// imaginary
	// y - z = ac + bc - ac + ad = ad + bc
	i := y - z
	return r, i
}
