package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/svennjegac/adventofcode/2019/day11"
	"github.com/svennjegac/adventofcode/2019/day11/computer"
)

func main() {
	memory, err := day11.Intcode("2019/day13/intcode.txt")
	if err != nil {
		panic(err)
	}
	// mem[0] = 2 for free play
	memory["0"] = new(big.Int).SetInt64(2)

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
		case <-time.After(time.Millisecond):
			in <- new(big.Int).SetInt64(printGameTerminal(computerOutput))
		}
	}

	printGameTerminal(computerOutput)
}

func printGameTerminal(computerOutput []int64) int64 {
	var maxX, maxY, score int64 = 0, 0, 0
	compOut2D := make(map[int64]map[int64]int64)
	var paddleX, ballX int64 = 0, 0

	for i := 0; i < len(computerOutput); i += 3 {
		x, y, val := computerOutput[i], computerOutput[i+1], computerOutput[i+2]
		if x == -1 && y == 0 {
			score = val
			continue
		}
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}

		ysForX, ok := compOut2D[x]
		if !ok {
			ysForX = make(map[int64]int64)
			compOut2D[x] = ysForX
		}
		ysForX[y] = val

		if val == 3 {
			paddleX = x
		}
		if val == 4 {
			ballX = x
		}
	}

	fmt.Println("Your current score is:", score)
	fmt.Println("Game terminal state:")
	for j := int64(0); j < maxY; j++ {
		for i := int64(0); i < maxX; i++ {
			switch compOut2D[i][j] {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print(".")
			case 2:
				fmt.Print("X")
			case 3:
				fmt.Print("P")
			case 4:
				fmt.Print("O")
			default:
				panic("Unknown object")
			}

		}
		fmt.Println()
	}

	if paddleX < ballX {
		return 1
	} else if paddleX > ballX {
		return -1
	} else {
		return 0
	}
}
