package day17

import (
	"io/ioutil"
	"math/big"
	"os"
	"strings"
)

func Intcode(f string) (map[string]*big.Int, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	strNums := strings.Split(string(fileBytes), ",")
	memory := make(map[string]*big.Int)
	index := big.NewInt(0)
	for _, strNum := range strNums {
		bigNum, _ := new(big.Int).SetString(strNum, 10)
		memory[index.String()] = bigNum
		index.Add(index, big.NewInt(1))
	}
	return memory, nil
}
