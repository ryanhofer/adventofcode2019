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
	RelOffsetOpcode   = 9
	HaltOpcode        = 99
)

const (
	PositionMode  = 0
	ImmediateMode = 1
	RelativeMode  = 2
)

type Config struct {
	Noun   nulls.Int
	Verb   nulls.Int
	Input  <-chan Word
	Output chan<- Word
}

type process struct {
	mem map[Word]Word
	pos Word
	rel Word
}

func (proc *process) read(addr Word) Word {
	return proc.mem[addr]
}

func (proc *process) write(addr Word, value Word) {
	proc.mem[addr] = value
}

func (proc *process) next() Word {
	val := proc.read(proc.pos)
	proc.pos++
	return val
}

func Exec(program Program, cfg *Config) (exit Word, err error) {
	if cfg.Output == nil {
		cfg.Output = make(chan Word)
	}
	defer close(cfg.Output)

	proc := process{}
	proc.mem = make(map[Word]Word, len(program))
	for addr, value := range program {
		proc.mem[Word(addr)] = value
	}

	if cfg.Noun.Valid && len(proc.mem) > 1 {
		proc.write(1, Word(cfg.Noun.Int))
	}
	if cfg.Verb.Valid && len(proc.mem) > 2 {
		proc.write(2, Word(cfg.Verb.Int))
	}

	for {
		instr := proc.next()
		opcode := instr % 100
		modes := instr / 100

		param := func(dest bool) Word {
			v := proc.next()
			m := modes % 10
			modes /= 10
			switch m {
			case ImmediateMode:
				return v
			case RelativeMode:
				v += proc.rel
			}
			if !dest {
				v = proc.read(v)
			}
			return v
		}

		operand := func() Word {
			return param(false)
		}

		destination := func() Word {
			return param(true)
		}

		switch opcode {
		case AddOpcode:
			a := operand()
			b := operand()
			dest := destination()
			proc.write(dest, a+b)

		case MultOpcode:
			a := operand()
			b := operand()
			dest := destination()
			proc.write(dest, a*b)

		case InputOpcode:
			dest := destination()
			input := <-cfg.Input
			proc.write(dest, input)

		case OutputOpcode:
			output := operand()
			cfg.Output <- output

		case JumpIfTrueOpcode:
			val := operand()
			pos := operand()
			if val != 0 {
				proc.pos = pos
			}

		case JumpIfFalseOpcode:
			val := operand()
			pos := operand()
			if val == 0 {
				proc.pos = pos
			}

		case LessThanOpcode:
			a := operand()
			b := operand()
			dest := destination()
			val := Word(0)
			if a < b {
				val = 1
			}
			proc.write(dest, val)

		case EqualsOpcode:
			a := operand()
			b := operand()
			dest := destination()
			val := Word(0)
			if a == b {
				val = 1
			}
			proc.write(dest, val)

		case RelOffsetOpcode:
			val := operand()
			proc.rel += val

		case HaltOpcode:
			exit = proc.read(0)
			return

		default:
			err = fmt.Errorf("unexpected opcode %d at position %d", opcode, proc.pos)
			return
		}
	}
}
