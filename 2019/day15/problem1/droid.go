package main

import (
	"fmt"
	"github.com/svennjegac/adventofcode/2019/day15"
	"github.com/svennjegac/adventofcode/2019/day15/computer"
	"github.com/svennjegac/adventofcode/2019/day15/point"
	"math/big"
)

const (
	north = 1
	south = 2
	west  = 3
	east  = 4

	wall   = 0
	moved  = 1
	finish = 2
)

var oppositeDirection = map[int]int{
	north: south,
	south: north,
	west:  east,
	east:  west,
}

func main() {
	memory, err := day15.Intcode("2019/day15/intcode.txt")
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

	d := &droid{}
	area := make(map[point.Point]int)
	area[d.location] = moved

	dfs(d, area, in, out)

	printArea(area)

	fmt.Println("Bfs:", bfs(point.Point{}, area))
}

func dfs(d *droid, area map[point.Point]int, in, out chan *big.Int) {
	for i := 1; i <= 4; i++ {
		if occupied(d, area, i) {
			continue
		}

		in <- new(big.Int).SetInt64(int64(i))
		response := int((<-out).Int64())

		markLocation(d, area, i, response)

		if response == wall {
			continue
		}

		d.move(i)
		dfs(d, area, in, out)
		d.move(oppositeDirection[i])

		// reset droid
		in <- new(big.Int).SetInt64(int64(oppositeDirection[i]))
		<-out
	}
}

func bfs(p point.Point, area map[point.Point]int) int {
	q := &queue{}
	q.add(locPath{length: 0, p: p})

	visited := make(map[point.Point]struct{})

	for {
		next := q.remove()
		visited[next.p] = struct{}{}

		mark, ok := area[next.p]
		if !ok || mark == wall {
			continue
		}

		if mark == finish {
			return next.length
		}

		for _, n := range next.p.Neighs() {
			if _, ok := visited[n]; ok {
				continue
			}
			q.add(locPath{p: n, length: next.length + 1})
		}
	}
}

func occupied(d *droid, area map[point.Point]int, direction int) bool {
	d.move(direction)
	_, ok := area[d.location]
	d.move(oppositeDirection[direction])
	return ok
}

func markLocation(d *droid, area map[point.Point]int, direction, marker int) {
	d.move(direction)
	area[d.location] = marker
	d.move(oppositeDirection[direction])
}

type droid struct {
	location point.Point
}

func (d *droid) move(m int) {
	switch m {
	case north:
		d.location.Y++
	case south:
		d.location.Y--
	case east:
		d.location.X++
	case west:
		d.location.X--
	default:
		panic("Unknown direction:" + fmt.Sprint(m))
	}
}

type queue struct {
	vals []locPath
}

type locPath struct {
	p      point.Point
	length int
}

func (q *queue) remove() locPath {
	v := q.vals[0]
	q.vals = q.vals[1:]
	return v
}

func (q *queue) add(l locPath) {
	q.vals = append(q.vals, l)
}

func printArea(area map[point.Point]int) {
	minX, minY, maxX, maxY := 0, 0, 0, 0
	for p := range area {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	for j := maxY; j >= minY; j-- {
		for i := minX; i <= maxX; i++ {
			val, ok := area[point.Point{X: i, Y: j}]
			if !ok {
				fmt.Print(" ")
				continue
			}
			switch val {
			case wall:
				fmt.Print("#")
			case moved:
				fmt.Print(".")
			case finish:
				fmt.Print("F")
			}
		}
		fmt.Println(j)
	}
}
