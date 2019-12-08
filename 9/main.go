package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

func main() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	var in, out chan intcode.Word
	var output intcode.Word

	in = make(chan intcode.Word)
	out = make(chan intcode.Word)

	go func() {
		cfg := &intcode.Config{
			Input:  in,
			Output: out,
		}
		_, err = intcode.Exec(program, cfg)
		check(err)
	}()

	in <- 1 // test mode
	for output = range out {
	}

	fmt.Println("Part 1:", output)

	in = make(chan intcode.Word)
	out = make(chan intcode.Word)

	go func() {
		cfg := &intcode.Config{
			Input:  in,
			Output: out,
		}
		_, err = intcode.Exec(program, cfg)
		check(err)
	}()

	in <- 2 // sensor boost mode
	for output = range out {
	}

	fmt.Println("Part 2:", output)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
