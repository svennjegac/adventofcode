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
	ox = 3
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

	len, fin := bfs(point.Point{}, area)
	fmt.Println("Bfs:", len, "->", fin)
	fmt.Println("Oxygen:", oxygen(point.Point{X:16, Y:-18}, area))
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

func bfs(p point.Point, area map[point.Point]int) (int, point.Point) {
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
			return next.length, next.p
		}

		for _, n := range next.p.Neighs() {
			if _, ok := visited[n]; ok {
				continue
			}
			q.add(locPath{p: n, length: next.length + 1})
		}
	}
}

func oxygen(p point.Point, area map[point.Point]int) int {
	minutes := 0

	q := &queue{}
	q.add(locPath{p: p, length: -1})

	qFuture := &queue{}

	visited := make(map[point.Point]struct{})
	for {
		for q.size() > 0 {
			next := q.remove()
			visited[next.p] = struct{}{}

			mark, ok := area[next.p]
			if !ok || mark == wall {
				continue
			}

			area[next.p] = ox

			for _, n := range next.p.Neighs() {
				if _, ok := visited[n]; ok {
					continue
				}
				qFuture.add(locPath{p: n, length: -1})
			}
		}
		if qFuture.size() == 0 {
			break
		}
		q.vals = qFuture.vals
		qFuture.vals = nil

		minutes++
		fmt.Println("New step, minutes:", minutes)
		printArea(area)
	}

	return minutes
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

func (q *queue) size() int {
	return len(q.vals)
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
			case ox:
				fmt.Print("O")
			}
		}
		fmt.Println(j)
	}
}
