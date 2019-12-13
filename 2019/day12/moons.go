package day12

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/svennjegac/adventofcode/2019/day12/moon"
)

func Moons() []*moon.Moon {
	file, err := os.Open("2019/day12/moons.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")
	moons := make([]*moon.Moon, 0, len(lines))
	for i, line := range lines {
		splits := strings.Split(line[1:len(line)-1], ",")
		x, _ := strconv.Atoi(splits[0][strings.Index(splits[0], "=")+1:])
		y, _ := strconv.Atoi(splits[1][strings.Index(splits[1], "=")+1:])
		z, _ := strconv.Atoi(splits[2][strings.Index(splits[2], "=")+1:])
		moon := &moon.Moon{
			Id: i,
			Position: &moon.Point{
				X: x,
				Y: y,
				Z: z,
			},
			Velocity: &moon.Point{},
		}
		moons = append(moons, moon)
	}
	return moons
}
