package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

type Robot struct {
	pos Coord
	dir Direction
}

type Coord struct {
	X, Y int
}

type Color int

const (
	Black Color = 0
	White Color = 1
)

var hull map[Coord]Color
var debug bool

func main() {
	hull = map[Coord]Color{}
	startHullPaintingRobot()
	fmt.Println("Part 1:", len(hull))

	hull = map[Coord]Color{}
	hull[Coord{0, 0}] = White
	startHullPaintingRobot()
	fmt.Println("Part 2:")
	printTheHull()
}

func startHullPaintingRobot() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	in, out, halt := intcode.Spawn(program)

	bot := Robot{
		pos: Coord{0, 0},
		dir: Up,
	}

	for {
		select {
		case err := <-halt:
			check(err)
			return
		case in <- intcode.Word(hull[bot.pos]):
		}

		// repaint it
		newColor := Color(<-out)
		hull[bot.pos] = newColor

		if debug {
			fmt.Printf("PAINT [%d,%d]\n", bot.pos.X, bot.pos.Y)
		}

		// rotate
		switch <-out {
		case 0:
			bot.dir = bot.dir.CCW()
		case 1:
			bot.dir = bot.dir.CW()
		}

		// move
		switch bot.dir {
		case Up:
			bot.pos.Y--
		case Right:
			bot.pos.X++
		case Down:
			bot.pos.Y++
		case Left:
			bot.pos.X--
		}
	}
}

func printTheHull() {
	// find the bounds of the painted region
	var min, max Coord
	for c := range hull {
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
	}

	// print it
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			coord := Coord{x, y}
			color := hull[coord]
			switch color {
			case Black:
				fmt.Print(" ")
			case White:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
