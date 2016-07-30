package clrs

func (g *graph) dfs() {
	for _, u := range g.v {
		u.color = WHITE
		u.p = nil
	}
	g.time = 0
	for _, u := range g.v {
		if u.color == WHITE {
			g.dfsVisit(u)
		}
	}
}

func (g *graph) dfsVisit(u *vertex) {
	g.time = g.time + 1
	u.d = g.time
	u.color = GRAY
	for n := g.adj[u]; n != nil; n = n.next {
		v := n.v
		if v.color == WHITE {
			v.p = u
			g.dfsVisit(v)
		}
	}
	u.color = SCHWARZ
	g.time = g.time + 1
	u.f = g.time
}
