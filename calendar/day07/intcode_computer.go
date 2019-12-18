package day07

import (
	"fmt"
)

const (
	opAdd = 1
	opMul = 2
	opInp = 3
	opOut = 4
	opJnz = 5
	opJz  = 6
	opLt  = 7
	opE   = 8
	opRet = 99
)

type intcodeComputer struct {
	intcode []int
	mem     []int
	pc      int
}

func (c *intcodeComputer) init(intcode []int) {
	c.intcode = intcode
	c.mem = make([]int, len(intcode))
}

func (c *intcodeComputer) getValue(pmode int, offset int) int {
	if pmode != 0 {
		switch offset {
		case 2:
			pmode /= 10
		case 3:
			pmode /= 100
		}
	}
	if pmode&1 == 1 {
		return c.mem[c.pc+offset]
	} else {
		return c.mem[c.mem[c.pc+offset]]
	}
}

func (c *intcodeComputer) exec(inC <-chan int, outC chan<- int) error {
	copy(c.mem, c.intcode)
loop:
	for c.pc = 0; ; {
		instruction := c.mem[c.pc]
		opcode := instruction % 100
		pmode := instruction / 100
		switch opcode {
		case opAdd:
			p1, p2, p3 := c.getValue(pmode, 1), c.getValue(pmode, 2), c.mem[c.pc+3]
			c.mem[p3] = p1 + p2
			c.pc += 4
		case opMul:
			p1, p2, p3 := c.getValue(pmode, 1), c.getValue(pmode, 2), c.mem[c.pc+3]
			c.mem[p3] = p1 * p2
			c.pc += 4
		case opInp:
			input := <-inC
			c.mem[c.mem[c.pc+1]] = input
			c.pc += 2
		case opOut:
			output := c.getValue(pmode, 1)
			outC <- output
			c.pc += 2
		case opJnz:
			p1, p2 := c.getValue(pmode, 1), c.getValue(pmode, 2)
			if p1 != 0 {
				c.pc = p2
			} else {
				c.pc += 3
			}
		case opJz:
			p1, p2 := c.getValue(pmode, 1), c.getValue(pmode, 2)
			if p1 == 0 {
				c.pc = p2
			} else {
				c.pc += 3
			}
		case opLt:
			p1, p2, p3 := c.getValue(pmode, 1), c.getValue(pmode, 2), c.mem[c.pc+3]
			if p1 < p2 {
				c.mem[p3] = 1
			} else {
				c.mem[p3] = 0
			}
			c.pc += 4
		case opE:
			p1, p2, p3 := c.getValue(pmode, 1), c.getValue(pmode, 2), c.mem[c.pc+3]
			if p1 == p2 {
				c.mem[p3] = 1
			} else {
				c.mem[p3] = 0
			}
			c.pc += 4
		case opRet:
			break loop
		default:
			return fmt.Errorf("invalid opcode %d", opcode)
		}
	}
	return nil
}
