package pathfind

type Node interface {
	Edges() []Edge
}

type Edge interface {
	Nodes() (from, to Node)
	Cost() float64
}
