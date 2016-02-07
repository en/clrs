package clrs

import (
	"reflect"
	"testing"
)

func TestDaryHeap(t *testing.T) {
	in := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	want := []int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1}
	want2 := []int{15, 14, 10, 8, 7, 9, 3, 2, 1, 1}
	want3 := []int{20, 15, 10, 8, 14, 9, 3, 2, 1, 7}
	want4 := []int{1, 1, 1, 1, 1, 1, 3, 1, 1, 7}
	h := &daryHeap{}
	h.heapSize = 10
	h.children = 2
	h.a = make([]int, len(in))
	copy(h.a, in)
	h.heapIncreaseKey(8, 15)
	if !reflect.DeepEqual(h.a, want) {
		t.Errorf("orig %v", in)
		t.Errorf(" got %v", h.a)
		t.Errorf("want %v", want)
	}
	got, err := h.heapExtractMax()
	if got != 16 || h.heapSize != 9 || !reflect.DeepEqual(h.a, want2) || err != nil {
		t.Errorf("got %v", got)
		t.Errorf("err %v", err)
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
	h.maxHeapInsert(20)
	if h.heapSize != 10 || !reflect.DeepEqual(h.a, want3) {
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
	for i := 0; i < 9; i++ {
		h.heapExtractMax()
	}
	got, err = h.heapExtractMax()
	if got != 1 || h.heapSize != 0 || !reflect.DeepEqual(h.a, want4) || err != nil {
		t.Errorf("got %v", got)
		t.Errorf("err %v", err)
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
	in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want5 := []int{10, 6, 9, 3, 1, 4, 5, 2, 7, 8}
	h = &daryHeap{}
	h.heapSize = 0
	h.children = 3
	h.a = make([]int, len(in2))
	for _, v := range in2 {
		h.maxHeapInsert(v)
	}
	if h.heapSize != 10 || !reflect.DeepEqual(h.a, want5) {
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
		t.Errorf("want %v", want5)
	}
	want6 := 10
	want7 := []int{1, 1, 2, 3, 1, 4, 5, 2, 7, 8}
	for i := 0; i < 10; i++ {
		got, err = h.heapExtractMax()
		if got != want6 || err != nil {
			t.Errorf("got %v", got)
			t.Errorf("want %v", want6)
			t.Errorf("err %v", err)
		}
		want6 = want6 - 1
	}
	if h.heapSize != 0 || !reflect.DeepEqual(h.a, want7) {
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
		t.Errorf("want %v", want7)
	}
}
