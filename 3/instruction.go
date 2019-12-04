package main

import (
	"fmt"
	"strings"
)

type Instruction struct {
	Direction Direction
	Distance  int
}

func Parse(in string) []Instruction {
	tokens := strings.Split(in, ",")
	instructions := make([]Instruction, 0, len(tokens))

	for _, token := range tokens {
		var directionstr string
		var distance int
		_, err := fmt.Sscanf(token, "%1s%d", &directionstr, &distance)
		check(err)

		direction := str2dir(directionstr)
		instruction := Instruction{
			Direction: direction,
			Distance:  distance,
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
