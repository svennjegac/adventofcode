package day16

import (
	"io/ioutil"
	"os"
	"strconv"
)

func Input(f string) []int {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	res := make([]int, 0, len(fileBytes))
	for _, s := range string(fileBytes) {
		num, err := strconv.Atoi(string(s))
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}
