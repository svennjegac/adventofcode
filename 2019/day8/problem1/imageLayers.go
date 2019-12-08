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

	minZeros := math.MaxInt64
	numOnes := 0
	numTwos := 0
	for i := 0; i < len(image); i += width * height {
		layerZeros := 0
		layerOnes := 0
		layerTwos := 0
		for j := i; j < i+width*height; j++ {
			if image[j] == 0 {
				layerZeros++
			} else if image[j] == 1 {
				layerOnes++
			} else if image[j] == 2 {
				layerTwos++
			}
		}
		if layerZeros < minZeros {
			minZeros = layerZeros
			numOnes = layerOnes
			numTwos = layerTwos
		}
	}

	fmt.Println("Checksum:", numOnes*numTwos)
}
