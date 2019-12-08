package day8

import (
	"io/ioutil"
	"os"
	"strconv"
)

func Image() ([]int, error) {
	file, err := os.Open("2019/day8/image.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	img := make([]int, 0, len(fileBytes))
	for i := 0; i < len(fileBytes); i++ {
		digit, err := strconv.Atoi(string(fileBytes[i]))
		if err != nil {
			return nil, err
		}
		img = append(img, digit)
	}
	return img, nil
}
