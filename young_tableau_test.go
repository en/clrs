package clrs

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestYoungTableau(t *testing.T) {
	yt := &youngTableau{}
	yt.n = 4
	yt.m = 4
	yt.a = [][]cell{
		[]cell{{2, false}, {3, false}, {12, false}, {14, false}},
		[]cell{{4, false}, {8, false}, {16, false}, {0, true}},
		[]cell{{5, false}, {9, false}, {0, true}, {0, true}},
		[]cell{{0, true}, {0, true}, {0, true}, {0, true}},
	}
	want := []cell{
		{2, false}, {3, false}, {4, false}, {5, false},
		{8, false}, {9, false}, {12, false}, {14, false},
		{16, false}, {0, true}, {0, true}, {0, true},
		{0, true}, {0, true}, {0, true}, {0, true},
	}
	index := 0
	for i := 0; i < 16; i++ {
		c, err := yt.extractMin()
		if err != nil || !reflect.DeepEqual(c, want[index]) {
			t.Errorf(" err %v", err)
			t.Errorf(" got %v", c)
			t.Errorf("want %v", want[index])
		}
		index = index + 1
	}
	// reset
	yt.a = [][]cell{
		[]cell{{2, false}, {3, false}, {12, false}, {14, false}},
		[]cell{{4, false}, {8, false}, {16, false}, {0, true}},
		[]cell{{5, false}, {9, false}, {0, true}, {0, true}},
		[]cell{{0, true}, {0, true}, {0, true}, {0, true}},
	}
	want1 := [][]cell{
		[]cell{{2, false}, {3, false}, {6, false}, {14, false}},
		[]cell{{4, false}, {8, false}, {12, false}, {16, false}},
		[]cell{{5, false}, {9, false}, {0, true}, {0, true}},
		[]cell{{0, true}, {0, true}, {0, true}, {0, true}},
	}
	yt.insert(cell{6, false})
	if !reflect.DeepEqual(yt.a, want1) {
		t.Errorf(" got %v", yt.a)
		t.Errorf("want %v", want1)
	}
	want2 := [][]cell{
		[]cell{{1, false}, {3, false}, {6, false}, {14, false}},
		[]cell{{2, false}, {4, false}, {8, false}, {16, false}},
		[]cell{{5, false}, {9, false}, {12, false}, {0, true}},
		[]cell{{0, true}, {0, true}, {0, true}, {0, true}},
	}
	yt.insert(cell{1, false})
	if !reflect.DeepEqual(yt.a, want2) {
		t.Errorf(" got %v", yt.a)
		t.Errorf("want %v", want1)
	}
}
func TestYoungTableauSort(t *testing.T) {
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	a := rand.Perm(9)
	yt := &youngTableau{}
	yt.n = 3
	yt.m = 3
	yt.a = make([][]cell, yt.m)
	for i := 0; i < yt.m; i++ {
		yt.a[i] = make([]cell, yt.n)
		for j := 0; j < yt.n; j++ {
			yt.a[i][j] = cell{0, true}
		}
	}
	for i := 0; i < 9; i++ {
		yt.insert(cell{a[i], false})
	}
	got := make([]int, 0)
	for i := 0; i < 9; i++ {
		c, err := yt.extractMin()
		if err == nil {
			got = append(got, c.value)
		} else {
			t.Errorf("extract min error %v", err)
		}
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
func TestYoungTableauFind(t *testing.T) {
	yt := &youngTableau{}
	yt.n = 4
	yt.m = 4
	yt.a = [][]cell{
		[]cell{{2, false}, {3, false}, {12, false}, {14, false}},
		[]cell{{4, false}, {8, false}, {16, false}, {0, true}},
		[]cell{{5, false}, {9, false}, {0, true}, {0, true}},
		[]cell{{0, true}, {0, true}, {0, true}, {0, true}},
	}

	for _, v := range []int{2, 3, 4, 5, 8, 9, 12, 14, 16} {
		if !yt.find(v) {
			t.Errorf("find %v", v)
		}
	}
	for _, v := range []int{0, 100} {
		if yt.find(v) {
			t.Errorf("find %v", v)
		}
	}
}
