package clrs

import (
	"container/list"
	"fmt"
)

func (g graph) bfs(s *vertex) {
	for v := range g.adj {
		if v != s {
			v.color = WHITE
			v.d = -1
			v.p = nil
		}
	}
	s.color = GRAY
	s.d = 0
	s.p = nil
	q := list.New()
	q.PushBack(s)
	for q.Len() > 0 {
		e := q.Front()
		u := e.Value.(*vertex)
		q.Remove(e)
		for n := g.adj[u]; n != nil; n = n.next {
			v := n.v
			if v.color == WHITE {
				v.color = GRAY
				v.d = u.d + 1
				v.p = u
				q.PushBack(v)
			}
		}
		u.color = SCHWARZ
	}
}

func (g graph) printPath(_s, _v *vertex) string {
	var (
		out string
		f   func(*vertex, *vertex)
	)
	f = func(s, v *vertex) {
		if v == s {
			out = out + fmt.Sprintf("%c", s.key)
		} else if v.p == nil {
			out += fmt.Sprintf("no path from %c to %c exists", s.key, v.key)
		} else {
			f(s, v.p)
			out = out + fmt.Sprintf("%c", v.key)
		}
	}
	f(_s, _v)
	return out
}
