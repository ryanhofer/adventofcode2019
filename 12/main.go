package main

import (
	"fmt"

	"github.com/ryanhofer/adventofcode2019/input"
)

var moons []Moon

func main() {
	moons = []Moon{}

	for line := range input.Lines() {
		pos := vec3{}
		_, err := fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &pos.x, &pos.y, &pos.z)
		check(err)
		moon := Moon{
			pos: pos,
			vel: vec3{},
		}
		moons = append(moons, moon)
	}

	runSimulation(1000)
	fmt.Println("Part 1:", totalEnergy())

	steps := findCycleLength()
	fmt.Println("Part 2:", steps)
}

func printState(t int) {
	fmt.Printf("t=%d\n", t)
	for i := range moons {
		fmt.Println(moons[i])
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
