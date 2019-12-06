package main

import (
	"fmt"

	"github.com/gobuffalo/nulls"
	"github.com/ryanhofer/adventofcode2019/input"
	"github.com/ryanhofer/adventofcode2019/intcode"
)

func main() {
	program, err := intcode.Parse(input.Contents())
	check(err)

	cfg := &intcode.Config{
		Noun: nulls.NewInt(12),
		Verb: nulls.NewInt(2),
	}
	exit, _, err := intcode.Exec(program, cfg)
	check(err)

	fmt.Println("Part 1:", exit)

	noun, verb := search(program, 19690720)
	fmt.Println("Part 2:", 100*noun+verb)
}

func search(program []int, needle int) (noun, verb int) {
search:
	for noun = 0; noun <= 99; noun++ {
		for verb = 0; verb <= 99; verb++ {
			cfg := &intcode.Config{
				Noun: nulls.NewInt(noun),
				Verb: nulls.NewInt(verb),
			}
			exit, _, err := intcode.Exec(program, cfg)
			check(err)

			if exit == needle {
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
