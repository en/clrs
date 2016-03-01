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
