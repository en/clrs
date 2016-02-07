package clrs

import (
	"errors"
	"math"
)

type daryHeap struct {
	a        []int
	heapSize int
	children int
}

func (h *daryHeap) parent(i int) int {
	return (i - 1) / h.children
}

func (h *daryHeap) child(i, index int) int {
	return h.children*i + index
}

func (h *daryHeap) maxHeapify(i int) {
	largest := i
	for index := 1; index <= h.children; index++ {
		c := h.child(i, index)
		if c < h.heapSize {
			if h.a[c] > h.a[largest] {
				largest = c
			}
		} else {
			break
		}
	}
	if largest != i {
		h.a[i], h.a[largest] = h.a[largest], h.a[i]
		h.maxHeapify(largest)
	}
}

func (h *daryHeap) heapExtractMax() (int, error) {
	if h.heapSize < 0 {
		return 0, errors.New("heap underflow")
	}
	max := h.a[0]
	h.a[0] = h.a[h.heapSize-1]
	h.heapSize = h.heapSize - 1
	h.maxHeapify(0)
	return max, nil
}

func (h *daryHeap) heapIncreaseKey(i, key int) error {
	if key < h.a[i] {
		return errors.New("new key is smaller than current key")
	}
	h.a[i] = key
	for i > 0 && h.a[h.parent(i)] < h.a[i] {
		p := h.parent(i)
		h.a[i], h.a[p] = h.a[p], h.a[i]
		i = p
	}
	return nil
}

func (h *daryHeap) maxHeapInsert(key int) {
	h.heapSize = h.heapSize + 1
	h.a[h.heapSize-1] = math.MinInt32
	h.heapIncreaseKey(h.heapSize-1, key)
}
