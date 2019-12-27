package pathfind

type Node interface {
	Edges() []Edge
}

type Edge interface {
	Nodes() (Node, Node)
	Cost() float64
}
