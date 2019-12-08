package main

import "github.com/ryanhofer/adventofcode2019/input"

func loadInput() {
	asteroids = [height][width]bool{}
	x, y := 0, 0
	for line := range input.Lines() {
		x = 0
		for _, r := range line {
			if r == '#' {
				asteroids[y][x] = true
			}
			x++
			if x >= width {
				break
			}
		}
		y++
		x = 0
		if y >= height {
			break
		}
	}
}
