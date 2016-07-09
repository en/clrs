package clrs

import (
	"reflect"
	"testing"
)

func TestHuffman(t *testing.T) {
	i := priorityQueue{
		{'a', 45, 0, nil, nil},
		{'b', 13, 1, nil, nil},
		{'c', 12, 2, nil, nil},
		{'d', 16, 3, nil, nil},
		{'e', 9, 4, nil, nil},
		{'f', 5, 5, nil, nil},
	}
	want := map[rune]string{
		'a': "0",
		'b': "101",
		'c': "100",
		'd': "111",
		'e': "1101",
		'f': "1100",
	}
	got := huffman(i)
	if !reflect.DeepEqual(got, want) {
		t.Errorf(" got %v", got)
		t.Errorf("want %v", want)
	}
}
