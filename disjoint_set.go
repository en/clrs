package clrs

type component struct {
	v []rune
	e []struct {
		e1, e2 rune
	}
	s []map[rune]bool
}

func (g *component) makeSet(x rune) {
	m := make(map[rune]bool)
	m[x] = true
	g.s = append(g.s, m)
}

func (g *component) union(x, y rune) {
	var is []int
	xi := g.findSet(x)
	yi := g.findSet(y)
	if xi > yi {
		xi = xi - 1
		is = append(is, yi)
		is = append(is, xi)
	} else {
		yi = yi - 1
		is = append(is, xi)
		is = append(is, yi)
	}
	u := make(map[rune]bool)
	for _, i := range is {
		for k := range g.s[i] {
			u[k] = true
		}
		g.s = append(g.s[:i], g.s[i+1:]...)
	}
	g.s = append(g.s, u)
}

func (g component) findSet(x rune) int {
	for i, m := range g.s {
		if _, ok := m[x]; ok {
			return i
		}
	}
	return -1
}

func (g *component) connectedComponents() {
	for _, v := range g.v {
		g.makeSet(v)
	}
	for _, e := range g.e {
		f1 := g.findSet(e.e1)
		f2 := g.findSet(e.e2)
		if (f1 != -1) && (f2 != -1) && (f1 != f2) {
			g.union(e.e1, e.e2)
		}
	}
}

func (g component) sameComponent(u, v rune) bool {
	f1 := g.findSet(u)
	f2 := g.findSet(v)
	return (f1 != -1) && (f1 == f2)
}
