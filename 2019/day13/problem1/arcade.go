package main

import (
	"fmt"
	"math/big"

	"github.com/svennjegac/adventofcode/2019/day13"
	"github.com/svennjegac/adventofcode/2019/day13/computer"
)

func main() {
	memory, err := day13.Intcode("2019/day13/intcode.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)

	done := make(chan struct{})
	go func() {
		comp.Run()
		close(done)
	}()

	computerOutput := make([]int64, 0)
loop:
	for {
		select {
		case <-done:
			fmt.Println("Intcode done")
			break loop
		case outVal := <-out:
			computerOutput = append(computerOutput, outVal.Int64())
		}
	}

	blockTiles := 0
	for i := 2; i < len(computerOutput); i += 3 {
		if computerOutput[i] == 2 {
			blockTiles++
		}
	}
	fmt.Println("Block tiles:", blockTiles)
}
