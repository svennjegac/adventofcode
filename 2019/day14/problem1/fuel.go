package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day14"
	"github.com/svennjegac/adventofcode/2019/day14/chemical"
)

const (
	fuel = "FUEL"
	ore  = "ORE"
)

var usedOres = 0

func main() {
	reactionSystem := day14.Reactions("2019/day14/reactions.txt")
	chemicalStash := make(map[string]int)

	resolve(chemical.Chemical{Units: 1, Id: fuel}, reactionSystem, chemicalStash)

	fmt.Println("OREs used:", usedOres)
}

func resolve(resolvingChem chemical.Chemical, reactionSystem map[string]chemical.Reaction, chemicalStash map[string]int) {
	if resolvingChem.Id == ore {
		usedOres += resolvingChem.Units
		return
	}

	reaction := reactionSystem[resolvingChem.Id]
	for stashed := chemicalStash[resolvingChem.Id]; stashed < resolvingChem.Units; stashed = chemicalStash[resolvingChem.Id] {
		chemicalStash[resolvingChem.Id] += reaction.Product.Units
		for _, reactor := range reaction.Reactors {
			resolve(reactor, reactionSystem, chemicalStash)
		}
	}

	chemicalStash[resolvingChem.Id] -= resolvingChem.Units
}
