package clrs

import (
	"math"
)

const (
	// RED ...
	RED = true
	// BLACK ...
	BLACK = false
)

var (
	nilNode *rbnode
)

type rbnode struct {
	color bool
	key   int
	left  *rbnode
	right *rbnode
	p     *rbnode
}

type rbtree struct {
	root *rbnode
}

func (t *rbtree) toArray() []int {
	var nodes []*rbnode
	layer := 0
	for {
		currentLayerNodes := int(math.Pow(2.0, float64(layer)))
		down := false
		for i := currentLayerNodes - 1; i <= 2*currentLayerNodes-2; i++ {
			if i == 0 {
				if t.root != nil {
					down = true
					nodes = append(nodes, t.root)
				}
			} else {
				p := (i - 1) / 2
				isLeft := i == (2*p + 1)
				if nodes[p] == nil || nodes[p] == nilNode {
					nodes = append(nodes, nil)
				} else {
					var node *rbnode
					if isLeft {
						node = nodes[p].left
					} else {
						node = nodes[p].right
					}
					if (node != nil) && (node != nilNode) {
						down = true
					}
					nodes = append(nodes, node)
				}
			}
		}
		if !down {
			nodes = nodes[:currentLayerNodes-1]
			break
		}
		layer++
	}
	a := make([]int, len(nodes))
	for i, node := range nodes {
		if (node != nil) && (node != nilNode) {
			a[i] = node.key
		} else {
			a[i] = -1
		}
	}
	return a
}

func (t *rbtree) leftRotate(x *rbnode) {
	y := x.right
	x.right = y.left
	if y.left != nilNode {
		y.left.p = x
	}
	y.p = x.p
	if x.p == nilNode {
		t.root = y
	} else if x == x.p.left {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
}

func (t *rbtree) rightRotate(y *rbnode) {
	x := y.left
	y.left = x.right
	if x.right != nilNode {
		x.right.p = y
	}
	x.p = y.p
	if y.p == nilNode {
		t.root = x
	} else if y == y.p.left {
		y.p.left = x
	} else {
		y.p.right = x
	}
	x.right = y
	y.p = x
}
