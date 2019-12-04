package main

import "fmt"

const (
	Add  = 1
	Mult = 2
	Halt = 99
)

func Exec(program []int, noun, verb int) int {
	mem := make([]int, 0, len(program))
	mem = append(mem, program...)
	mem[1] = noun
	mem[2] = verb
	pos := 0

loop:
	for {
		opcode := mem[pos]
		switch opcode {
		case Add:
			a := mem[mem[pos+1]]
			b := mem[mem[pos+2]]
			mem[mem[pos+3]] = a + b
			pos += 4
		case Mult:
			a := mem[mem[pos+1]]
			b := mem[mem[pos+2]]
			mem[mem[pos+3]] = a * b
			pos += 4
		case Halt:
			break loop
		default:
			panic(fmt.Sprintf("unexpected opcode %d at position %d", opcode, pos))
		}
	}

	return mem[0]
}
