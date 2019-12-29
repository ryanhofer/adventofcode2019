package main

import (
	"fmt"
	"unicode"

	"github.com/ryanhofer/adventofcode2019/geom"
	"github.com/ryanhofer/adventofcode2019/input"
)

type tile rune

const (
	Wall     tile = '#'
	Open     tile = '.'
	Entrance tile = '@'
)

func (t tile) isKey() bool {
	return unicode.IsLower(rune(t))
}

func (t tile) isDoor() bool {
	return unicode.IsUpper(rune(t))
}

var maze map[geom.Vec]tile

func main() {
	maze = map[geom.Vec]tile{}

	keys := map[rune]geom.Vec{}
	doors := map[rune]geom.Vec{}

	var c, entrance geom.Vec
	for line := range input.Lines() {
		for _, r := range line {
			t := tile(r)
			switch true {
			case unicode.IsLower(r):
				keys[r] = c
			case unicode.IsUpper(r):
				doors[unicode.ToLower(r)] = c
			case t == Entrance:
				entrance = c
			case t == Open:
			case t == Wall:
			}
			maze[c] = t
			c.X++
		}
		c.X = 0
		c.Y++
	}

	fmt.Printf("entrance=%s\n", entrance)
}
