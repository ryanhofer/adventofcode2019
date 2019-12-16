package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

var w *bufio.Writer

var maze map[Coord]Tile
var pos Coord
var moves []Dir
var unexplored map[Coord]bool

var in chan<- intcode.Word
var out <-chan intcode.Word

func main() {
	w = bufio.NewWriter(os.Stdout)

	program, err := intcode.Parse(input.Contents())
	check(err)

	in, out, _ = intcode.Spawn(program)

	maze = map[Coord]Tile{}
	pos = Coord{0, 0}
	unexplored = map[Coord]bool{}

	for {
		// Check adjacent coords for unexplored tiles
		candidates := []Dir{}
		for _, dir := range dirs {
			c := pos.Translate(dir)
			if _, ok := maze[c]; !ok {
				unexplored[c] = true
				candidates = append(candidates, dir)
			}
		}

		// Watch that repair droid go!
		// printMaze()
		// time.Sleep(50 * time.Millisecond)

		if len(unexplored) == 0 {
			break
		}

		if len(candidates) == 0 {
			// Dead end; time to backtrack
			backtrack()
			continue
		}

		move(candidates[0])
	}

	printMaze()

	start := Coord{0, 0}
	goal := findGoalCoord()
	path, err := pathfind(start, goal)
	check(err)
	fmt.Println("Part 1:", len(path)-1)

	// Find furthest point from oxygen system
	var farthest int
	for c, t := range maze {
		if t != Traversable {
			continue
		}
		path, err := pathfind(goal, c)
		check(err)
		distance := len(path) - 1
		if distance > farthest {
			farthest = distance
		}
	}
	fmt.Println("Part 2:", farthest)
}

func findGoalCoord() Coord {
	for c, t := range maze {
		if t == Goal {
			return c
		}
	}
	panic("goal not found")
}

func move(dir Dir) {
	c := pos.Translate(dir)

	in <- intcode.Word(dir)
	t := Tile(<-out)

	maze[c] = t
	delete(unexplored, c)

	switch t {
	case Goal:
		fallthrough
	case Traversable:
		pos = pos.Translate(dir)
		moves = append(moves, dir)
	}
}

func backtrack() {
	n := len(moves) - 1
	dir := moves[n].Opposite()
	moves = moves[:n]

	c := pos.Translate(dir)

	in <- intcode.Word(dir)
	t := Tile(<-out)
	if t == Wall {
		panic(fmt.Sprintf("unexpected wall during backtrack at %d,%d", c.X, c.Y))
	}

	pos = pos.Translate(dir)
}

func printMaze() {
	var min, max Coord
	for c := range maze {
		if c.X < min.X {
			min.X = c.X
		}
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
	}

	for y := min.Y - 1; y <= max.Y+1; y++ {
		for x := min.X - 1; x <= max.X+1; x++ {
			c := Coord{x, y}
			if c == pos {
				w.WriteRune('D')
			} else if t, ok := maze[c]; ok {
				w.WriteString(t.String())
			} else {
				w.WriteRune(' ')
			}
		}
		w.WriteRune('\n')
	}
	w.Flush()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
