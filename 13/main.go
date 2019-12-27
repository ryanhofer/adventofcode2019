package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/arcade"
	"github.com/ryanhofer/adventofcode2019/imath"
	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

func main() {
	part1()
	part2()
}

func part1() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	_, out, _ := intcode.Spawn(program)

	screen := arcade.Screen{}
	for output := range out {
		x := int(output)
		y := int(<-out)
		tile := arcade.Tile(<-out)
		screen.SetTile(x, y, tile)
	}

	numBlocks := 0
	for _, t := range screen {
		if t == arcade.Block {
			numBlocks++
		}
	}
	fmt.Println("Part 1:", numBlocks)
}

func part2() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	// Insert quarters
	program[0] = 2

	in, out, halt := intcode.Spawn(program)

	var ballX, paddleX int
	var nextInput intcode.Word

	screen := arcade.Screen{}
	var score intcode.Word

gameloop:
	for {
		var output intcode.Word

		select {
		case err := <-halt:
			check(err)
			break gameloop
		case in <- nextInput:
			continue gameloop
		case output = <-out:
		}

		if output == -1 {
			// Score instruction
			_, score = <-out, <-out
			// screen.Print()
			// fmt.Println("SCORE:", score)
			continue
		}

		// Tile instruction
		x := int(output)
		y := int(<-out)
		tile := arcade.Tile(<-out)
		screen.SetTile(x, y, tile)

		switch tile {
		case arcade.Ball:
			ballX = x
			nextInput = intcode.Word(imath.Cmp(ballX, paddleX))
		case arcade.HorizontalPaddle:
			paddleX = x
			nextInput = intcode.Word(imath.Cmp(ballX, paddleX))
		}
	}

	fmt.Println("Part 2:", score)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
