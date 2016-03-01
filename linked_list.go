package clrs

import (
	"errors"
)

type sNode struct {
	key  int
	next *sNode
}

type singlyLinkedList struct {
	head *sNode
}

func (l *singlyLinkedList) listSearch(k int) *sNode {
	x := l.head
	for x != nil && x.key != k {
		x = x.next
	}
	return x
}

func (l *singlyLinkedList) listInsert(x *sNode) {
	x.next = l.head
	l.head = x
}

func (l *singlyLinkedList) listSearchPrev(n *sNode) (*sNode, error) {
	x := l.head
	var prev *sNode
	for x != nil && x != n {
		prev = x
		x = x.next
	}
	if x != n {
		return nil, errors.New("node not exists")
	}
	return prev, nil
}

func (l *singlyLinkedList) listDelete(x *sNode) error {
	prev, err := l.listSearchPrev(x)
	if err != nil {
		return err
	}
	if prev == nil {
		l.head = x.next
	} else {
		prev.next = x.next
	}
	return nil
}

func (l *singlyLinkedList) reverse() {
	var prev *sNode
	x := l.head
	for x != nil {
		next := x.next
		x.next = prev
		prev = x
		x = next
	}
	l.head = prev
}

type dNode struct {
	prev *dNode
	key  int
	next *dNode
}

type doublyLinkedList struct {
	head *dNode
}

func (l *doublyLinkedList) listSearch(k int) *dNode {
	x := l.head
	for x != nil && x.key != k {
		x = x.next
	}
	return x
}

func (l *doublyLinkedList) listInsert(x *dNode) {
	x.next = l.head
	if l.head != nil {
		l.head.prev = x
	}
	l.head = x
	x.prev = nil
}

func (l *doublyLinkedList) listDelete(x *dNode) {
	if x.prev != nil {
		x.prev.next = x.next
	} else {
		l.head = x.next
	}
	if x.next != nil {
		x.next.prev = x.prev
	}
}

// circular, doubly linked list with a sentinel
type dsLinkedList struct {
	sentinel *dNode
}

func (l *dsLinkedList) New() {
	x := new(dNode)
	x.prev = x
	x.next = x
	l.sentinel = x
}

func (l *dsLinkedList) listSearch(k int) *dNode {
	x := l.sentinel.next
	for x != l.sentinel && x.key != k {
		x = x.next
	}
	return x
}

func (l *dsLinkedList) listInsert(x *dNode) {
	x.next = l.sentinel.next
	l.sentinel.next.prev = x
	l.sentinel.next = x
	x.prev = l.sentinel
}

func (l *dsLinkedList) listDelete(x *dNode) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

type multipleArrayLinkedList struct {
	head int
	free int
	prev []int
	key  []int
	next []int
}

func (l *multipleArrayLinkedList) New(size int) {
	l.head = -1
	l.free = -1
	if size > 0 {
		l.prev = make([]int, size)
		l.key = make([]int, size)
		l.next = make([]int, size)
		for i := 0; i <= size-2; i++ {
			l.next[i] = i + 1
		}
		l.next[size-1] = -1
		l.free = 0
	}
}

func (l *multipleArrayLinkedList) allocateObject() (int, error) {
	if l.free == -1 {
		return -1, errOutOfSpace
	}
	x := l.free
	l.free = l.next[x]
	return x, nil
}

func (l *multipleArrayLinkedList) freeObject(x int) {
	l.next[x] = l.free
	l.free = x
}

func (l *multipleArrayLinkedList) listSearch(k int) int {
	x := l.head
	for x != -1 && l.key[x] != k {
		x = l.next[x]
	}
	return x
}

func (l *multipleArrayLinkedList) listInsert(k int) error {
	i, err := l.allocateObject()
	if err != nil {
		return err
	}
	l.key[i] = k
	l.next[i] = l.head
	if l.head != -1 {
		l.prev[l.head] = i
	}
	l.head = i
	l.prev[i] = -1
	return nil
}

func (l *multipleArrayLinkedList) listDelete(x int) {
	if l.prev[x] != -1 {
		l.next[l.prev[x]] = l.next[x]
	} else {
		l.head = l.next[x]
	}
	if l.next[x] != -1 {
		l.prev[l.next[x]] = l.prev[x]
	}
	l.freeObject(x)
}
