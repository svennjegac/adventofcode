package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/svennjegac/adventofcode/2019/day12"
	"github.com/svennjegac/adventofcode/2019/day12/moon"
)

func main() {
	moons := day12.Moons()

	x := equilibrium(copyMoons(moons), (*moon.Point).XC)
	y := equilibrium(copyMoons(moons), (*moon.Point).YC)
	z := equilibrium(copyMoons(moons), (*moon.Point).ZC)
	fmt.Println("Time when universe repeats itself:", leastCommonMultiple(x, y, z)*2)
}

func copyMoons(moons []*moon.Moon) []*moon.Moon {
	moonsCpy := make([]*moon.Moon, 0, len(moons))
	for _, moon := range moons {
		moonsCpy = append(moonsCpy, moon.Copy())
	}
	return moonsCpy
}

func equilibrium(moons []*moon.Moon, coord func(p *moon.Point) int) int {
	previousStates := make(map[string]struct{})
	for i := 0; ; i++ {
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
		previousStates[hashMoons(moons, coord)] = struct{}{}

		for _, moon := range moons {
			moon.Move()
		}

		if _, ok := previousStates[hashMoons(moons, coord)]; ok {
			return i + 1
		}
	}
}

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func hashMoons(moons []*moon.Moon, coord func(p *moon.Point) int) string {
	sb := &strings.Builder{}
	for _, moon := range moons {
		sb.WriteString(strconv.Itoa(coord(moon.Position)) + strconv.Itoa(coord(moon.Velocity)))
	}
	return sb.String()
}

func leastCommonMultiple(nums ...int) int {
	allOnes := func() bool {
		for _, n := range nums {
			if n != 1 {
				return false
			}
		}
		return true
	}

	lcm := 1
	divisor := 2
	for {
		if allOnes() {
			break
		}
		divided := false
		for i, n := range nums {
			if n%divisor == 0 {
				nums[i] /= divisor
				divided = true
			}
		}
		if divided {
			lcm *= divisor
		} else {
			divisor++
		}
	}
	return lcm
}
