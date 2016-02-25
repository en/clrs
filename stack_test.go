package clrs

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	snapshot1 := stack{top: 3, data: []int{15, 6, 2, 9}}
	snapshot2 := stack{top: 5, data: []int{15, 6, 2, 9, 17, 3}}
	snapshot3 := stack{top: 4, data: []int{15, 6, 2, 9, 17, 3}}
	snapshot4 := stack{top: 5, data: []int{15, 6, 2, 9, 17, 4}}
	s := stack{}
	s.New()
	if !s.empty() {
		t.Errorf("%v not empty", s)
	}
	n, err := s.pop()
	if err != ERR_UNDERFLOW {
		t.Errorf(" got %v", err)
		t.Errorf("want %v", ERR_UNDERFLOW)
	}
	s.push(15)
	s.push(6)
	s.push(2)
	s.push(9)
	if !reflect.DeepEqual(s, snapshot1) {
		t.Errorf(" got %v", s)
		t.Errorf("want %v", snapshot1)
	}
	s.push(17)
	s.push(3)
	if !reflect.DeepEqual(s, snapshot2) {
		t.Errorf(" got %v", s)
		t.Errorf("want %v", snapshot2)
	}
	n, err = s.pop()
	if err != nil {
		t.Errorf(" got %v", s)
		t.Errorf(" err %v", err)
	} else if n != 3 {
		t.Errorf(" got %v", n)
		t.Errorf("want %v", 3)
	}
	if !reflect.DeepEqual(s, snapshot3) {
		t.Errorf(" got %v", s)
		t.Errorf("want %v", snapshot3)
	}
	s.push(4)
	if !reflect.DeepEqual(s, snapshot4) {
		t.Errorf(" got %v", s)
		t.Errorf("want %v", snapshot4)
	}
}
