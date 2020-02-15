package main

import (
	"fmt"
	"github.com/svennjegac/adventofcode/2019/day17"
	"github.com/svennjegac/adventofcode/2019/day17/computer"
	"math/big"
	"time"
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

	finalOut := 0
loop:
	for {
		select {
		case <-done:
			break loop
		case co := <-out:
			finalOut = int(co.Int64())
		case <-time.After(time.Second):
			go loadRules(in)
		}
	}

	fmt.Println("Dust:", finalOut)
}

func loadRules(in chan *big.Int) {
	// main: A,B,A,C,B,A,C,B,A,C10
	mainRoutine := []int{65,44,66,44,65,44,67,44,66,44,65,44,67,44,66,44,65,44,67,10}
	// a: L,12,L,12,L,6,L,610
	a := []int{76,44,49,50,44,76,44,49,50,44,76,44,54,44,76,44,54,10}
	// b: R,8,R,4,L,12,10
	b := []int{82,44,56,44,82,44,52,44,76,44,49,50,10}
	// c: L,12,L,6,R,12,R,8
	c := []int{76,44,49,50,44,76,44,54,44,82,44,49,50,44,82,44,56,10}
	// video
	video := []int{110,10}

	load := func(args []int) {
		for _, a := range args {
			in <- new(big.Int).SetInt64(int64(a))
		}
	}

	load(mainRoutine)
	load(a)
	load(b)
	load(c)
	load(video)
}
