package clrs

import (
	"reflect"
	"testing"
)

func TestDeque(t *testing.T) {
	d := deque{}
	d.New(5)
	snapshot1 := deque{head: 0, tail: 0, data: []int{0, 0, 0, 0, 0}}
	snapshot2 := deque{head: 0, tail: 1, data: []int{15, 0, 0, 0, 0}}
	snapshot3 := deque{head: 4, tail: 1, data: []int{15, 0, 0, 0, 6}}
	snapshot4 := deque{head: 3, tail: 1, data: []int{15, 0, 0, 9, 6}}
	snapshot5 := deque{head: 3, tail: 2, data: []int{15, 8, 0, 9, 6}}
	snapshot6 := deque{head: 4, tail: 2, data: []int{15, 8, 0, 9, 6}}
	snapshot7 := deque{head: 4, tail: 3, data: []int{15, 8, 2, 9, 6}}
	snapshot8 := deque{head: 0, tail: 3, data: []int{15, 8, 2, 9, 6}}
	snapshot9 := deque{head: 0, tail: 2, data: []int{15, 8, 2, 9, 6}}
	snapshot10 := deque{head: 1, tail: 2, data: []int{15, 8, 2, 9, 6}}
	snapshot11 := deque{head: 1, tail: 1, data: []int{15, 8, 2, 9, 6}}
	if !reflect.DeepEqual(d, snapshot1) {
		t.Errorf(" got %v", d)
		t.Errorf("want %v", snapshot1)
	}
	_, err := d.pop()
	if err != errUnderflow {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", errUnderflow)
	}
	_, err = d.shift()
	if err != errUnderflow {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", errUnderflow)
	}
	err = d.push(15)
	if err != nil || !reflect.DeepEqual(d, snapshot2) {
		t.Errorf(" got %v", d)
		t.Errorf(" err %v", err)
		t.Errorf("want %v", snapshot2)
	}
	err = d.unshift(6)
	if err != nil || !reflect.DeepEqual(d, snapshot3) {
		t.Errorf(" got %v", d)
		t.Errorf(" err %v", err)
		t.Errorf("want %v", snapshot3)
	}
	err = d.unshift(9)
	if err != nil || !reflect.DeepEqual(d, snapshot4) {
		t.Errorf(" got %v", d)
		t.Errorf(" err %v", err)
		t.Errorf("want %v", snapshot4)
	}
	err = d.push(8)
	if err != nil || !reflect.DeepEqual(d, snapshot5) {
		t.Errorf(" got %v", d)
		t.Errorf(" err %v", err)
		t.Errorf("want %v", snapshot5)
	}
	err = d.unshift(2)
	if err != errOverflow {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", errOverflow)
	}
	err = d.push(2)
	if err != errOverflow {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", errOverflow)
	}
	x, err := d.shift()
	if err != nil || x != 9 || !reflect.DeepEqual(d, snapshot6) {
		t.Errorf("  got %v", x)
		t.Errorf("  err %v", err)
		t.Errorf("deque %v", d)
		t.Errorf(" want %v", snapshot6)
	}
	err = d.push(2)
	if err != nil || !reflect.DeepEqual(d, snapshot7) {
		t.Errorf(" got %v", d)
		t.Errorf(" err %v", err)
		t.Errorf("want %v", snapshot7)
	}
	x, err = d.shift()
	if err != nil || x != 6 || !reflect.DeepEqual(d, snapshot8) {
		t.Errorf("  got %v", x)
		t.Errorf("  err %v", err)
		t.Errorf("deque %v", d)
		t.Errorf(" want %v", snapshot8)
	}
	x, err = d.pop()
	if err != nil || x != 2 || !reflect.DeepEqual(d, snapshot9) {
		t.Errorf("  got %v", x)
		t.Errorf("  err %v", err)
		t.Errorf("deque %v", d)
		t.Errorf(" want %v", snapshot9)
	}
	x, err = d.shift()
	if err != nil || x != 15 || !reflect.DeepEqual(d, snapshot10) {
		t.Errorf("  got %v", x)
		t.Errorf("  err %v", err)
		t.Errorf("deque %v", d)
		t.Errorf(" want %v", snapshot10)
	}
	x, err = d.pop()
	if err != nil || x != 8 || !reflect.DeepEqual(d, snapshot11) {
		t.Errorf("  got %v", x)
		t.Errorf("  err %v", err)
		t.Errorf("deque %v", d)
		t.Errorf(" want %v", snapshot11)
	}
}
