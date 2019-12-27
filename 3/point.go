package main

import "github.com/ryanhofer/adventofcode2019/imath"

type Point struct {
	X int
	Y int
}

var Origin Point = Point{
	X: 0,
	Y: 0,
}

func (p Point) IsOrigin() bool {
	return p.Equals(Origin)
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) ManhattanDistance(q Point) int {
	return imath.Abs(p.X-q.X) + imath.Abs(p.Y-q.Y)
}
