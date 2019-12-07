package main

import (
	"fmt"
)

func main() {
	validPasswords := calculateNumOfValidPasswords(137683, 596253)
	fmt.Println("Valid passwords:", validPasswords)
}

func calculateNumOfValidPasswords(l int, r int) int {
	validPasswords := 0
	for i := l; i <= r; i++ {
		if isValid(i) {
			validPasswords++
		}
	}
	return validPasswords
}

func isValid(i int) bool {
	return hasAdjDuplicate(i) && dontHaveDecreasingDigits(i)
}

func hasAdjDuplicate(i int) bool {
	digitsToIndexes := make(map[int][]int)

	index := 0
	for i > 0 {
		lastDigit := i % 10
		digitsToIndexes[lastDigit] = append(digitsToIndexes[lastDigit], index)
		i /= 10
		index++
	}

	for _, indexes := range digitsToIndexes {
		if len(indexes) == 2 && indexes[1]-indexes[0] == 1 {
			return true
		}
	}

	return false
}

func dontHaveDecreasingDigits(i int) bool {
	removedDigit := i % 10
	i /= 10
	for i > 0 {
		lastDigit := i % 10
		if lastDigit > removedDigit {
			return false
		}
		removedDigit = lastDigit
		i /= 10
	}
	return true
}
