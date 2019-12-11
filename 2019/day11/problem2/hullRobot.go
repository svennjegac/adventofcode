package main

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/svennjegac/adventofcode/2019/day11"
	"github.com/svennjegac/adventofcode/2019/day11/computer"
)

func main() {
	memory, err := day11.Intcode("2019/day11/intcode.txt")
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

	hullRobot := &robot{
		grid: map[point]int{
			point{0, 0}: 1,
		},
	}

	for {
		in <- big.NewInt(int64(hullRobot.scan()))
		select {
		case <-done:
			fmt.Println("Painted fields:", len(hullRobot.grid))
			paint(hullRobot.grid)
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
	position  point
	grid      map[point]int
}

func (r *robot) turn(t int) {
	if t == 0 {
		t = -1
	}
	r.direction = (r.direction + t + 4) % 4
	move := directionToMove[r.direction]
	r.position = point{r.position.x + move.x, r.position.y + move.y}
}

func (r *robot) scan() int {
	return r.grid[r.position]
}

func (r *robot) paint(p int) {
	r.grid[r.position] = p
}

var directionToMove = map[int]point{
	0: {0, 1},
	1: {1, 0},
	2: {0, -1},
	3: {-1, 0},
}

type point struct {
	x, y int
}

func paint(grid map[point]int) {
	panels := make([]panel, 0)
	for p, clr := range grid {
		panels = append(panels, panel{p, clr})
	}

	sort.Slice(panels, func(i, j int) bool {
		if panels[i].x < panels[j].x {
			return true
		} else if panels[i].x == panels[j].x && panels[i].y < panels[j].y {
			return true
		}
		return false
	})

	row := panels[0].x
	for _, pnl := range panels {
		if pnl.x != row {
			fmt.Println()
			row = pnl.x
		}
		if pnl.color == 0 {
			fmt.Print(".")
		} else {
			fmt.Print("#")
		}
	}
}

type panel struct {
	point
	color int
}
