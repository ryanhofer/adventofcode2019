package intcode

import (
	"fmt"

	"github.com/gobuffalo/nulls"
)

const (
	AddOpcode         = 1
	MultOpcode        = 2
	InputOpcode       = 3
	OutputOpcode      = 4
	JumpIfTrueOpcode  = 5
	JumpIfFalseOpcode = 6
	LessThanOpcode    = 7
	EqualsOpcode      = 8
	HaltOpcode        = 99
)

const (
	PositionMode  = 0
	ImmediateMode = 1
)

type Config struct {
	Noun  nulls.Int
	Verb  nulls.Int
	Input int
}

type process struct {
	mem []int
	pos int
}

func (proc *process) next() int {
	val := proc.mem[proc.pos]
	proc.pos++
	return val
}

func Exec(program []int, cfg *Config) (exit int, output []int, err error) {
	proc := process{}
	proc.mem = append(proc.mem, program...)

	if cfg.Noun.Valid && len(proc.mem) > 1 {
		proc.mem[1] = cfg.Noun.Int
	}
	if cfg.Verb.Valid && len(proc.mem) > 2 {
		proc.mem[2] = cfg.Verb.Int
	}

loop:
	for {
		instr := proc.next()
		opcode := instr % 100
		modes := instr / 100

		nextArg := func() int {
			v := proc.next()
			m := modes % 10
			modes /= 10
			if m == PositionMode {
				v = proc.mem[v]
			}
			return v
		}

		switch opcode {
		case AddOpcode:
			a := nextArg()
			b := nextArg()
			dest := proc.next()
			proc.mem[dest] = a + b
		case MultOpcode:
			a := nextArg()
			b := nextArg()
			dest := proc.next()
			proc.mem[dest] = a * b
		case InputOpcode:
			in := cfg.Input
			dest := proc.next()
			proc.mem[dest] = in
		case OutputOpcode:
			out := nextArg()
			output = append(output, out)
		case JumpIfTrueOpcode:
			val := nextArg()
			pos := nextArg()
			if val != 0 {
				proc.pos = pos
			}
		case JumpIfFalseOpcode:
			val := nextArg()
			pos := nextArg()
			if val == 0 {
				proc.pos = pos
			}
		case LessThanOpcode:
			a := nextArg()
			b := nextArg()
			dest := proc.next()
			val := 0
			if a < b {
				val = 1
			}
			proc.mem[dest] = val
		case EqualsOpcode:
			a := nextArg()
			b := nextArg()
			dest := proc.next()
			val := 0
			if a == b {
				val = 1
			}
			proc.mem[dest] = val
		case HaltOpcode:
			break loop
		default:
			err = fmt.Errorf("unexpected opcode %d at position %d", opcode, proc.pos)
			return
		}
	}

	exit = proc.mem[0]
	return
}
