package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
)

func main() {
	plots := []Plot{}
	for line := range input.Lines() {
		plots = append(plots, Draw(Parse(line)))
	}

	if len(plots) < 2 {
		panic("expected 2 plots")
	}

	intersect := Intersect(plots[0], plots[1])
	bestManhattan := 0
	bestSteps := 0

	for p, steps := range intersect {
		if p.IsOrigin() {
			continue
		}

		manhattan := Origin.ManhattanDistance(p)
		if bestManhattan == 0 || manhattan < bestManhattan {
			bestManhattan = manhattan
		}

		if bestSteps == 0 || steps < bestSteps {
			bestSteps = steps
		}
	}

	fmt.Println("Part 1:", bestManhattan)
	fmt.Println("Part 2:", bestSteps)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
