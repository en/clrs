package clrs

type mat struct {
	rmin, rmax int
	cmin, cmax int
}

func squareMatrixAdd(a, b [][]int, ma, mb mat) [][]int {
	n := ma.rmax - ma.rmin + 1
	c := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		c[i] = make([]int, n)
	}
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-1; j++ {
			c[i][j] = a[ma.rmin+i][ma.cmin+j] + b[mb.rmin+i][mb.cmin+j]
		}
	}
	return c
}

func squareMatrixSubtract(a, b [][]int, ma, mb mat) [][]int {
	n := ma.rmax - ma.rmin + 1
	c := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		c[i] = make([]int, n)
	}
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-1; j++ {
			c[i][j] = a[ma.rmin+i][ma.cmin+j] - b[mb.rmin+i][mb.cmin+j]
		}
	}
	return c
}

func squareMatrixMultiply(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		c[i] = make([]int, n)
	}
	for i := 0; i <= n-1; i++ {
		for j := 0; j <= n-1; j++ {
			c[i][j] = 0
			for k := 0; k <= n-1; k++ {
				c[i][j] = c[i][j] + a[i][k]*b[k][j]
			}
		}
	}
	return c
}

func squareMatrixMultiplyRecursive(a, b [][]int, ma, mb mat) [][]int {
	n := ma.rmax - ma.rmin + 1
	c := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		c[i] = make([]int, n)
	}
	if n == 1 {
		c[0][0] = a[ma.rmin][ma.cmin] * b[mb.rmin][mb.cmin]
	} else {
		aRowMid := (ma.rmin + ma.rmax) / 2
		aColMid := (ma.cmin + ma.cmax) / 2
		a11 := mat{rmin: ma.rmin, rmax: aRowMid, cmin: ma.cmin, cmax: aColMid}
		a12 := mat{rmin: ma.rmin, rmax: aRowMid, cmin: aColMid + 1, cmax: ma.cmax}
		a21 := mat{rmin: aRowMid + 1, rmax: ma.rmax, cmin: ma.cmin, cmax: aColMid}
		a22 := mat{rmin: aRowMid + 1, rmax: ma.rmax, cmin: aColMid + 1, cmax: ma.cmax}

		bRowMid := (mb.rmin + mb.rmax) / 2
		bColMid := (mb.cmin + mb.cmax) / 2
		b11 := mat{rmin: mb.rmin, rmax: bRowMid, cmin: mb.cmin, cmax: bColMid}
		b12 := mat{rmin: mb.rmin, rmax: bRowMid, cmin: bColMid + 1, cmax: mb.cmax}
		b21 := mat{rmin: bRowMid + 1, rmax: mb.rmax, cmin: mb.cmin, cmax: bColMid}
		b22 := mat{rmin: bRowMid + 1, rmax: mb.rmax, cmin: bColMid + 1, cmax: mb.cmax}

		m := mat{0, n/2 - 1, 0, n/2 - 1}
		c11 := squareMatrixAdd(squareMatrixMultiplyRecursive(a, b, a11, b11), squareMatrixMultiplyRecursive(a, b, a12, b21), m, m)
		c12 := squareMatrixAdd(squareMatrixMultiplyRecursive(a, b, a11, b12), squareMatrixMultiplyRecursive(a, b, a12, b22), m, m)
		c21 := squareMatrixAdd(squareMatrixMultiplyRecursive(a, b, a21, b11), squareMatrixMultiplyRecursive(a, b, a22, b21), m, m)
		c22 := squareMatrixAdd(squareMatrixMultiplyRecursive(a, b, a21, b12), squareMatrixMultiplyRecursive(a, b, a22, b22), m, m)

		row := 0
		for i := range c11 {
			c[row] = append(c11[i], c12[i]...)
			row = row + 1
		}
		for i := range c21 {
			c[row] = append(c21[i], c22[i]...)
			row = row + 1
		}
	}
	return c
}

func squareMatrixMultiplyStrassen(a, b [][]int, ma, mb mat) [][]int {
	n := ma.rmax - ma.rmin + 1
	c := make([][]int, n)
	for i := 0; i <= n-1; i++ {
		c[i] = make([]int, n)
	}
	if n == 1 {
		c[0][0] = a[ma.rmin][ma.cmin] * b[mb.rmin][mb.cmin]
	} else {
		aRowMid := (ma.rmin + ma.rmax) / 2
		aColMid := (ma.cmin + ma.cmax) / 2
		a11 := mat{rmin: ma.rmin, rmax: aRowMid, cmin: ma.cmin, cmax: aColMid}
		a12 := mat{rmin: ma.rmin, rmax: aRowMid, cmin: aColMid + 1, cmax: ma.cmax}
		a21 := mat{rmin: aRowMid + 1, rmax: ma.rmax, cmin: ma.cmin, cmax: aColMid}
		a22 := mat{rmin: aRowMid + 1, rmax: ma.rmax, cmin: aColMid + 1, cmax: ma.cmax}

		bRowMid := (mb.rmin + mb.rmax) / 2
		bColMid := (mb.cmin + mb.cmax) / 2
		b11 := mat{rmin: mb.rmin, rmax: bRowMid, cmin: mb.cmin, cmax: bColMid}
		b12 := mat{rmin: mb.rmin, rmax: bRowMid, cmin: bColMid + 1, cmax: mb.cmax}
		b21 := mat{rmin: bRowMid + 1, rmax: mb.rmax, cmin: mb.cmin, cmax: bColMid}
		b22 := mat{rmin: bRowMid + 1, rmax: mb.rmax, cmin: bColMid + 1, cmax: mb.cmax}

		s1 := squareMatrixSubtract(b, b, b12, b22)
		s2 := squareMatrixAdd(a, a, a11, a12)
		s3 := squareMatrixAdd(a, a, a21, a22)
		s4 := squareMatrixSubtract(b, b, b21, b11)
		s5 := squareMatrixAdd(a, a, a11, a22)
		s6 := squareMatrixAdd(b, b, b11, b22)
		s7 := squareMatrixSubtract(a, a, a12, a22)
		s8 := squareMatrixAdd(b, b, b21, b22)
		s9 := squareMatrixSubtract(a, a, a11, a21)
		s10 := squareMatrixAdd(b, b, b11, b12)

		m := mat{0, n/2 - 1, 0, n/2 - 1}

		p1 := squareMatrixMultiplyStrassen(a, s1, a11, m)
		p2 := squareMatrixMultiplyStrassen(s2, b, m, b22)
		p3 := squareMatrixMultiplyStrassen(s3, b, m, b11)
		p4 := squareMatrixMultiplyStrassen(a, s4, a22, m)
		p5 := squareMatrixMultiplyStrassen(s5, s6, m, m)
		p6 := squareMatrixMultiplyStrassen(s7, s8, m, m)
		p7 := squareMatrixMultiplyStrassen(s9, s10, m, m)

		m1 := squareMatrixAdd(p5, p4, m, m)
		m2 := squareMatrixSubtract(m1, p2, m, m)
		c11 := squareMatrixAdd(m2, p6, m, m)
		c12 := squareMatrixAdd(p1, p2, m, m)
		c21 := squareMatrixAdd(p3, p4, m, m)
		m3 := squareMatrixAdd(p5, p1, m, m)
		m4 := squareMatrixSubtract(m3, p3, m, m)
		c22 := squareMatrixSubtract(m4, p7, m, m)

		row := 0
		for i := range c11 {
			c[row] = append(c11[i], c12[i]...)
			row = row + 1
		}
		for i := range c21 {
			c[row] = append(c21[i], c22[i]...)
			row = row + 1
		}
	}
	return c
}
