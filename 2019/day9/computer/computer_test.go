package computer_test

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/svennjegac/adventofcode/2019/day9"
	"github.com/svennjegac/adventofcode/2019/day9/computer"
)

func TestOutputItself(t *testing.T) {
	memory, err := day9.Intcode("tests/resources/output_itself.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int, 100)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)
	comp.Run()

	res := make([]int, 0)
	for i := 0; i < 16; i++ {
		res = append(res, int((<-out).Int64()))
	}

	if !reflect.DeepEqual(res, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}) {
		fmt.Println(res)
		t.Fatal("output itself failed")
	}
}

func Test16DigitNumber(t *testing.T) {
	memory, err := day9.Intcode("tests/resources/s16_dig_num.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int, 100)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)
	comp.Run()

	res := make([]int, 0)
	for i := 0; i < 1; i++ {
		res = append(res, int((<-out).Int64()))
	}

	if !reflect.DeepEqual(res, []int{1219070632396864}) {
		fmt.Println(res)
		t.Fatal("16 dig num failed")
	}
}

func TestMiddleNum(t *testing.T) {
	memory, err := day9.Intcode("tests/resources/middle_num.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int, 100)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)
	comp.Run()

	res := make([]int, 0)
	for i := 0; i < 1; i++ {
		res = append(res, int((<-out).Int64()))
	}

	if !reflect.DeepEqual(res, []int{1125899906842624}) {
		fmt.Println(res)
		t.Fatal("middle num failed")
	}
}

func TestMyInputProblem1(t *testing.T) {
	memory, err := day9.Intcode("tests/resources/my_input.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int, 100)
	in <- big.NewInt(1)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)
	comp.Run()

	res := make([]int, 0)
	for i := 0; i < 1; i++ {
		res = append(res, int((<-out).Int64()))
	}

	if !reflect.DeepEqual(res, []int{3454977209}) {
		fmt.Println(res)
		t.Fatal("my input v1 failed")
	}
}

func TestMyInputProblem2(t *testing.T) {
	memory, err := day9.Intcode("tests/resources/my_input.txt")
	if err != nil {
		panic(err)
	}

	in := make(chan *big.Int, 1)
	out := make(chan *big.Int, 100)
	in <- big.NewInt(2)
	comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)
	comp.Run()

	res := make([]int, 0)
	for i := 0; i < 1; i++ {
		res = append(res, int((<-out).Int64()))
	}

	if !reflect.DeepEqual(res, []int{50120}) {
		fmt.Println(res)
		t.Fatal("my input v2 failed")
	}
}
