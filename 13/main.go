package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/arcade"
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

	out := make(chan intcode.Word)
	_ = startProgram(program, nil, out)

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

	in := make(chan intcode.Word)
	out := make(chan intcode.Word)
	done := startProgram(program, in, out)

	var ballX, paddleX int
	var nextInput intcode.Word

	screen := arcade.Screen{}
	var score intcode.Word

gameloop:
	for {
		var output intcode.Word

		select {
		case <-done:
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
			nextInput = intcode.Word(cmp(ballX, paddleX))
		case arcade.HorizontalPaddle:
			paddleX = x
			nextInput = intcode.Word(cmp(ballX, paddleX))
		}
	}

	fmt.Println("Part 2:", score)
}

func startProgram(program intcode.Program, in, out chan intcode.Word) <-chan bool {
	done := make(chan bool)

	go func(in <-chan intcode.Word, out chan<- intcode.Word) {
		cfg := &intcode.Config{
			Input:  in,
			Output: out,
		}
		_, err := intcode.Exec(program, cfg)
		check(err)
		done <- true
	}(in, out)

	return done
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
