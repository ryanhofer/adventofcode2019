package main

import (
	"fmt"
	"strconv"

	"github.com/ryanhofer/adventofcode2019/input"
)

func main() {
	total := 0
	totalIncludingFuel := 0

	for line := range input.Lines() {
		mass, err := strconv.Atoi(line)
		check(err)

		total += mass2fuel(mass)

		for mass > 0 {
			fuel := mass2fuel(mass)
			totalIncludingFuel += fuel
			mass = fuel
		}
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", totalIncludingFuel)
}

func mass2fuel(mass int) int {
	fuel := (mass / 3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
