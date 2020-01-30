package day14

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/svennjegac/adventofcode/2019/day14/chemical"
)

func Reactions(f string) map[string]chemical.Reaction {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileBytes), "\n")

	reactionSystem := make(map[string]chemical.Reaction)
	for _, line := range lines {
		splittedEquation := strings.Split(line, "=>")

		splittedEq1 := strings.Split(splittedEquation[0], ",")
		reactors := make([]chemical.Chemical, 0, len(splittedEq1))
		for _, s := range splittedEq1 {
			reactors = append(reactors, createChemical(s))
		}

		chemRes := createChemical(splittedEquation[1])
		if _, ok := reactionSystem[chemRes.Id]; ok {
			panic("already existing chem")
		}
		reactionSystem[chemRes.Id] = chemical.Reaction{
			Reactors: reactors,
			Product:  chemRes,
		}
	}

	return reactionSystem
}

func createChemical(s string) chemical.Chemical {
	s = strings.TrimSpace(s)
	sSplit := strings.Split(s, " ")

	units, err := strconv.Atoi(sSplit[0])
	if err != nil {
		panic(err)
	}

	return chemical.Chemical{
		Units: units,
		Id:    sSplit[1],
	}
}
