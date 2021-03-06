package clrs

type chainedHash struct {
	table []doublyLinkedList
	h     func(int) int
}

func (ch *chainedHash) New(m int, h func(int) int) {
	ch.table = make([]doublyLinkedList, m)
	ch.h = h
}

func (ch *chainedHash) Insert(key int) {
	i := ch.h(key)
	node := new(dNode)
	node.key = key
	ch.table[i].listInsert(node)
}

func (ch *chainedHash) Search(key int) bool {
	i := ch.h(key)
	node := ch.table[i].listSearch(key)
	return node != nil
}

func (ch *chainedHash) Delete(key int) {
	i := ch.h(key)
	node := ch.table[i].listSearch(key)
	if node != nil {
		ch.table[i].listDelete(node)
	}
}

type openAddressingHash struct {
	m     int
	table []int
	h     func(int, int) int
}

func (oah *openAddressingHash) New(m int, h func(int, int) int) {
	oah.m = m
	oah.table = make([]int, m)
	for i := 0; i < m; i++ {
		oah.table[i] = -1
	}
	oah.h = h
}

func (oah *openAddressingHash) Insert(key int) int {
	for i := 0; i < oah.m; i++ {
		j := oah.h(key, i)
		if oah.table[j] == -1 || oah.table[j] == -2 {
			oah.table[j] = key
			return j
		}
	}
	return -1
}

func (oah *openAddressingHash) Search(key int) int {
	for i := 0; i < oah.m; i++ {
		j := oah.h(key, i)
		if oah.table[j] == -1 {
			return -1
		} else if oah.table[j] == key {
			return j
		}
	}
	return -1
}

func (oah *openAddressingHash) Delete(key int) {
	j := oah.Search(key)
	if j != -1 {
		oah.table[j] = -2 // deleted
	}
}
