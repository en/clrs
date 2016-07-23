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

func TestProtoVEBMinimum(t *testing.T) {
	v := buildTestProtoVEBTree()
	got := v.minimum()
	want := 2
	if got != want {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestProtoVEBMaximum(t *testing.T) {
	v := buildTestProtoVEBTree()
	got := v.maximum()
	want := 15
	if got != want {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestProtoVEBSuccessor(t *testing.T) {
	v := buildTestProtoVEBTree()
	s := []int{2, 3, 4, 5, 7, 14, 15, -1}
	for i, x := range s[:7] {
		got := v.successor(x)
		want := s[i+1]
		if got != want {
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}
}

func TestProtoVEBPredecessor(t *testing.T) {
	v := buildTestProtoVEBTree()
	s := []int{-1, 2, 3, 4, 5, 7, 14, 15}
	for i, x := range s[1:] {
		got := v.predecessor(x)
		want := s[i]
		if got != want {
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}
}

func TestProtoVEBInsert(t *testing.T) {
	v := buildTestProtoVEBTree()
	v.insert(10)
	for i := 0; i < 16; i++ {
		got := v.member(i)
		want := 0
		switch i {
		case 2, 3, 4, 5, 7, 10, 14, 15:
			want = 1
		}
		if got != want {
			t.Errorf("member(%v)", i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}
}

func TestProtoVEBDelete(t *testing.T) {
	v := buildTestProtoVEBTree()
	v.delete(14)
	for i := 0; i < 16; i++ {
		got := v.member(i)
		want := 0
		switch i {
		case 2, 3, 4, 5, 7, 15:
			want = 1
		}
		if got != want {
			t.Errorf("member(%v)", i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}

	ssa1 := v.summary.summary.a[1]
	sc1a1 := v.summary.cluster[1].a[1]
	c3sa1 := v.cluster[3].summary.a[1]
	if ssa1 != 1 || sc1a1 != 1 || c3sa1 != 1 {
		t.Errorf("after delete 14")
		t.Errorf(" ssa1 %v", ssa1)
		t.Errorf("sc1a1 %v", sc1a1)
		t.Errorf("c3sa1 %v", c3sa1)
	}

	v.delete(15)
	for i := 0; i < 16; i++ {
		got := v.member(i)
		want := 0
		switch i {
		case 2, 3, 4, 5, 7:
			want = 1
		}
		if got != want {
			t.Errorf("member(%v)", i)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want)
		}
	}

	ssa1 = v.summary.summary.a[1]
	sc1a1 = v.summary.cluster[1].a[1]
	c3sa1 = v.cluster[3].summary.a[1]
	if ssa1 != 0 || sc1a1 != 0 || c3sa1 != 0 {
		t.Errorf("after delete 15")
		t.Errorf(" ssa1 %v", ssa1)
		t.Errorf("sc1a1 %v", sc1a1)
		t.Errorf("c3sa1 %v", c3sa1)
	}
}
