package clrs

import (
	"reflect"
	"testing"
)

func hMod9(k int) int {
	return k % 9
}

func h(k int) int {
	return k
}

func hh(k int) int {
	m := 11
	return 1 + (k % (m - 1))
}

func hLinearProbing11(k, i int) int {
	m := 11
	return (h(k) + i) % m
}

func hQuadraticProbing11(k, i int) int {
	m := 11
	c1 := 1
	c2 := 3
	return (h(k) + c1*i + c2*i*i) % m
}

func hDoubleHashing11(k, i int) int {
	m := 11
	return (h(k) + i*hh(k)) % m
}

func TestChainedHash(t *testing.T) {
	ht := new(chainedHash)
	ht.New(9, hMod9)
	for _, v := range []int{5, 28, 19, 15, 20, 33, 12, 17, 10} {
		ht.Insert(v)
	}
	want := [][]int{
		[]int{},
		[]int{10, 19, 28},
		[]int{20},
		[]int{12},
		[]int{},
		[]int{5},
		[]int{33, 15},
		[]int{},
		[]int{17},
	}
	got := [][]int{}
	for i := 0; i < 9; i++ {
		got = append(got, ht.table[i].toArray())
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}

	if !ht.Search(19) {
		t.Errorf("search %v", 19)
		t.Errorf("   got %v", false)
		t.Errorf("  want %v", true)
	}
	ht.Delete(19)
	if ht.Search(19) {
		t.Errorf("search %v", 19)
		t.Errorf("   got %v", true)
		t.Errorf("  want %v", false)
	}
}

func TestLinearProbing(t *testing.T) {
	ht := new(openAddressingHash)
	ht.New(11, hLinearProbing11)
	for _, v := range []int{10, 22, 31, 4, 15, 28, 17, 88, 59} {
		ht.Insert(v)
	}
	want := []int{22, 88, -1, -1, 4, 15, 28, 17, 59, 31, 10}
	want2 := []int{22, 88, -1, -1, -2, 15, 28, 17, 59, 31, 10}
	if !reflect.DeepEqual(ht.table, want) {
		t.Errorf(" got %v", ht.table)
		t.Errorf("want %v", want)
	}
	if i := ht.Search(4); i != 4 {
		t.Errorf("search %v", 4)
		t.Errorf("   got %v", i)
		t.Errorf("  want %v", 4)
	}
	ht.Delete(4)
	if i := ht.Search(4); i != -1 {
		t.Errorf("search %v", 4)
		t.Errorf("   got %v", i)
		t.Errorf("  want %v", -1)
	}
	if !reflect.DeepEqual(ht.table, want2) {
		t.Errorf(" got %v", ht.table)
		t.Errorf("want %v", want2)
	}
}

func TestQuadraticProbing(t *testing.T) {
	ht := new(openAddressingHash)
	ht.New(11, hQuadraticProbing11)
	for _, v := range []int{10, 22, 31, 4, 15, 28, 17, 88, 59} {
		ht.Insert(v)
	}
	want := []int{22, -1, 88, 17, 4, -1, 28, 59, 15, 31, 10}
	want2 := []int{22, -1, 88, 17, -2, -1, 28, 59, 15, 31, 10}
	if !reflect.DeepEqual(ht.table, want) {
		t.Errorf(" got %v", ht.table)
		t.Errorf("want %v", want)
	}
	if i := ht.Search(4); i != 4 {
		t.Errorf("search %v", 4)
		t.Errorf("   got %v", i)
		t.Errorf("  want %v", 4)
	}
	ht.Delete(4)
	if i := ht.Search(4); i != -1 {
		t.Errorf("search %v", 4)
		t.Errorf("   got %v", i)
		t.Errorf("  want %v", -1)
	}
	if !reflect.DeepEqual(ht.table, want2) {
		t.Errorf(" got %v", ht.table)
		t.Errorf("want %v", want2)
	}
}

func TestDoubleHasing(t *testing.T) {
	ht := new(openAddressingHash)
	ht.New(11, hDoubleHashing11)
	for _, v := range []int{10, 22, 31, 4, 15, 28, 17, 88, 59} {
		ht.Insert(v)
	}
	want := []int{22, -1, 59, 17, 4, 15, 28, 88, -1, 31, 10}
	want2 := []int{22, -1, 59, 17, -2, 15, 28, 88, -1, 31, 10}
	if !reflect.DeepEqual(ht.table, want) {
		t.Errorf(" got %v", ht.table)
		t.Errorf("want %v", want)
	}
	if i := ht.Search(4); i != 4 {
		t.Errorf("search %v", 4)
		t.Errorf("   got %v", i)
		t.Errorf("  want %v", 4)
	}
	ht.Delete(4)
	if i := ht.Search(4); i != -1 {
		t.Errorf("search %v", 4)
		t.Errorf("   got %v", i)
		t.Errorf("  want %v", -1)
	}
	if !reflect.DeepEqual(ht.table, want2) {
		t.Errorf(" got %v", ht.table)
		t.Errorf("want %v", want2)
	}
}
