package main

import (
	"unicode"

	. "github.com/ryanhofer/adventofcode2019/geom"
)

type maze map[Vec]tile

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
