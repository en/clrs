package clrs

import (
	// "reflect"
	"testing"
)

func buildTestProtoVEBTree() *protoVEB {
	var (
		p2s []*protoVEB
		p4s []*protoVEB
	)
	for _, as := range [][]int{
		{1, 1},
		{1, 1},
		{0, 1},
		{0, 1},
		{0, 0},
		{1, 1},
		{1, 1},
		{1, 1},
		{0, 1},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 1},
		{0, 0},
		{1, 1},
	} {
		p := new(protoVEB)
		p.u = 2
		p.a = as
		p2s = append(p2s, p)
	}
	for i := 0; i < 5; i++ {
		p := new(protoVEB)
		p.u = 4
		p.summary = p2s[i*3]
		p.cluster = make([]*protoVEB, 2)
		p.cluster[0] = p2s[i*3+1]
		p.cluster[1] = p2s[i*3+2]
		p4s = append(p4s, p)
	}
	p16 := new(protoVEB)
	p16.u = 16
	p16.summary = p4s[0]
	p16.cluster = make([]*protoVEB, 4)
	for i := 0; i < 4; i++ {
		p16.cluster[i] = p4s[i+1]
	}
	return p16
}

func TestProtoVEBMember(t *testing.T) {
	v := buildTestProtoVEBTree()
	for i := 0; i < 16; i++ {
		got := v.member(i)
		want := 0
		switch i {
		case 2, 3, 4, 5, 7, 14, 15:
			want = 1
		}
		if got != want {
			t.Errorf("member(%v)", i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}
}
