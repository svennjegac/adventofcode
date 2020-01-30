package main

import (
	"fmt"

	"github.com/svennjegac/adventofcode/2019/day14"
	"github.com/svennjegac/adventofcode/2019/day14/chemical"
)

const (
	fuel = "FUEL"
	ore  = "ORE"

	factor      = 500
	ingredients = 1000000000000
)

var oresUsed = 0

func main() {
	reactionSystem := day14.Reactions("2019/day14/reactions.txt")
	chemicalStash := make(map[string]int)

	f := 1
	f *= factor
	modify(reactionSystem, func(i int) int {
		return i * factor
	})

	fuelCounter := 0
	lowed := false
	for {
		resolve(chemical.Chemical{Units: f, Id: fuel}, reactionSystem, chemicalStash)
		fuelCounter += f

		if oresUsed > ingredients-9999999999 && !lowed {
			f /= factor
			modify(reactionSystem, func(i int) int {
				return i / factor
			})
			lowed = true
		}
		if oresUsed > ingredients {
			fuelCounter -= f
			break
		}
	}

	fmt.Println("Fuel:", fuelCounter)
}

func resolve(resolvingChem chemical.Chemical, reactionSystem map[string]chemical.Reaction, chemicalStash map[string]int) {
	if resolvingChem.Id == ore {
		oresUsed += resolvingChem.Units
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

func modify(reactionSystem map[string]chemical.Reaction, modifier func(int) int) {
	for id, reaction := range reactionSystem {
		reaction.Product.Units = modifier(reaction.Product.Units)
		for i, reactor := range reaction.Reactors {
			reactor.Units = modifier(reactor.Units)
			reaction.Reactors[i] = reactor
		}
		reactionSystem[id] = reaction
	}
}
