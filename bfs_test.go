package clrs

import (
	"testing"
)

func TestBfs(t *testing.T) {
	g := new(graph)
	//            0    1    2    3    4    5    6    7
	rs := []rune{'r', 's', 't', 'u', 'v', 'w', 'x', 'y'}
	e := [][]int{
		{1, 4},
		{0, 5},
		{3, 5, 6},
		{2, 6, 7},
		{0},
		{1, 2, 6},
		{2, 3, 5, 7},
		{3, 6},
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
	g.bfs(g.v[1])
	for i, want := range []string{
		"sr",
		"s",
		"swt",
		"swxu",
		"srv",
		"sw",
		"swx",
		"swxy",
	} {
		got := g.printPath(g.v[1], g.v[i])
		if got != want {
			t.Errorf("path %c=>%c", g.v[1].key, g.v[i].key)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}
}
