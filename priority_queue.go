package clrs

import (
	"errors"
	"math"
)

func heapMaximum(h *heap) int {
	return h.a[0]
}

func heapExtractMax(h *heap) (int, error) {
	if h.heapSize < 0 {
		return 0, errors.New("heap underflow")
	}
	max := h.a[0]
	h.a[0] = h.a[h.heapSize-1]
	h.heapSize = h.heapSize - 1
	maxHeapify(h, 0)
	return max, nil
}

func heapIncreaseKey(h *heap, i, key int) error {
	if key < h.a[i] {
		return errors.New("new key is smaller than current key")
	}
	h.a[i] = key
	for i > 0 && h.a[parent(i)] < h.a[i] {
		p := parent(i)
		h.a[i], h.a[p] = h.a[p], h.a[i]
		i = p
	}
	return nil
}

func maxHeapInsert(h *heap, key int) {
	h.heapSize = h.heapSize + 1
	h.a[h.heapSize-1] = math.MinInt32
	heapIncreaseKey(h, h.heapSize-1, key)
}
