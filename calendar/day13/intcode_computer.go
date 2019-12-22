package day13

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
	opRb  = 9
	opRet = 99
)

type intcodeComputer struct {
	intcode []int
	mem     []int
	pc      int
	relbase int
	pmode   pmode
}

func (c *intcodeComputer) init(intcode []int) {
	c.intcode = intcode
}

func (c *intcodeComputer) read() int {
	pmode := c.pmode.next()
	addr := 0
	switch pmode {
	case 0:
		addr = c.mem[c.next()]
	case 1:
		addr = c.next()
	case 2:
		addr = c.mem[c.next()] + c.relbase
	default:
		panic(fmt.Errorf("invalid parameter mode %d", pmode))
	}
	return c.mem[addr]
}

func (c *intcodeComputer) write(value int) {
	pmode := c.pmode.next()
	addr := 0
	switch pmode {
	case 0:
		addr = c.mem[c.next()]
	case 2:
		p := c.mem[c.next()]
		addr = p + c.relbase
	default:
		panic(fmt.Errorf("invalid relative parameter mode %d", pmode))
	}
	c.mem[addr] = value
}

func (c *intcodeComputer) next() int {
	i := c.pc
	c.pc++
	return i
}

type pmode int

func (p *pmode) next() int {
	i := *p % 10
	*p /= 10
	return int(i)
}

func (c *intcodeComputer) exec(inC <-chan int, outC chan<- int) {
	c.relbase = 0
	c.mem = make([]int, 16384)
	copy(c.mem, c.intcode)
loop:
	for c.pc = 0; ; {
		instruction := c.mem[c.next()]
		opcode := instruction % 100
		c.pmode = pmode(instruction / 100)
		switch opcode {
		case opAdd:
			p1, p2 := c.read(), c.read()
			c.write(p1 + p2)
		case opMul:
			p1, p2 := c.read(), c.read()
			c.write(p1 * p2)
		case opInp:
			input := <-inC
			c.write(input)
		case opOut:
			output := c.read()
			outC <- output
		case opJnz:
			p1, p2 := c.read(), c.read()
			if p1 != 0 {
				c.pc = p2
			}
		case opJz:
			p1, p2 := c.read(), c.read()
			if p1 == 0 {
				c.pc = p2
			}
		case opLt:
			p1, p2 := c.read(), c.read()
			if p1 < p2 {
				c.write(1)
			} else {
				c.write(0)
			}
		case opE:
			p1, p2 := c.read(), c.read()
			if p1 == p2 {
				c.write(1)
			} else {
				c.write(0)
			}
		case opRb:
			p1 := c.read()
			c.relbase += p1
		case opRet:
			break loop
		}
	}
}
