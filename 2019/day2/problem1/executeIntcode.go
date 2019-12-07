package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day2"
)

func main() {
	intcode, err := day2.Intcode()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting intcode:", intcode)

	intcode[1] = 12
	intcode[2] = 2
	executeIntcode(intcode)

	fmt.Println("Executed intcode:", intcode)
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
