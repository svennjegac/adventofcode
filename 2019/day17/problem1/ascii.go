package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/svennjegac/adventofcode/2019/day17"
	"github.com/svennjegac/adventofcode/2019/day17/computer"
)

const (
	newLine = 10
	hash = 35
)

func main() {
	locs := make(map[loc]int)

	time.Sleep(time.Millisecond * 10)

	for i := 0; i < 2; i++ {
		down := 0
		right := 0

		memory, err := day17.Intcode("2019/day17/intcode.txt")
		if err != nil {
			panic(err)
		}

		in := make(chan *big.Int, 1)
		out := make(chan *big.Int)
		comp := computer.NewIntcodeComputer(memory, big.NewInt(0), big.NewInt(0), in, out)

		done := make(chan struct{})
		go func() {
			comp.Run()
			close(done)
		}()

	loop:
		for {
			select {
			case <-done:
				break loop
			case co := <-out:
				o := int(co.Int64())
				if o == newLine {
					down++
					right = 0
				} else {
					l := loc{
						down:  down,
						right: right,
					}
					right++
					locs[l] = o
				}
				fmt.Printf("%c", o)
			}
		}
	}

	res := 0
	for l, v := range locs {
		if v != hash {
			continue
		}
		if !isOfType(l.right - 1, l.down, hash, locs) {
			continue
		}
		if !isOfType(l.right + 1, l.down, hash, locs) {
			continue
		}
		if !isOfType(l.right, l.down - 1, hash, locs) {
			continue
		}
		if !isOfType(l.right, l.down + 1, hash, locs) {
			continue
		}
		res += l.down * l.right
	}

	fmt.Println("Res:", res)
}

func isOfType(r int, d int, t int, locs map[loc]int) bool {
	l := locs[loc{
		down:  d,
		right: r,
	}]
	return l == t
}

type loc struct {
	down int
	right int
}
