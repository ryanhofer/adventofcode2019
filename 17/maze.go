package main

import (
	. "github.com/ryanhofer/adventofcode2019/geom"
)

type Maze map[Vec]rune

const (
	Scaffold     rune = '#'
	OpenSpace    rune = '.'
	RobotN       rune = '^'
	RobotE       rune = '>'
	RobotS       rune = 'v'
	RobotW       rune = '<'
	RobotFalling rune = 'X'
)

func (m Maze) IsIntersection(v Vec) bool {
	if m[v] != Scaffold {
		return false
	}
	for _, dir := range Cardinals {
		w := v.Add(dir.Vec())
		if m[w] != Scaffold {
			return false
		}
	}
	return true
}
