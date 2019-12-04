package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
)

func main() {
	program := Parse(input.Contents())

	fmt.Println("Part 1:", Exec(program, 12, 2))

	noun, verb := search(program, 19690720)
	fmt.Println("Part 2:", 100*noun+verb)
}

func search(program []int, needle int) (noun, verb int) {
search:
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			if needle == Exec(program, noun, verb) {
				break search
			}
		}
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
