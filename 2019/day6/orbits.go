package day6

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/svennjegac/adventofcode/2019/day6/graph"
)

func Orbits() (graph.Graph, int, int, error) {
	file, err := os.Open("2019/day6/orbits.txt")
	if err != nil {
		return graph.Graph{}, 0, 0, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return graph.Graph{}, 0, 0, err
	}

	orbits := strings.Split(string(fileBytes), "\n")
	planetsToIndexes := make(map[string]int)
	index := 0
	edges := make([][]int, 0)

	youIndex := 0
	sanIndex := 0

	for _, orbit := range orbits {
		planets := strings.Split(orbit, ")")
		p1 := planets[0]
		p2 := planets[1]

		p1Index, ok := planetsToIndexes[p1]
		if !ok {
			planetsToIndexes[p1] = index
			p1Index = index
			index++
		}

		p2Index, ok := planetsToIndexes[p2]
		if !ok {
			planetsToIndexes[p2] = index
			p2Index = index
			index++
		}

		if p2 == "YOU" {
			youIndex = p2Index
		} else if p2 == "SAN" {
			sanIndex = p2Index
		}

		edges = append(edges, []int{p2Index, p1Index})
	}

	g := graph.New(len(planetsToIndexes))
	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	return g, youIndex, sanIndex, nil
}
