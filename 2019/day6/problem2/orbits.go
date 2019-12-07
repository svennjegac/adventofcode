package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day6"
	"github.com/svennjegac/adventofcode/2019/day6/graph"
)

func main() {
	orbits, youIndex, sanIndex, err := day6.Orbits()
	if err != nil {
		panic(err)
	}

	trsfrs := transfers{g: orbits, visited: make(map[int]int), you: youIndex, san: sanIndex}
	fmt.Println("Transfers:", trsfrs.get()-2)
}

type transfers struct {
	g       graph.Graph
	visited map[int]int
	you     int
	san     int
}

func (d transfers) get() int {
	d.dfs(d.you, 0)
	return d.dfs(d.san, 0)
}

func (d transfers) dfs(v int, dist int) int {
	d.visited[v] = dist
	for w := range d.g.Adj(v) {
		if otherDist, ok := d.visited[w]; !ok {
			return d.dfs(w, dist+1)
		} else {
			return otherDist + dist + 1
		}
	}
	return -1
}
