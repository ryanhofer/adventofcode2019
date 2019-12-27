package main

import (
	"fmt"
	"sort"

	"github.com/ryanhofer/adventofcode2019/imath"
)

const width = 26
const height = 26

var asteroids [height][width]bool

func main() {
	loadInput()

	bestCount := 0
	bestX, bestY := 0, 0
	for y := 0; y < height; y++ {
		for x := 0; x < height; x++ {
			if !asteroids[y][x] {
				continue
			}
			count := countVisible(x, y)
			// fmt.Printf("[%2d,%2d] -> %d\n", x, y, count)
			if count > bestCount {
				bestCount = count
				bestX, bestY = x, y
			}
		}
	}
	fmt.Printf("BEST [%2d,%2d] -> %d\n", bestX, bestY, bestCount)
	printVisible(bestX, bestY)
	fmt.Println("Part 1:", bestCount)

	targets := make([]Target, 0, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if x == bestX && y == bestY {
				continue
			}
			if asteroids[y][x] {
				t := Target{
					x: x - bestX,
					y: y - bestY,
				}
				targets = append(targets, t)
			}
		}
	}

	sort.Sort(ByAngle(targets))
	// for _, t := range targets {
	// 	fmt.Printf("[%3d,%3d]\n", t.x, t.y)
	// }

	numVaporized := 0
	for _, t := range targets {
		if t.x == 0 && t.y == 0 {
			// already vaporized
			continue
		}
		toX, toY := bestX+t.x, bestY+t.y
		if checkVisible(bestX, bestY, toX, toY) {
			t.x, t.y = 0, 0
			numVaporized++
			// fmt.Printf("VAPORIZED #%d at [%3d,%3d]\n", numVaporized, toX, toY)
		}
		if numVaporized == 200 {
			fmt.Println("Part 2:", (100*toX + toY))
			return
		}
	}
}

func printVisible(fromX, fromY int) {
	const (
		emptyRune   = ' '
		hiddenRune  = '∙'
		visibleRune = 'o'
		stationRune = 'X'
	)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !asteroids[y][x] {
				fmt.Print(string(emptyRune))
				continue
			}
			if x == fromX && y == fromY {
				fmt.Print(string(stationRune))
			} else if checkVisible(fromX, fromY, x, y) {
				fmt.Print(string(visibleRune))
			} else {
				fmt.Print(string(hiddenRune))
			}
		}
		fmt.Println()
	}
}

func countVisible(fromX, fromY int) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if !asteroids[y][x] {
				continue
			}
			if checkVisible(fromX, fromY, x, y) {
				count++
			}
		}
	}
	return count
}

func checkVisible(fromX, fromY, toX, toY int) bool {
	if fromX == toX && fromY == toY {
		return false
	}

	dx, dy := norm(toX-fromX, toY-fromY)
	x, y := fromX+dx, fromY+dy

	for !(x == toX && y == toY) {
		if asteroids[y][x] {
			return false
		}
		x, y = x+dx, y+dy
	}
	return true
}

func norm(x, y int) (int, int) {
	if x == 0 && y == 0 {
		return x, y
	}
	if x == 0 {
		return x, y / imath.Abs(y)
	}
	if y == 0 {
		return x / imath.Abs(x), y
	}
	d := imath.GCD(imath.Abs(x), imath.Abs(y))
	return x / d, y / d
}
