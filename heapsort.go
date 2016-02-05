package clrs

type heap struct {
	a        []int
	heapSize int
}

func (h *heap) length() int {
	return len(h.a)
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func maxHeapify(h *heap, i int) {
	l := left(i)
	r := right(i)
	var largest int
	if l < h.heapSize && h.a[l] > h.a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < h.heapSize && h.a[r] > h.a[largest] {
		largest = r
	}
	if largest != i {
		h.a[i], h.a[largest] = h.a[largest], h.a[i]
		maxHeapify(h, largest)
	}
}

func buildMaxHeap(h *heap) {
	h.heapSize = h.length()
	for i := (h.length() / 2) - 1; i >= 0; i-- {
		maxHeapify(h, i)
	}
}

func heapsort(h *heap) {
	buildMaxHeap(h)
	for i := h.length() - 1; i >= 1; i-- {
		h.a[0], h.a[i] = h.a[i], h.a[0]
		h.heapSize = h.heapSize - 1
		maxHeapify(h, 0)
	}
}
