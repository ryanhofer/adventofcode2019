package main

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
	return abs(p.X-q.X) + abs(p.Y-q.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
