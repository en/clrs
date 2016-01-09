package clrs

import (
	"reflect"
	"strings"
	"testing"
)

var mongeArraysCases = []struct {
	a    [][]int
	want []int
}{
	{
		[][]int{
			[]int{37, 23},
			[]int{21, 6},
		},
		[]int{
			1,
			1,
		},
	},
	{
		[][]int{
			[]int{37, 23, 24, 32},
			[]int{21, 6, 7, 10},
			[]int{53, 34, 30, 31},
			[]int{32, 13, 9, 6},
			[]int{43, 21, 15, 8},
		},
		[]int{
			1,
			1,
			2,
			3,
			3,
		},
	},
	{
		[][]int{
			[]int{10, 17, 13, 28, 23},
			[]int{17, 22, 16, 29, 23},
			[]int{24, 28, 22, 34, 24},
			[]int{11, 13, 6, 17, 7},
			[]int{45, 44, 32, 37, 23},
			[]int{36, 33, 19, 21, 6},
			[]int{75, 66, 51, 53, 34},
		},
		[]int{
			0,
			2,
			2,
			2,
			4,
			4,
			4,
		},
	},
}

func printMongeArrays(t *testing.T, label string, m [][]int) {
	t.Errorf("%v", label)
	spaces := strings.Repeat(" ", len(label))
	for _, r := range m {
		t.Errorf("%v %v", spaces, r)
	}
}

func TestMongeArrays(t *testing.T) {
	for _, c := range mongeArraysCases {
		if !isMongeArrays(c.a) {
			printMongeArrays(t, "matrix A", c.a)
			t.Errorf("is not monge")
			continue
		}
		var rows []int
		for i := 0; i < len(c.a); i++ {
			rows = append(rows, i)
		}
		got := findMinimums(c.a, rows)
		if !reflect.DeepEqual(got, c.want) {
			printMongeArrays(t, "monge arrays", c.a)
			t.Errorf("         got %v", got)
			t.Errorf("        want %v", c.want)
		}
	}
}
