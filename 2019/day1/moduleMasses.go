package day1

import (
	"bufio"
	"os"
	"strconv"
)

func ModuleMasses() ([]int, error) {
	file, err := os.Open("2019/day1/masses.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	masses := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		masses = append(masses, mass)
	}
	return masses, nil
}
