package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

func main() {
	var err error
	var cfg *intcode.Config
	var in, out chan intcode.Word

	program, err := intcode.Parse(input.Contents())
	check(err)

	in = make(chan intcode.Word, 10)
	out = make(chan intcode.Word, 10)
	cfg = &intcode.Config{
		Input:  in,
		Output: out,
	}
	in <- 1
	_, err = intcode.Exec(program, cfg)
	check(err)

	var finalOutput intcode.Word
	for finalOutput = range out {
	}
	fmt.Println("Part 1:", finalOutput)

	in = make(chan intcode.Word, 10)
	out = make(chan intcode.Word, 10)
	cfg = &intcode.Config{
		Input:  in,
		Output: out,
	}
	in <- 5
	_, err = intcode.Exec(program, cfg)
	check(err)

	for finalOutput = range out {
	}
	fmt.Println("Part 2:", finalOutput)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
