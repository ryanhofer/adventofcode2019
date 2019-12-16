package main

type Tile int

const (
	Wall Tile = iota
	Traversable
	Goal
)

func (t Tile) String() string {
	switch t {
	case Wall:
		return "█"
	case Traversable:
		return "░"
	case Goal:
		return "X"
	default:
		return "?"
	}
}
