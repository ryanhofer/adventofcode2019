package intcode

import (
	"strconv"
	"strings"
)

type Word int64
type Program []Word

func Parse(in string) (Program, error) {
	opstrs := strings.Split(in, ",")
	program := Program{}

	for _, opstr := range opstrs {
		opcode, err := strconv.ParseInt(opstr, 10, 64)
		if err != nil {
			return Program{}, err
		}

		program = append(program, Word(opcode))
	}

	return program, nil
}
