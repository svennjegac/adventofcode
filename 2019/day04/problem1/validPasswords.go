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
	removedDigit := i % 10
	i /= 10
	for i > 0 {
		lastDigit := i % 10
		if lastDigit == removedDigit {
			return true
		}
		removedDigit = lastDigit
		i /= 10
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
