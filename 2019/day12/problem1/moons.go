package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day12"
)

func main() {
	moons := day12.Moons()

	for i := 0; i < 1000; i++ {
		for _, moon := range moons {
			for _, otherMoon := range moons {
				if moon.Id == otherMoon.Id {
					continue
				}

				dx := otherMoon.Position.X - moon.Position.X
				dy := otherMoon.Position.Y - moon.Position.Y
				dz := otherMoon.Position.Z - moon.Position.Z

				moon.Velocity.X += sign(dx)
				moon.Velocity.Y += sign(dy)
				moon.Velocity.Z += sign(dz)
			}
		}
		for _, moon := range moons {
			moon.Move()
		}
	}

	totalEnergy := 0
	for _, moon := range moons {
		totalEnergy += moon.TotalEnergy()
	}

	fmt.Println("Total energy:", totalEnergy)
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}
