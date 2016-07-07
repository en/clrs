package clrs

import (
	"reflect"
	"testing"
)

var (
	activitySelectorS = []int{0, 1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	activitySelectorF = []int{0, 4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
)

func TestRecursiveActivitySelector(t *testing.T) {
	want := []int{11, 8, 4, 1}
	got := recursiveActivitySelector(activitySelectorS, activitySelectorF, 0, len(activitySelectorS)-1)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}

func TestGreedyActivitySelector(t *testing.T) {
	want := []int{1, 4, 8, 11}
	got := greedyActivitySelector(activitySelectorS, activitySelectorF)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
