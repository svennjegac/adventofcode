package main

import (
	"fmt"
	"math"

	"github.com/svennjegac/adventofcode/2019/day3"
	"github.com/svennjegac/adventofcode/2019/day3/point"
)

func main() {
	wire1, wire2, err := day3.Wires()
	if err != nil {
		panic(err)
	}

	closestIntersection := calculateClosestIntersection(wire1, wire2)

	fmt.Println("Closest intersection:", closestIntersection)
}

func calculateClosestIntersection(wire1 map[point.Point]int, wire2 map[point.Point]int) interface{} {
	closest := math.MaxInt64
	for p1, length1 := range wire1 {
		if length2, ok := wire2[p1]; ok {
			distance := length1 + length2
			if distance < closest {
				closest = distance
			}
		}
	}
	return closest
}
