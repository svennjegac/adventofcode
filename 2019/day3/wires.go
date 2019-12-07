package day3

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/svennjegac/adventofcode/2019/day3/point"
)

func Wires() (map[point.Point]int, map[point.Point]int, error) {
	file, err := os.Open("2019/day3/wires.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	wireSpecifications := strings.Split(string(fileBytes), "\n")

	wire1, err := createWire(wireSpecifications[0])
	if err != nil {
		return nil, nil, err
	}

	wire2, err := createWire(wireSpecifications[1])
	if err != nil {
		return nil, nil, err
	}

	return wire1, wire2, nil
}

func createWire(wireSpecification string) (map[point.Point]int, error) {
	movingInstructions := strings.Split(wireSpecification, ",")

	wire := make(map[point.Point]int)
	current := point.New(0, 0)
	length := 0
	for _, movingInstruction := range movingInstructions {
		direction := movingInstruction[0]
		steps, err := strconv.Atoi(movingInstruction[1:])
		if err != nil {
			return nil, err
		}

		for i := 0; i < steps; i++ {
			length += 1
			current = current.Add(move[string(direction)])
			if _, ok := wire[current]; !ok {
				wire[current] = length
			}
		}
	}

	return wire, nil
}

var move = map[string]point.Point{
	"R": point.New(1, 0),
	"L": point.New(-1, 0),
	"U": point.New(0, 1),
	"D": point.New(0, -1),
}
