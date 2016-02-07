package clrs

import (
	"reflect"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	in := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	want := []int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1}
	want2 := []int{15, 14, 10, 8, 7, 9, 3, 2, 1, 1}
	want3 := []int{20, 15, 10, 8, 14, 9, 3, 2, 1, 7}
	want4 := []int{1, 1, 1, 1, 1, 1, 3, 1, 1, 7}
	h := &heap{}
	h.heapSize = 10
	h.a = make([]int, len(in))
	copy(h.a, in)
	heapIncreaseKey(h, 8, 15)
	if !reflect.DeepEqual(h.a, want) {
		t.Errorf("orig %v", in)
		t.Errorf(" got %v", h.a)
		t.Errorf("want %v", want)
	}
	got, err := heapExtractMax(h)
	if got != 16 || h.heapSize != 9 || !reflect.DeepEqual(h.a, want2) || err != nil {
		t.Errorf("got %v", got)
		t.Errorf("err %v", err)
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
	got = heapMaximum(h)
	if got != 15 {
		t.Errorf("got %v", got)
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
	maxHeapInsert(h, 20)
	if h.heapSize != 10 || !reflect.DeepEqual(h.a, want3) {
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
	for i := 0; i < 9; i++ {
		heapExtractMax(h)
	}
	got, err = heapExtractMax(h)
	if got != 1 || h.heapSize != 0 || !reflect.DeepEqual(h.a, want4) || err != nil {
		t.Errorf("got %v", got)
		t.Errorf("err %v", err)
		t.Errorf("heap size %v", h.heapSize)
		t.Errorf("heap %v", h.a)
	}
}
