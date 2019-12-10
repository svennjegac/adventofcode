package main

import (
	"fmt"
	"math/big"

	"github.com/svennjegac/adventofcode/2019/day09"
	"github.com/svennjegac/adventofcode/2019/day09/computer"
)

func main() {
	memory, err := day09.Intcode("2019/day09/intcode.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int, 100)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)

	in <- big.NewInt(1)
	comp.Run()

	fmt.Println(<-out)
}
