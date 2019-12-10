package main

import (
	"fmt"
	"math"
	"sync"

	"github.com/svennjegac/adventofcode/2019/day07"
)

const (
	amplifiers = 5
	phases     = 99999
)

func main() {
	max := math.MinInt64
	for phase := 0; phase < phases; phase++ {
		if isValid(phase) {
			out := tryPhase(phase)
			if out > max {
				fmt.Println("Valid phase:", phase)
				fmt.Println("Phase power:", out)
				max = out
			}
		}
	}
	fmt.Println("Max thruster power:", max)
}

func isValid(phase int) bool {
	phaseSettings := make(map[int]struct{})
	for i := 0; i < amplifiers; i++ {
		phaseSettings[phase%10] = struct{}{}
		phase /= 10
	}
	for p := range phaseSettings {
		if p < amplifiers {
			return false
		}
	}
	return len(phaseSettings) == amplifiers
}

func tryPhase(phase int) int {
	executors := make([]*executor, 0)
	for i := 0; i < amplifiers; i++ {
		intcode, err := day07.Intcode()
		if err != nil {
			panic(err)
		}

		out := make(chan int, 2)
		exec := &executor{
			intcode: intcode,
			out:     out,
		}
		executors = append(executors, exec)
	}

	phaseCodes := make([]int, amplifiers)
	phaseIndex := 0
	for phase > 0 {
		phaseCodes[amplifiers-phaseIndex-1] = phase % 10
		phase /= 10
		phaseIndex++
	}

	for i := 0; i < amplifiers; i++ {
		if i == 0 {
			executors[i].in = executors[len(executors)-1].out
		} else {
			executors[i].in = executors[i-1].out
		}
		executors[i].in <- phaseCodes[i]
	}

	executors[0].in <- 0

	wg := sync.WaitGroup{}
	wg.Add(amplifiers)
	for i := 0; i < amplifiers; i++ {
		go func(i int) {
			executors[i].executeIntcode()
			wg.Done()
		}(i)
	}

	wg.Wait()

	return <-executors[amplifiers-1].out
}

type executor struct {
	intcode []int
	in      chan int
	out     chan int
}

func (e *executor) executeIntcode() {
	for i := 0; i < len(e.intcode); {
		jmp, isJump, done := e.execute(i, e.intcode)
		if done {
			return
		} else if isJump {
			i = jmp
		} else {
			i += jmp
		}
	}
}

func (e *executor) execute(index int, intcode []int) (int, bool, bool) {
	instr, modes := e.decodeOpcode(intcode[index])

	switch instr {
	case add:
		intcode[intcode[index+3]] = e.getVal(index+1, intcode, modes[0]) + e.getVal(index+2, intcode, modes[1])
		return 4, false, false
	case mul:
		intcode[intcode[index+3]] = e.getVal(index+1, intcode, modes[0]) * e.getVal(index+2, intcode, modes[1])
		return 4, false, false
	case in:
		intcode[intcode[index+1]] = <-e.in
		return 2, false, false
	case out:
		e.out <- e.getVal(index+1, intcode, modes[0])
		return 2, false, false
	case jmpIfTrue:
		first := e.getVal(index+1, intcode, modes[0])
		second := e.getVal(index+2, intcode, modes[1])
		if first != 0 {
			return second, true, false
		}
		return 3, false, false
	case jmpIfFalse:
		first := e.getVal(index+1, intcode, modes[0])
		second := e.getVal(index+2, intcode, modes[1])
		if first == 0 {
			return second, true, false
		}
		return 3, false, false
	case lessThan:
		first := e.getVal(index+1, intcode, modes[0])
		second := e.getVal(index+2, intcode, modes[1])
		if first < second {
			intcode[intcode[index+3]] = 1
		} else {
			intcode[intcode[index+3]] = 0
		}
		return 4, false, false
	case equals:
		first := e.getVal(index+1, intcode, modes[0])
		second := e.getVal(index+2, intcode, modes[1])
		if first == second {
			intcode[intcode[index+3]] = 1
		} else {
			intcode[intcode[index+3]] = 0
		}
		return 4, false, false
	case halt:
		return -1, false, true
	default:
		return -2, false, true
	}
}

func (e *executor) decodeOpcode(opcode int) (instruction, []mode) {
	instr := instruction(opcode % 100)

	opcode /= 100

	modes := make([]mode, 0)
	for i := 0; i < 3; i++ {
		modes = append(modes, mode(opcode%10))
		opcode /= 10
	}

	return instr, modes
}

func (e *executor) getVal(index int, intcode []int, m mode) int {
	if m == position {
		return intcode[intcode[index]]
	}
	return intcode[index]
}

type instruction int

const (
	unknown instruction = iota
	add
	mul
	in
	out
	jmpIfTrue
	jmpIfFalse
	lessThan
	equals
	halt = 99
)

type mode int

const (
	position mode = iota
	immediate
)
