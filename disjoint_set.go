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
	xi := g.findSet(x)
	yi := g.findSet(y)
	for k := range g.s[yi] {
		g.s[xi][k] = true
	}
	g.s = append(g.s[:yi], g.s[yi+1:]...)
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
