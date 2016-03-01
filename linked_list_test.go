package clrs

import (
	"testing"
)

func TestSinglyLinkedList(t *testing.T) {
	l := singlyLinkedList{}
	var nodes []*sNode
	for _, v := range []int{15, 6, 9, 8} {
		x := new(sNode)
		x.key = v
		x.next = nil
		nodes = append(nodes, x)
	}
	for _, x := range nodes {
		l.listInsert(x)
	}
	x := l.head
	for i := len(nodes) - 1; i >= 0; i-- {
		if x != nodes[i] {
			t.Errorf(" got %v", x.key)
			t.Errorf("want %v", nodes[i].key)
		}
		x = x.next
	}
	for i, v := range []int{15, 6, 9, 8} {
		x := l.listSearch(v)
		if x != nodes[i] {
			t.Errorf(" got %v", x)
			t.Errorf("want %v", nodes[i])
		}
	}
	x = l.listSearch(7)
	if x != nil {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", nil)
	}
	err := l.listDelete(nodes[3])
	if err != nil {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", nil)
	}
	err = l.listDelete(nodes[0])
	if err != nil {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", nil)
	}
	err = l.listDelete(nodes[1])
	if err != nil {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", nil)
	}
	err = l.listDelete(l.head)
	if err != nil || l.head != nil {
		t.Errorf(" err %v", err)
		t.Errorf(" got %v", l.head)
		t.Errorf("want %v", nil)
	}
}

func TestDoublyLinkedList(t *testing.T) {
	l := doublyLinkedList{}
	var nodes []*dNode
	for _, v := range []int{15, 6, 9, 8} {
		x := new(dNode)
		x.key = v
		x.prev = nil
		x.next = nil
		nodes = append(nodes, x)
	}
	for _, x := range nodes {
		l.listInsert(x)
	}
	x := l.head
	for i := len(nodes) - 1; i >= 0; i-- {
		if x != nodes[i] {
			t.Errorf(" got %v", x.key)
			t.Errorf("want %v", nodes[i].key)
		}
		x = x.next
	}
	for i, v := range []int{15, 6, 9, 8} {
		x := l.listSearch(v)
		if x != nodes[i] {
			t.Errorf(" got %v", x)
			t.Errorf("want %v", nodes[i])
		}
	}
	x = l.listSearch(7)
	if x != nil {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", nil)
	}
	l.listDelete(nodes[3])
	l.listDelete(nodes[0])
	l.listDelete(nodes[1])
	l.listDelete(nodes[2])
	if l.head != nil {
		t.Errorf(" got %v", l.head)
		t.Errorf("want %v", nil)
	}
}

func TestSentinelLinkedList(t *testing.T) {
	l := dsLinkedList{}
	l.New()
	var nodes []*dNode
	for _, v := range []int{15, 6, 9, 8} {
		x := new(dNode)
		x.key = v
		x.prev = nil
		x.next = nil
		nodes = append(nodes, x)
	}
	for _, x := range nodes {
		l.listInsert(x)
	}
	x := l.sentinel.next
	for i := len(nodes) - 1; i >= 0; i-- {
		if x != nodes[i] {
			t.Errorf(" got %v", x.key)
			t.Errorf("want %v", nodes[i].key)
		}
		x = x.next
	}
	for i, v := range []int{15, 6, 9, 8} {
		x := l.listSearch(v)
		if x != nodes[i] {
			t.Errorf(" got %v", x)
			t.Errorf("want %v", nodes[i])
		}
	}
	x = l.listSearch(7)
	if x != l.sentinel {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", l.sentinel)
	}
	l.listDelete(nodes[3])
	l.listDelete(nodes[0])
	l.listDelete(nodes[1])
	l.listDelete(nodes[2])
	if l.sentinel.next != l.sentinel || l.sentinel.prev != l.sentinel {
		t.Errorf(" got sentinel.prev %v", l.sentinel.prev)
		t.Errorf("want sentinel.prev %v", l.sentinel)
		t.Errorf(" got sentinel.next %v", l.sentinel.next)
		t.Errorf("want sentinel.next %v", l.sentinel)
	}
}
