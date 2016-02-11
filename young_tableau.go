package clrs

type cell struct {
	value    int
	infinity bool
}

func (c cell) lt(a cell) bool {
	if c.infinity && a.infinity {
		return false
	}
	if !c.infinity && !a.infinity {
		return c.value < a.value
	}
	if a.infinity {
		return true
	}
	return false
}

func (c cell) gt(a cell) bool {
	if c.infinity && a.infinity {
		return false
	}
	if !c.infinity && !a.infinity {
		return c.value > a.value
	}
	if c.infinity {
		return true
	}
	return false
}

func (c cell) eq(a cell) bool {
	return (c.infinity == a.infinity) && (c.value == a.value)
}

type youngTableau struct {
	a [][]cell
	n int
	m int
}

func (t *youngTableau) right(i, j int) (int, int) {
	return i, j + 1
}

func (t *youngTableau) below(i, j int) (int, int) {
	return i + 1, j
}

func (t *youngTableau) above(i, j int) (int, int) {
	return i - 1, j
}

func (t *youngTableau) left(i, j int) (int, int) {
	return i, j - 1
}

func (t *youngTableau) minHeapify(i, j int) {
	rx, ry := t.right(i, j)
	bx, by := t.below(i, j)
	sx, sy := i, j
	if rx < t.n && ry < t.m && t.a[rx][ry].lt(t.a[sx][sy]) {
		sx, sy = rx, ry
	}
	if bx < t.n && by < t.m && t.a[bx][by].lt(t.a[sx][sy]) {
		sx, sy = bx, by
	}
	if sx != i || sy != j {
		t.a[sx][sy], t.a[i][j] = t.a[i][j], t.a[sx][sy]
		t.minHeapify(sx, sy)
	}
}

func (t *youngTableau) extractMin() (cell, error) {
	min := t.a[0][0]
	t.a[0][0] = cell{0, true}
	t.minHeapify(0, 0)
	return min, nil
}

func (t *youngTableau) insert(c cell) {
	i, j := t.m-1, t.n-1
	t.a[i][j] = c
	for {
		x, y := i, j
		if i > 0 {
			ax, ay := t.above(i, j)
			if t.a[ax][ay].gt(t.a[x][y]) {
				x, y = ax, ay
			}
		}
		if j > 0 {
			lx, ly := t.left(i, j)
			if t.a[lx][ly].gt(t.a[x][y]) {
				x, y = lx, ly
			}
		}
		if x == i && y == j {
			break
		} else {
			t.a[i][j], t.a[x][y] = t.a[x][y], t.a[i][j]
			i, j = x, y
		}
	}
}

func (t *youngTableau) find(v int) bool {
	i, j := t.m-1, 0
	c := cell{v, false}
	for i >= 0 && j < t.n {
		if t.a[i][j].eq(c) {
			return true
		} else if t.a[i][j].gt(c) {
			i, j = t.above(i, j)
		} else {
			i, j = t.right(i, j)
		}
	}
	return false
}
