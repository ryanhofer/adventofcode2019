package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

func main() {
	var err error
	var cfg *intcode.Config
	var output []int

	program, err := intcode.Parse(input.Contents())
	check(err)

	cfg = &intcode.Config{Input: 1}
	_, output, err = intcode.Exec(program, cfg)
	check(err)
	fmt.Println("Part 1:", output[len(output)-1])

	cfg = &intcode.Config{Input: 5}
	_, output, err = intcode.Exec(program, cfg)
	check(err)
	fmt.Println("Part 2:", output[len(output)-1])
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
