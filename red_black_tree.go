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

type rbdata struct {
	key   int
	color bool
}

type rbtree struct {
	root *rbnode
}

func init() {
	nilNode = new(rbnode)
	nilNode.color = BLACK
}

func (t *rbtree) toArray() []rbdata {
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
	a := make([]rbdata, len(nodes))
	for i, node := range nodes {
		a[i] = rbdata{}
		if (node != nil) && (node != nilNode) {
			a[i].key = node.key
			a[i].color = node.color
		} else {
			a[i].key = -1
			a[i].color = BLACK
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

func (t *rbtree) rbInsert(z *rbnode) {
	y := nilNode
	x := t.root
	for x != nilNode {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nilNode {
		t.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	z.left = nilNode
	z.right = nilNode
	z.color = RED
	t.rbInsertFixup(z)
}

func (t *rbtree) rbInsertFixup(z *rbnode) {
	for z.p.color == RED {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					t.leftRotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				t.rightRotate(z.p.p)
			}
		} else {
			y := z.p.p.left
			if y.color == RED {
				z.p.color = BLACK
				y.color = BLACK
				z.p.p.color = RED
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					t.rightRotate(z)
				}
				z.p.color = BLACK
				z.p.p.color = RED
				t.leftRotate(z.p.p)
			}
		}
	}
	t.root.color = BLACK
}
