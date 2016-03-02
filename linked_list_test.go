package clrs

import (
	"reflect"
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

func TestReverseSinglyLinkedList(t *testing.T) {
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
	l.reverse()
	x := l.head
	for _, n := range nodes {
		if x != n {
			t.Errorf(" got %v", x)
			t.Errorf("want %v", n)
		}
		x = x.next
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

func TestMultipleArrayLinkedList(t *testing.T) {
	snapshot1 := multipleArrayLinkedList{
		head: -1,
		free: 0,
		prev: []int{0, 0, 0, 0},
		key:  []int{0, 0, 0, 0},
		next: []int{1, 2, 3, -1},
	}
	snapshot2 := multipleArrayLinkedList{
		head: 3,
		free: -1,
		prev: []int{1, 2, 3, -1},
		key:  []int{15, 6, 9, 8},
		next: []int{-1, 0, 1, 2},
	}
	snapshot3 := multipleArrayLinkedList{
		head: 3,
		free: 2,
		prev: []int{1, 3, 3, -1},
		key:  []int{15, 6, 9, 8},
		next: []int{-1, 0, -1, 1},
	}
	snapshot4 := multipleArrayLinkedList{
		head: 3,
		free: 0,
		prev: []int{1, 3, 3, -1},
		key:  []int{15, 6, 9, 8},
		next: []int{2, -1, -1, 1},
	}
	snapshot5 := multipleArrayLinkedList{
		head: 1,
		free: 3,
		prev: []int{1, -1, 3, -1},
		key:  []int{15, 6, 9, 8},
		next: []int{2, -1, -1, 0},
	}
	snapshot6 := multipleArrayLinkedList{
		head: 3,
		free: 0,
		prev: []int{1, 3, 3, -1},
		key:  []int{15, 6, 9, 5},
		next: []int{2, -1, -1, 1},
	}
	l := multipleArrayLinkedList{}
	l.New(4)
	if !reflect.DeepEqual(l, snapshot1) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot1)
	}
	for _, v := range []int{15, 6, 9, 8} {
		err := l.listInsert(v)
		if err != nil {
			t.Errorf(" got %v", err)
			t.Errorf("want %v", nil)
		}
	}
	if !reflect.DeepEqual(l, snapshot2) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot2)
	}
	err := l.listInsert(5)
	if err != errOutOfSpace {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", errOutOfSpace)
	}
	l.listDelete(2)
	if !reflect.DeepEqual(l, snapshot3) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot3)
	}
	l.listDelete(0)
	if !reflect.DeepEqual(l, snapshot4) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot4)
	}
	l.listDelete(3)
	if !reflect.DeepEqual(l, snapshot5) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot5)
	}
	err = l.listInsert(5)
	if err != nil || !reflect.DeepEqual(l, snapshot6) {
		t.Errorf(" err %v", err)
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot6)
	}
	if x := l.listSearch(5); x != 3 {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", 3)
	}
	if x := l.listSearch(6); x != 1 {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", 1)
	}
	if x := l.listSearch(15); x != -1 {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", -1)
	}
}

func TestSingleArrayLinkedList(t *testing.T) {
	snapshot1 := singleArrayLinkedList{
		head: -1,
		free: 0,
		data: []int{
			0, 3, 0,
			0, 6, 0,
			0, 9, 0,
			0, -1, 0,
		},
	}
	snapshot2 := singleArrayLinkedList{
		head: 9,
		free: -1,
		data: []int{
			15, -1, 3,
			6, 0, 6,
			9, 3, 9,
			8, 6, -1,
		},
	}
	snapshot3 := singleArrayLinkedList{
		head: 9,
		free: 6,
		data: []int{
			15, -1, 3,
			6, 0, 9,
			9, -1, 9,
			8, 3, -1,
		},
	}
	snapshot4 := singleArrayLinkedList{
		head: 9,
		free: 0,
		data: []int{
			15, 6, 3,
			6, -1, 9,
			9, -1, 9,
			8, 3, -1,
		},
	}
	snapshot5 := singleArrayLinkedList{
		head: 3,
		free: 9,
		data: []int{
			15, 6, 3,
			6, -1, -1,
			9, -1, 9,
			8, 0, -1,
		},
	}
	snapshot6 := singleArrayLinkedList{
		head: 9,
		free: 0,
		data: []int{
			15, 6, 3,
			6, -1, 9,
			9, -1, 9,
			5, 3, -1,
		},
	}
	l := singleArrayLinkedList{}
	l.New(4)
	if !reflect.DeepEqual(l, snapshot1) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot1)
	}
	for _, v := range []int{15, 6, 9, 8} {
		err := l.listInsert(v)
		if err != nil {
			t.Errorf(" got %v", err)
			t.Errorf("want %v", nil)
		}
	}
	if !reflect.DeepEqual(l, snapshot2) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot2)
	}
	err := l.listInsert(5)
	if err != errOutOfSpace {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", errOutOfSpace)
	}
	l.listDelete(6)
	if !reflect.DeepEqual(l, snapshot3) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot3)
	}
	l.listDelete(0)
	if !reflect.DeepEqual(l, snapshot4) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot4)
	}
	l.listDelete(9)
	if !reflect.DeepEqual(l, snapshot5) {
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot5)
	}
	err = l.listInsert(5)
	if err != nil || !reflect.DeepEqual(l, snapshot6) {
		t.Errorf(" err %v", err)
		t.Errorf(" got %v", l)
		t.Errorf("want %v", snapshot6)
	}
	if x := l.listSearch(5); x != 9 {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", 9)
	}
	if x := l.listSearch(6); x != 3 {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", 3)
	}
	if x := l.listSearch(15); x != -1 {
		t.Errorf(" got %v", x)
		t.Errorf("want %v", -1)
	}
}
