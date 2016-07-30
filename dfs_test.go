package clrs

import (
	"testing"
)

func TestDfs(t *testing.T) {
	g := new(graph)
	//            0    1    2    3    4    5
	rs := []rune{'u', 'v', 'w', 'x', 'y', 'z'}
	e := [][]int{
		{3, 1},
		{4},
		{5, 4},
		{1},
		{3},
		{5},
	}
	g.v = make([]*vertex, 0)
	for _, r := range rs {
		k := new(vertex)
		k.key = r
		g.v = append(g.v, k)
	}
	g.adj = make(map[*vertex]*adjNode)
	for i, k := range g.v {
		for _, a := range e[i] {
			an := new(adjNode)
			an.v = g.v[a]
			if _, ok := g.adj[k]; ok {
				an.next = g.adj[k]
			}
			g.adj[k] = an
		}
	}
	g.dfs()
	for i, want := range [][]int{
		{1, 8},
		{2, 7},
		{9, 12},
		{4, 5},
		{3, 6},
		{10, 11},
	} {
		if g.v[i].d != want[0] || g.v[i].f != want[1] {
			t.Errorf("vertex:%c", g.v[i].key)
			t.Errorf(" got d:%v, f:%v", g.v[i].d, g.v[i].f)
			t.Errorf("want d:%v, f:%v", want[0], want[1])
		}
	}
}
