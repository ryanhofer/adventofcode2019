package main

import "fmt"

type Dir int

const (
	_ Dir = iota
	North
	South
	West
	East
)

var dirs = []Dir{
	North,
	East,
	South,
	West,
}

func (d Dir) Opposite() Dir {
	switch d {
	case North:
		return South
	case South:
		return North
	case West:
		return East
	case East:
		return West
	default:
		panic(fmt.Sprintf("unexpected dir: %d", d))
	}
}

func (d Dir) String() string {
	switch d {
	case North:
		return "N"
	case South:
		return "S"
	case West:
		return "W"
	case East:
		return "E"
	default:
		panic(fmt.Sprintf("unexpected dir: %d", d))
	}
}
