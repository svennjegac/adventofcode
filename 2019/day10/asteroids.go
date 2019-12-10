package day10

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/svennjegac/adventofcode/2019/day10/point"
)

func Asteroids() (map[point.Point]struct{}, error) {
	file, err := os.Open("2019/day10/asteroids.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileBytes), "\n")
	asteroids := make(map[point.Point]struct{})
	for i, line := range lines {
		for j, object := range line {
			if string(object) == "#" {
				asteroids[point.New(j, i)] = struct{}{}
			}
		}
	}
	return asteroids, nil
}

func Asteroids2(x, y int) (map[point.Point]struct{}, error) {
	file, err := os.Open("2019/day10/asteroids.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileBytes), "\n")
	asteroids := make(map[point.Point]struct{})
	for i, line := range lines {
		for j, object := range line {
			if string(object) == "#" {
				asteroids[point.New(j-x, y-i)] = struct{}{}
			}
		}
	}
	return asteroids, nil
}
