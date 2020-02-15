package main

import (
	"fmt"
	"github.com/svennjegac/adventofcode/2019/day16"
)

const (
	phases = 4
)

func main() {
	input := day16.Input("2019/day16/input.txt")
	basePattern := []int{0, 1, 0, -1}

	for i := 0; i < phases; i++ {
		input = fft(input, basePattern)
		fmt.Println(input)
	}
}

func fft(input []int, basePattern []int) []int {
	output := make([]int, 0, len(input))
	for i := 0; i < len(input); i++ {
		pattern := createPattern(basePattern, i, len(input))
		digit := calculateDigit(input, pattern)
		output = append(output, digit)
	}
	return output
}

func createPattern(basePattern []int, index, length int) []int {
	pattern := make([]int, 0, length+1)
	patternLen := 0
	for patternLen < length+1 {
		for _, bpd := range basePattern {
			for i := 0; i <= index; i++ {
				pattern = append(pattern, bpd)
				patternLen++
			}
		}
	}
	return pattern[1:]
}

func calculateDigit(input []int, pattern []int) int {
	res := 0
	for i := 0; i < len(input); i++ {
		res += input[i] * pattern[i]
	}
	return abs(res) % 10
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
