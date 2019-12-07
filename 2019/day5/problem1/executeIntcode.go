package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/svennjegac/adventofcode/2019/day5"
)

func main() {
	intcode, err := day5.Intcode()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting intcode:", intcode)
	executeIntcode(intcode)
	fmt.Println("Executed intcode:", intcode)
}

func executeIntcode(intcode []int) {
	exec := newExecutor()
	for i := 0; i < len(intcode); {
		jmp, done := exec.Execute(i, intcode)
		if done {
			fmt.Println("Done,", jmp)
			return
		}
		i += jmp
	}
}

func newExecutor() *executor {
	return &executor{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

type executor struct {
	scanner *bufio.Scanner
}

func (e *executor) Execute(index int, intcode []int) (int, bool) {
	instr, modes := e.decodeOpcode(intcode[index])

	switch instr {
	case add:
		intcode[intcode[index+3]] = e.getVal(index+1, intcode, modes[0]) + e.getVal(index+2, intcode, modes[1])
		return 4, false
	case mul:
		intcode[intcode[index+3]] = e.getVal(index+1, intcode, modes[0]) * e.getVal(index+2, intcode, modes[1])
		return 4, false
	case in:
		e.scanner.Scan()
		userInput := e.scanner.Text()
		num, err := strconv.Atoi(userInput)
		if err != nil {
			panic(err)
		}
		intcode[intcode[index+1]] = num
		return 2, false
	case out:
		fmt.Println(e.getVal(index+1, intcode, modes[0]))
		return 2, false
	case halt:
		return -1, true
	default:
		return -2, true
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
	halt = 99
)

type mode int

const (
	position mode = iota
	immediate
)
