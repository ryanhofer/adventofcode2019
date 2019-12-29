package main

import (
	"errors"

	. "github.com/ryanhofer/adventofcode2019/geom"
)

var ErrPathNotFound = errors.New("path not found")

// https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode
func pathfind(start Vec, goal Vec) ([]Vec, error) {
	// Using Manhattan distance as the heuristic function
	h := func(c Vec) int {
		return abs(c.X-goal.X) + abs(c.Y-goal.Y)
	}

	cameFrom := map[Vec]Vec{}
	openSet := map[Vec]bool{
		start: true,
	}
	gScore := map[Vec]int{
		start: 0,
	}
	fScore := map[Vec]int{
		start: h(start),
	}

	for len(openSet) > 0 {
		var current Vec
		bestFScore := 1000
		for c := range openSet {
			f, ok := fScore[c]
			if ok && f < bestFScore {
				bestFScore = f
				current = c
			}
		}

		if current == goal {
			return reconstructPath(cameFrom, current), nil
		}

		delete(openSet, current)

		for _, dir := range Cardinals {
			neighbor := current.Add(dir.Vec())
			if t, ok := maze[neighbor]; !ok || t == Wall {
				continue
			}

			tentativeGScore := gScore[current] + 1

			if g, ok := gScore[neighbor]; !ok || tentativeGScore < g {
				cameFrom[neighbor] = current

				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + h(neighbor)

				if _, ok := openSet[neighbor]; !ok {
					openSet[neighbor] = true
				}
			}
		}
	}

	return nil, ErrPathNotFound
}

func reconstructPath(cameFrom map[Vec]Vec, current Vec) []Vec {
	path := []Vec{current}
	for {
		next, ok := cameFrom[current]
		if !ok {
			break
		}
		current = next
		// Prepending is expensive; just append and reverse it later
		path = append(path, current)
	}

	// Reverse it
	for i := 0; i < len(path)/2; i++ {
		j := len(path) - 1 - i
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
