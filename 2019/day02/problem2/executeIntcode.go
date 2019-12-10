package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day02"
)

func main() {
	intcode, err := day02.Intcode()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting intcode:", intcode)

	noun, verb := findNounAndVerb(intcode)
	fmt.Println("n:", noun, "v:", verb)
	fmt.Println("100 * n + v = ", 100*noun+verb)
}

func findNounAndVerb(intcode []int) (int, int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			intcodeCopy := make([]int, len(intcode))
			copy(intcodeCopy, intcode)

			intcodeCopy[1] = noun
			intcodeCopy[2] = verb

			executeIntcode(intcodeCopy)

			if intcodeCopy[0] == 19690720 {
				return noun, verb
			}
		}
	}
	return -1, -1
}

func executeIntcode(intcode []int) {
	for i := 0; i < len(intcode); i += 4 {
		opcode := intcode[i]
		if opcode == 99 {
			return
		}
		operand1 := intcode[intcode[i+1]]
		operand2 := intcode[intcode[i+2]]
		resultAddr := intcode[i+3]
		intcode[resultAddr] = executeOperation(opcode, operand1, operand2)
	}
}

func executeOperation(opcode int, operand1 int, operand2 int) int {
	if opcode == 1 {
		return operand1 + operand2
	}
	return operand1 * operand2
}
