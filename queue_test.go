package clrs

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := queue{}
	q.New(12)
	snapshot1 := queue{head: 0, tail: 0, data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	snapshot2 := queue{head: 0, tail: 11, data: []int{2, 3, 4, 5, 6, 7, 15, 6, 9, 8, 4, 0}}
	snapshot3 := queue{head: 6, tail: 11, data: []int{2, 3, 4, 5, 6, 7, 15, 6, 9, 8, 4, 0}}
	snapshot4 := queue{head: 6, tail: 2, data: []int{3, 5, 4, 5, 6, 7, 15, 6, 9, 8, 4, 17}}
	snapshot5 := queue{head: 6, tail: 5, data: []int{3, 5, 7, 8, 9, 7, 15, 6, 9, 8, 4, 17}}
	snapshot6 := queue{head: 7, tail: 5, data: []int{3, 5, 7, 8, 9, 7, 15, 6, 9, 8, 4, 17}}
	snapshot7 := queue{head: 7, tail: 6, data: []int{3, 5, 7, 8, 9, 10, 15, 6, 9, 8, 4, 17}}
	if !reflect.DeepEqual(q, snapshot1) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot1)
	}
	_, err := q.dequeue()
	if err != ERR_UNDERFLOW {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", ERR_UNDERFLOW)
	}
	for _, x := range []int{2, 3, 4, 5, 6, 7, 15, 6, 9, 8, 4} {
		q.enqueue(x)
	}
	if !reflect.DeepEqual(q, snapshot2) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot2)
	}
	for _, want := range []int{2, 3, 4, 5, 6, 7} {
		got, err := q.dequeue()
		if got != want || err != nil {
			t.Errorf(" got %v", got)
			t.Errorf(" err %v", err)
			t.Errorf("want %v", want)
		}
	}
	if !reflect.DeepEqual(q, snapshot3) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot3)
	}
	for _, x := range []int{17, 3, 5} {
		q.enqueue(x)
	}
	if !reflect.DeepEqual(q, snapshot4) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot4)
	}
	for _, x := range []int{7, 8, 9} {
		q.enqueue(x)
	}
	if !reflect.DeepEqual(q, snapshot5) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot5)
	}
	err = q.enqueue(10)
	if err != ERR_OVERFLOW {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", ERR_OVERFLOW)
	}
	n, err := q.dequeue()
	if n != 15 || err != nil {
		t.Errorf(" got %v", n)
		t.Errorf(" err %v", err)
		t.Errorf("want %v", 15)
	}
	if !reflect.DeepEqual(q, snapshot6) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot6)
	}
	q.enqueue(10)
	if !reflect.DeepEqual(q, snapshot7) {
		t.Errorf(" got %v", q)
		t.Errorf("want %v", snapshot7)
	}
}
