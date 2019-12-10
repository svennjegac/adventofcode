package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day06"
	"github.com/svennjegac/adventofcode/2019/day06/graph"
)

func main() {
	orbits, _, _, err := day06.Orbits()
	if err != nil {
		panic(err)
	}

	suroundingOrbits := suroundingOrbits{g: orbits, visited: make(map[int]int)}
	fmt.Println("Surounding orbits:", suroundingOrbits.get())
}

type suroundingOrbits struct {
	g       graph.Graph
	visited map[int]int
}

func (d suroundingOrbits) get() int {
	suborbits := 0
	for v := 0; v < d.g.V(); v++ {
		if _, ok := d.visited[v]; !ok {
			d.dfs(v)
		}
		suborbits += d.visited[v]
	}
	return suborbits
}

func (d suroundingOrbits) dfs(v int) {
	suborbits := 0
	for w := range d.g.Adj(v) {
		if _, ok := d.visited[w]; !ok {
			d.dfs(w)
		}
		suborbits = d.visited[w] + 1
	}
	d.visited[v] = suborbits
}
