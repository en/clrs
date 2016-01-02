package clrs

import (
	"reflect"
	"strings"
	"testing"
)

var multiplyCases = []struct {
	a    [][]int
	b    [][]int
	want [][]int
}{
	{
		[][]int{
			[]int{1, 3},
			[]int{7, 5},
		},
		[][]int{
			[]int{6, 8},
			[]int{4, 2},
		},
		[][]int{
			[]int{18, 14},
			[]int{62, 66},
		},
	},
	{
		[][]int{
			[]int{1, 2, 3, 4},
			[]int{5, 6, 7, 8},
			[]int{9, 10, 11, 12},
			[]int{13, 14, 15, 16},
		},
		[][]int{
			[]int{17, 18, 19, 20},
			[]int{21, 22, 23, 24},
			[]int{25, 26, 27, 28},
			[]int{29, 30, 31, 32},
		},
		[][]int{
			[]int{125 * 2, 130 * 2, 135 * 2, 140 * 2},
			[]int{309 * 2, 322 * 2, 335 * 2, 348 * 2},
			[]int{493 * 2, 514 * 2, 535 * 2, 556 * 2},
			[]int{677 * 2, 706 * 2, 735 * 2, 764 * 2},
		},
	},
	{
		[][]int{
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
			[]int{9, 10, 11, 12, 13, 14, 15, 16},
			[]int{17, 18, 19, 20, 21, 22, 23, 24},
			[]int{25, 26, 27, 28, 29, 30, 31, 32},
			[]int{33, 34, 35, 36, 37, 38, 39, 40},
			[]int{41, 42, 43, 44, 45, 46, 47, 48},
			[]int{49, 50, 51, 52, 53, 54, 55, 56},
			[]int{57, 58, 59, 60, 61, 62, 63, 64},
		},
		[][]int{
			[]int{65, 66, 67, 68, 69, 70, 71, 72},
			[]int{73, 74, 75, 76, 77, 78, 79, 80},
			[]int{81, 82, 83, 84, 85, 86, 87, 88},
			[]int{89, 90, 91, 92, 93, 94, 95, 96},
			[]int{97, 98, 99, 100, 101, 102, 103, 104},
			[]int{105, 106, 107, 108, 109, 110, 111, 112},
			[]int{113, 114, 115, 116, 117, 118, 119, 120},
			[]int{121, 122, 123, 124, 125, 126, 127, 128},
		},
		[][]int{
			[]int{3684, 3720, 3756, 3792, 3828, 3864, 3900, 3936},
			[]int{9636, 9736, 9836, 9936, 10036, 10136, 10236, 10336},
			[]int{15588, 15752, 15916, 16080, 16244, 16408, 16572, 16736},
			[]int{21540, 21768, 21996, 22224, 22452, 22680, 22908, 23136},
			[]int{27492, 27784, 28076, 28368, 28660, 28952, 29244, 29536},
			[]int{33444, 33800, 34156, 34512, 34868, 35224, 35580, 35936},
			[]int{39396, 39816, 40236, 40656, 41076, 41496, 41916, 42336},
			[]int{45348, 45832, 46316, 46800, 47284, 47768, 48252, 48736},
		},
	},
}

func printMatrix(t *testing.T, label string, m [][]int) {
	t.Errorf("%v", label)
	spaces := strings.Repeat(" ", len(label))
	for _, r := range m {
		t.Errorf("%v%v", spaces, r)
	}
}

func TestSquareMatrixMultiply(t *testing.T) {
	for _, c := range multiplyCases {
		got := squareMatrixMultiply(c.a, c.b)
		if !reflect.DeepEqual(got, c.want) {
			printMatrix(t, "matrix A", c.a)
			printMatrix(t, "matrix B", c.b)
			printMatrix(t, "     got", got)
			printMatrix(t, "    want", c.want)
		}
	}
}

func TestSquareMatrixMultiplyRecursive(t *testing.T) {
	for _, c := range multiplyCases {
		n := len(c.a)
		got := squareMatrixMultiplyRecursive(c.a, c.b, mat{0, n - 1, 0, n - 1}, mat{0, n - 1, 0, n - 1})
		if !reflect.DeepEqual(got, c.want) {
			printMatrix(t, "matrix A", c.a)
			printMatrix(t, "matrix B", c.b)
			printMatrix(t, "     got", got)
			printMatrix(t, "    want", c.want)
		}
	}
}

func TestSquareMatrixMultiplyStrassen(t *testing.T) {
	for _, c := range multiplyCases {
		n := len(c.a)
		got := squareMatrixMultiplyStrassen(c.a, c.b, mat{0, n - 1, 0, n - 1}, mat{0, n - 1, 0, n - 1})
		if !reflect.DeepEqual(got, c.want) {
			printMatrix(t, "matrix A", c.a)
			printMatrix(t, "matrix B", c.b)
			printMatrix(t, "     got", got)
			printMatrix(t, "    want", c.want)
		}
	}
}
