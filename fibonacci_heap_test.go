package clrs

import (
	"fmt"
	"testing"
)

func buildTestFibHeap() *fibHeap {
	h := new(fibHeap)
	h.n = 15

	n23 := new(fibNode)
	n23.key = 23

	n7 := new(fibNode)
	n7.key = 7

	n21 := new(fibNode)
	n21.key = 21

	n3 := new(fibNode)
	n3.key = 3
	n3.degree = 3

	n17 := new(fibNode)
	n17.key = 17
	n17.degree = 1

	n24 := new(fibNode)
	n24.key = 24
	n24.degree = 2

	n23.left = n24
	n23.right = n7
	n7.left = n23
	n7.right = n21
	n21.left = n7
	n21.right = n3
	n3.left = n21
	n3.right = n17
	n17.left = n3
	n17.right = n24
	n24.left = n17
	n24.right = n23

	h.min = n3

	n18 := new(fibNode)
	n18.key = 18
	n18.p = n3
	n18.mark = true
	n18.degree = 1

	n52 := new(fibNode)
	n52.key = 52
	n52.p = n3

	n38 := new(fibNode)
	n38.key = 38
	n38.p = n3
	n38.degree = 1

	n18.left = n38
	n18.right = n52
	n52.left = n18
	n52.right = n38
	n38.left = n52
	n38.right = n18

	n3.child = n18

	n30 := new(fibNode)
	n30.key = 30
	n30.p = n17
	n30.left = n30
	n30.right = n30

	n17.child = n30

	n26 := new(fibNode)
	n26.key = 26
	n26.p = n24
	n26.mark = true
	n26.degree = 1

	n46 := new(fibNode)
	n46.key = 46
	n46.p = n24

	n26.left = n46
	n26.right = n46
	n46.left = n26
	n46.right = n26

	n24.child = n26

	n39 := new(fibNode)
	n39.key = 39
	n39.p = n18
	n39.mark = true

	n39.left = n39
	n39.right = n39
	n18.child = n39

	n41 := new(fibNode)
	n41.key = 41
	n41.p = n38

	n41.left = n41
	n41.right = n41
	n38.child = n41

	n35 := new(fibNode)
	n35.key = 35
	n35.p = n26

	n35.left = n35
	n35.right = n35
	n26.child = n35

	return h
}

func e(x *fibNode) int {
	if x != nil {
		return x.key
	}
	return -1
}

func TestFibHeapExtractMin(t *testing.T) {
	h := buildTestFibHeap()
	min := h.fibHeapExtractMin()
	if min.key != 3 {
		t.Errorf(" got %v", min.key)
		t.Errorf("want %v", 3)
	}
	p := func(p, c, l, r, d int, m bool) string {
		return fmt.Sprintf("p:%v child:%v left:%v right:%v degree:%v mark:%v", p, c, l, r, d, m)
	}
	want := map[int]string{
		7:  p(-1, 23, 38, 18, 3, false),
		18: p(-1, 39, 7, 38, 2, true),
		38: p(-1, 41, 18, 7, 1, false),
		24: p(7, 26, 17, 23, 2, false),
		17: p(7, 30, 23, 24, 1, false),
		23: p(7, -1, 24, 17, 0, false),
		21: p(18, 52, 39, 39, 1, false),
		39: p(18, -1, 21, 21, 0, true),
		41: p(38, -1, 41, 41, 0, false),
		26: p(24, 35, 46, 46, 1, true),
		46: p(24, -1, 26, 26, 0, false),
		30: p(17, -1, 30, 30, 0, false),
		52: p(21, -1, 52, 52, 0, false),
		35: p(26, -1, 35, 35, 0, false),
	}
	w := func(x *fibNode) int {
		if x != nil {
			return x.key
		}
		return -1
	}
	nodes := h.fibHeapWalk()
	for _, n := range nodes {
		got := p(w(n.p), w(n.child), w(n.left), w(n.right), n.degree, n.mark)
		if got != want[n.key] {
			t.Errorf("node %v", n.key)
			t.Errorf(" got %v", got)
			t.Errorf("want %v", want[n.key])
		}
	}
}
