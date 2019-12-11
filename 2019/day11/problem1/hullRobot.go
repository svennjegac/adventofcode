package main

import (
	"fmt"
	"math/big"

	"github.com/svennjegac/adventofcode/2019/day11"
	"github.com/svennjegac/adventofcode/2019/day11/computer"
)

func main() {
	memory, err := day11.Intcode("2019/day11/intcode.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 10)
	out := make(chan *big.Int, 20)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)

	done := make(chan struct{})
	go func() {
		comp.Run()
		close(done)
	}()

	hullRobot := &robot{
		grid: make(map[point]int),
	}

	for {
		in <- big.NewInt(int64(hullRobot.scan()))
		select {
		case <-done:
			fmt.Println("Painted fields:", len(hullRobot.grid))
			return
		case paint := <-out:
			turn := <-out
			hullRobot.paint(int(paint.Int64()))
			hullRobot.turn(int(turn.Int64()))
		}
	}
}

type robot struct {
	direction int // 0=up, 1=right, 2=down, 3=left
	position point
	grid map[point]int
}

func (r *robot) turn(t int) {
	if t == 0 {
		r.direction = (r.direction - 1 + 4) % 4
	} else {
		r.direction = (r.direction + 1 + 4) % 4
	}
	move := directionToMove[r.direction]
	r.position = point{r.position.x+move.x, r.position.y+move.y}
}

func (r *robot) scan() int {
	return r.grid[r.position]
}

func (r *robot) paint(p int) {
	r.grid[r.position] = p
}

var directionToMove = map[int]point{
	0: {0,1},
	1: {1, 0},
	2: {0, -1},
	3: {-1, 0},
}

type point struct {
	x, y int
}
