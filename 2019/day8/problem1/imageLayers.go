package main

import (
	"fmt"
	"math"

	"github.com/svennjegac/adventofcode/2019/day8"
)

const (
	width  = 25
	height = 6
)

func main() {
	image, err := day8.Image()
	if err != nil {
		panic(err)
	}

	counts := make([]int, 3)
	counts[0] = math.MaxInt64
	for i := 0; i < len(image); i += width * height {
		layerCounts := make([]int, 3)
		for j := i; j < i+width*height; j++ {
			layerCounts[image[j]] += 1
		}

		if layerCounts[0] < counts[0] {
			counts = layerCounts
		}
	}

	fmt.Println("Checksum:", counts[1]*counts[2])
}
