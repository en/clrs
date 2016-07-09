package clrs

import (
	cheap "container/heap"
)

type item struct {
	r     rune
	freq  int
	index int
	left  *item
	right *item
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = j
	pq[j].index = i
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	i := x.(*item)
	i.index = n
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	i := old[n-1]
	i.index = -1
	*pq = old[0 : n-1]
	return i
}

func (pq *priorityQueue) Update(i *item, r rune, freq int) {
	i.r = r
	i.freq = freq
	cheap.Fix(pq, i.index)
}

func huffman(c priorityQueue) map[rune]string {
	cheap.Init(&c)
	n := len(c)
	for k := 1; k <= n-1; k++ {
		i := new(item)
		i.left = cheap.Pop(&c).(*item)
		i.right = cheap.Pop(&c).(*item)
		i.freq = i.left.freq + i.right.freq
		cheap.Push(&c, i)
	}
	out := make(map[rune]string)
	x := c[0]
	var f func(int, *item, string)
	f = func(dir int, i *item, code string) {
		if i != nil {
			if dir == 0 { // left
				code = code + "0"
			} else if dir == 1 { // right
				code = code + "1"
			}
			f(0, i.left, code)
			f(1, i.right, code)
			if i.r != 0 {
				out[i.r] = code
			}
		}
	}
	f(-1, x, "")

	return out
}
