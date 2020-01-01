package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"github.com/ryanhofer/adventofcode2019/geom"
)

type pathfinder struct {
	m     maze
	pos   geom.Vec
	steps []geom.Vec
	keys  string
}

func (p *pathfinder) clone() *pathfinder {
	cloned := pathfinder{
		m:     p.m,
		pos:   p.pos,
		steps: p.steps,
		keys:  p.keys,
	}
	return &cloned
}

func (p *pathfinder) passable(t tile) bool {
	if t == Wall {
		return false
	}
	if t.isDoor() {
		k := unicode.ToLower(rune(t))
		return strings.ContainsRune(p.keys, k)
	}
	return true
}

func (p *pathfinder) solve(keys map[rune]geom.Vec) (int, bool) {
	if len(keys) == len(p.keys) {
		// Found all the keys
		fmt.Println("DONE:", len(p.steps))
		return len(p.steps), true
	}

	var bestSteps *int

	for k, c := range keys {
		// TODO: Do we already know the best steps from this state?
		score, ok := Cached(p.pos, p.keys)

		// Has pathfinder already collected this key?
		if strings.ContainsRune(p.keys, k) {
			continue
		}
		// Can pathfinder reach this key?
		path, ok := p.search(c)
		if !ok {
			continue
		}
		// Got the key
		q := p.clone()
		q.keys += string(k)
		q.pos = path[len(path)-1]
		q.steps = append(p.steps, path[1:]...)

		steps, ok := q.solve(keys)
		if !ok {
			continue
		}
		if bestSteps == nil || steps < *bestSteps {
			bestSteps = &steps
		}
	}

	if bestSteps == nil {
		return 0, false
	}
	return *bestSteps, true
}

// https://en.wikipedia.org/wiki/A*_search_algorithm#Pseudocode
func (p *pathfinder) search(goal geom.Vec) (path []geom.Vec, ok bool) {
	start := p.pos
	h := manhattanDistance

	cameFrom := map[geom.Vec]geom.Vec{}
	openSet := map[geom.Vec]bool{
		start: true,
	}
	gScore := map[geom.Vec]float64{
		start: 0,
	}
	fScore := map[geom.Vec]float64{
		start: h(start, goal),
	}

	for len(openSet) > 0 {
		var current geom.Vec
		bestFScore := math.Inf(1)
		for c := range openSet {
			f, ok := fScore[c]
			if ok && f < bestFScore {
				bestFScore = f
				current = c
			}
		}

		if current == goal {
			return reconstructPath(cameFrom, current), true
		}

		delete(openSet, current)

		for _, dir := range geom.Cardinals {
			neighbor := current.Add(dir.Vec())
			t, ok := p.m[neighbor]
			if !ok || !p.passable(t) {
				continue
			}
			if t.isKey() && neighbor != goal {
				// Don't want to to pick up the wrong key.
				continue
			}

			tentativeGScore := gScore[current] + 1

			if g, ok := gScore[neighbor]; !ok || tentativeGScore < g {
				cameFrom[neighbor] = current

				gScore[neighbor] = tentativeGScore
				fScore[neighbor] = tentativeGScore + h(neighbor, goal)

				if _, ok := openSet[neighbor]; !ok {
					openSet[neighbor] = true
				}
			}
		}
	}

	return nil, false
}

func reconstructPath(cameFrom map[geom.Vec]geom.Vec, current geom.Vec) []geom.Vec {
	path := []geom.Vec{current}
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

func manhattanDistance(from, to geom.Vec) float64 {
	diff := from.Add(to.Negate())
	return math.Abs(float64(diff.X)) + math.Abs(float64(diff.Y))
}
