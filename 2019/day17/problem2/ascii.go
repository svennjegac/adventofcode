package main

import (
	"fmt"
	"github.com/svennjegac/adventofcode/2019/day17"
	"github.com/svennjegac/adventofcode/2019/day17/computer"
	"math/big"
)

func main() {
	memory, err := day17.Intcode("2019/day17/intcode.txt")
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

	go loadRules(in)

	finalOut := 0
loop:
	for {
		select {
		case <-done:
			break loop
		case co := <-out:
			finalOut = int(co.Int64())
		}
	}

	fmt.Println("Dust:", finalOut)
}

func loadRules(in chan *big.Int) {
	mainRoutine := "A,B,A,C,B,A,C,B,A,C\n"
	a := "L,12,L,12,L,6,L,6\n"
	b := "R,8,R,4,L,12\n"
	c := "L,12,L,6,R,12,R,8\n"
	video := "n\n"

	ls := func(s string) {
		for _, r := range s {
			in <- new(big.Int).SetInt64(int64(r))
		}
	}

	ls(mainRoutine)
	ls(a)
	ls(b)
	ls(c)
	ls(video)
}
