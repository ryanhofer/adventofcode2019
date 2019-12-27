package main

import (
	"unicode"

	"github.com/ryanhofer/adventofcode2019/input"
)

const (
	Wall     rune = '#'
	Open     rune = '.'
	Entrance rune = '@'
)

func main() {
	for line := range input.Lines() {
		for _, r := range line {
			switch true {
			case unicode.IsLower(r):
				// key
			case unicode.IsLetter(r):
				// door
			case r == Wall:
			case r == Open:
			case r == Entrance:
			}
		}
	}
}
