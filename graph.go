package clrs

const (
	// WHITE ...
	WHITE = 0
	// GRAY ...
	GRAY = 1
	// SCHWARZ ...
	SCHWARZ = 2
)

type vertex struct {
	color int
	d     int
	p     *vertex
	key   rune
}

type adjNode struct {
	v    *vertex
	next *adjNode
}

type graph struct {
	adj map[*vertex]*adjNode
}
