package main

import (
	"fmt"
	"math/big"

	"github.com/svennjegac/adventofcode/2019/day10"
	"github.com/svennjegac/adventofcode/2019/day10/point"
)

func main() {
	asteroids, err := day10.Asteroids()
	if err != nil {
		panic(err)
	}

	max := 0
	var best point.Point
	for asteroid := range asteroids {
		canSee := make(map[string]struct{})
		for otherAsteroid := range asteroids {
			if asteroid == otherAsteroid {
				continue
			}
			dx := otherAsteroid.X - asteroid.X
			dy := otherAsteroid.Y - asteroid.Y

			if dx < 0 && dy < 0 {
				canSee["neg"+rStr(otherAsteroid, asteroid)] = struct{}{}
			} else if dx > 0 && dy > 0 {
				canSee[rStr(otherAsteroid, asteroid)] = struct{}{}
			} else if dx == 0 {
				if dy > 0 {
					canSee["Inf"] = struct{}{}
				} else {
					canSee["negInf"] = struct{}{}
				}
			} else if dy == 0 {
				if dx > 0 {
					canSee["Asi"] = struct{}{}
				} else {
					canSee["negAsi"] = struct{}{}
				}
			} else if dx > 0 {
				canSee[rStr(otherAsteroid, asteroid)] = struct{}{}
			} else {
				canSee["neg"+rStr(otherAsteroid, asteroid)] = struct{}{}
			}
		}
		if len(canSee) > max {
			max = len(canSee)
			best = asteroid
		}
	}

	fmt.Println("Max:", max)
	fmt.Println("Best:", best)
}

func rStr(other, aster point.Point) string {
	return big.NewRat(int64(other.Y-aster.Y), int64(other.X-aster.X)).String()
}
