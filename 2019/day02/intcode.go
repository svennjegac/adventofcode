package day02

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Intcode() ([]int, error) {
	file, err := os.Open("2019/day02/intcode.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	strcode := strings.Split(string(fileBytes), ",")
	intcode := make([]int, 0, len(strcode))
	for _, s := range strcode {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		intcode = append(intcode, i)
	}
	return intcode, nil
}
