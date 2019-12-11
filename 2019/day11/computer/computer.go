package computer

import (
	"math/big"
)

func NewIntcodeComputer(memory map[string]*big.Int, pc, rb *big.Int, in, out chan *big.Int) *IntcodeComputer {
	return &IntcodeComputer{
		memory: memory,
		pc:     pc,
		rb:     rb,
		in:     in,
		out:    out,
	}
}

type IntcodeComputer struct {
	memory map[string]*big.Int
	pc     *big.Int
	rb     *big.Int

	in  chan *big.Int
	out chan *big.Int
}

func (e *IntcodeComputer) Run() {
	for done := e.next(); done == false; done = e.next() {
	}
}

func (e *IntcodeComputer) next() bool {
	instruction, modes := e.opcode()

	switch instruction {
	case add:
		resAddr := e.addr(3, modes[2])
		op1 := e.val(1, modes[0])
		op2 := e.val(2, modes[1])
		e.memory[resAddr] = new(big.Int).Add(op1, op2)
		e.pc.Add(e.pc, big.NewInt(4))
		return false
	case mul:
		resAddr := e.addr(3, modes[2])
		op1 := e.val(1, modes[0])
		op2 := e.val(2, modes[1])
		e.memory[resAddr] = new(big.Int).Mul(op1, op2)
		e.pc.Add(e.pc, big.NewInt(4))
		return false
	case in:
		resAddr := e.addr(1, modes[0])
		e.memory[resAddr] = copyBigInt(<-e.in)
		e.pc.Add(e.pc, big.NewInt(2))
		return false
	case out:
		e.out <- copyBigInt(e.val(1, modes[0]))
		e.pc.Add(e.pc, big.NewInt(2))
		return false
	case jmpIfTrue:
		op1 := e.val(1, modes[0])
		op2 := e.val(2, modes[1])
		if op1.String() != "0" {
			e.pc = copyBigInt(op2)
			return false
		}
		e.pc.Add(e.pc, big.NewInt(3))
		return false
	case jmpIfFalse:
		op1 := e.val(1, modes[0])
		op2 := e.val(2, modes[1])
		if op1.String() == "0" {
			e.pc = copyBigInt(op2)
			return false
		}
		e.pc.Add(e.pc, big.NewInt(3))
		return false
	case lessThan:
		op1 := e.val(1, modes[0])
		op2 := e.val(2, modes[1])
		if op1.Cmp(op2) < 0 {
			e.memory[e.addr(3, modes[2])] = big.NewInt(1)
		} else {
			e.memory[e.addr(3, modes[2])] = big.NewInt(0)
		}
		e.pc.Add(e.pc, big.NewInt(4))
		return false
	case equals:
		op1 := e.val(1, modes[0])
		op2 := e.val(2, modes[1])
		if op1.Cmp(op2) == 0 {
			e.memory[e.addr(3, modes[2])] = big.NewInt(1)
		} else {
			e.memory[e.addr(3, modes[2])] = big.NewInt(0)
		}
		e.pc.Add(e.pc, big.NewInt(4))
		return false
	case adjRelBase:
		op1 := e.val(1, modes[0])
		e.rb.Add(e.rb, op1)
		e.pc.Add(e.pc, big.NewInt(2))
		return false
	case halt:
		return true
	default:
		panic("unknown next")
	}
}

func (e *IntcodeComputer) opcode() (instruction, []mode) {
	opcodeBigInt := e.memory[e.pc.String()]
	opcode := opcodeBigInt.Int64()

	instr := instruction(opcode % 100)

	opcode /= 100
	modes := make([]mode, 0)
	for i := 0; i < 3; i++ {
		modes = append(modes, mode(opcode%10))
		opcode /= 10
	}

	return instr, modes
}

func (e *IntcodeComputer) addr(offset int64, m mode) string {
	switch m {
	case position:
		return e.relToPc(offset).String()
	case relative:
		return new(big.Int).Add(e.rb, e.relToPc(offset)).String()
	default:
		panic("address resolving should not happen")
	}
}

func (e *IntcodeComputer) val(offset int64, m mode) *big.Int {
	switch m {
	case position:
		if res, ok := e.memory[e.relToPc(offset).String()]; ok {
			return res
		}
		return big.NewInt(0)
	case immediate:
		return e.relToPc(offset)
	case relative:
		return e.relToRb(e.relToPc(offset).Int64())
	default:
		panic("unknown memory access mode")
	}
}

func (e *IntcodeComputer) relToPc(offset int64) *big.Int {
	return e.relTo(offset, e.pc)
}

func (e *IntcodeComputer) relToRb(offset int64) *big.Int {
	return e.relTo(offset, e.rb)
}

func (e *IntcodeComputer) relTo(offset int64, memoryAddr *big.Int) *big.Int {
	addr := new(big.Int).Add(memoryAddr, big.NewInt(offset)).String()
	if val, ok := e.memory[addr]; ok {
		return val
	}
	return big.NewInt(0)
}

func copyBigInt(b *big.Int) *big.Int {
	if b == nil {
		return big.NewInt(0)
	}
	b2, _ := new(big.Int).SetString(b.String(), 10)
	return b2
}

type instruction int

const (
	_ instruction = iota
	add
	mul
	in
	out
	jmpIfTrue
	jmpIfFalse
	lessThan
	equals
	adjRelBase
	halt = 99
)

type mode int

const (
	position mode = iota
	immediate
	relative
)
