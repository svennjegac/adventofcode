package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day01"
)

func main() {
	moduleMasses, err := day01.ModuleMasses()
	if err != nil {
		panic(err)
	}

	totalFuel := calculateTotalFuel(moduleMasses)
	fmt.Println("Total fuel:", totalFuel)
}

func calculateTotalFuel(moduleMasses []int) int {
	totalFuel := 0
	for _, moduleMass := range moduleMasses {
		totalFuel += moduleMass/3 - 2
	}
	return totalFuel
}
