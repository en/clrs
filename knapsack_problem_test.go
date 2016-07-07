package clrs

import (
	"reflect"
	"testing"
)

func TestDynamic01Knapsack(t *testing.T) {
	v := []int{60, 100, 120}
	w := []int{10, 20, 30}
	want := []int{3, 2}
	c := dynamic01Knapsack(v, w, len(v), 50)
	got := printKnapsackSolution(c, w, 50)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
