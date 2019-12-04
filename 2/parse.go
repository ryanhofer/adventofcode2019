package main

import (
	"strconv"
	"strings"
)

func Parse(in string) []int {
	opstrs := strings.Split(in, ",")
	program := make([]int, 0, len(opstrs))

	for _, opstr := range opstrs {
		opcode, err := strconv.Atoi(opstr)
		check(err)

		program = append(program, opcode)
	}

	return program
}
