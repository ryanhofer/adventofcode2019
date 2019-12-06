package intcode

import (
	"strconv"
	"strings"
)

func Parse(in string) ([]int, error) {
	opstrs := strings.Split(in, ",")
	program := make([]int, 0, len(opstrs))

	for _, opstr := range opstrs {
		opcode, err := strconv.Atoi(opstr)
		if err != nil {
			return []int{}, err
		}

		program = append(program, opcode)
	}

	return program, nil
}
