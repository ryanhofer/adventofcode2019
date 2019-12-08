package main

import "math"

type Target struct{ x, y int }

type ByAngle []Target

func (s ByAngle) Len() int {
	return len(s)
}

func (s ByAngle) Less(i, j int) bool {
	xi, yi := float64(s[i].x), float64(s[i].y)
	xj, yj := float64(s[j].x), float64(s[j].y)

	// Rotate 90deg clockwise
	xi, yi = yi, -xi
	xj, yj = yj, -xj

	ti := math.Atan2(yi, xi)
	tj := math.Atan2(yj, xj)
	return ti < tj
}

func (s ByAngle) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
