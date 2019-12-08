package main

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) CCW() Direction {
	return (d - 1 + 4) % 4
}

func (d Direction) CW() Direction {
	return (d + 1 + 4) % 4
}
