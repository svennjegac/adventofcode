package main

import (
	"fmt"
	"github.com/svennjegac/adventofcode/2019/day16"
)

const (
	phases          = 100
	inputMultiplier = 10000
)

func main() {
	input := day16.Input("2019/day16/input.txt")
	basePattern := []int{0, 1, 0, -1}

	mulInput := make([]int, inputMultiplier*len(input))
	for i := 0; i < inputMultiplier; i++ {
		copy(mulInput[i*len(input):], input)
	}
	input = mulInput

	patterns := createPatterns(basePattern, len(input))
	for i := 0; i < phases; i++ {
		input = fft(input, patterns)
	}

	fmt.Println("Done:", input[5976809:5976817])
}

type sparsePatternDigit struct {
	startOffset int
	endOffset   int
	digit       int
}

func createPatterns(basePattern []int, length int) [][]sparsePatternDigit {
	patterns := make([][]sparsePatternDigit, 0, length)

	for i := 1; i <= length; i++ {
		pattern := make([]sparsePatternDigit, 0)
		patternLen := 0

		for patternLen < length+1 {
			for _, bpd := range basePattern {
				patternLen += i

				if bpd == 0 {
					continue
				}

				pattern = append(pattern, sparsePatternDigit{
					startOffset: patternLen - i - 1,
					endOffset:   patternLen - 1,
					digit:       bpd,
				})
			}
		}

		patterns = append(patterns, pattern)
	}

	return patterns
}

func fft(input []int, patterns [][]sparsePatternDigit) []int {
	sumsArr := createSumsArr(input)

	output := make([]int, 0, len(input))
	for i := 0; i < len(input); i++ {
		digit := calculateDigit(input, patterns[i], sumsArr)
		output = append(output, digit)
	}
	return output
}

func createSumsArr(input []int) []int {
	res := make([]int, 0, len(input))
	res = append(res, input[0])
	for i := 1; i < len(input); i++ {
		res = append(res, input[i]+res[i-1])
	}
	return res
}

func calculateDigit(input []int, pattern []sparsePatternDigit, sumsArr []int) int {
	res := 0
	for _, spd := range pattern {
		l := spd.startOffset - 1
		if l >= len(input)-1 {
			break
		}
		r := spd.endOffset - 1
		if r >= len(input) {
			r = len(input) - 1
		}

		lSum := 0
		if l >= 0 {
			lSum = sumsArr[l]
		}

		rSum := sumsArr[r]

		res += spd.digit * (rSum - lSum)
	}
	return abs(res) % 10
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
