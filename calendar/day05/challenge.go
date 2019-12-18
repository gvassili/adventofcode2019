package day05

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Day5 struct {
	program []int
	memory  []int
	pc      int
}

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

func (d Day5) InputPath() string {
	return "calendar/day05/input"
}

func (d *Day5) Prepare(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, ','); i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		opCode, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d.program = append(d.program, opCode)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	d.memory = make([]int, len(d.program))
	return nil
}

func (d *Day5) getValue(pc int, pmode int, offset int) int {
	if pmode != 0 {
		switch offset {
		case 2:
			pmode /= 10
		case 3:
			pmode /= 100
		}
	}
	if pmode&1 == 1 {
		return d.memory[pc+offset]
	} else {
		return d.memory[d.memory[pc+offset]]
	}
}

func (d *Day5) runByteCode(input []int) (int, error) {
	copy(d.memory, d.program)
	output := 0
loop:
	for d.pc = 0; ; {
		instruction := d.memory[d.pc]
		opcode := instruction % 100
		pmode := instruction / 100
		switch opcode {
		case opAdd:
			p1, p2, p3 := d.getValue(d.pc, pmode, 1), d.getValue(d.pc, pmode, 2), d.memory[d.pc+3]
			d.memory[p3] = p1 + p2
			d.pc += 4
		case opMul:
			p1, p2, p3 := d.getValue(d.pc, pmode, 1), d.getValue(d.pc, pmode, 2), d.memory[d.pc+3]
			d.memory[p3] = p1 * p2
			d.pc += 4
		case opInp:
			d.memory[d.memory[d.pc+1]] = input[0]
			input = input[1:]
			d.pc += 2
		case opOut:
			output = d.getValue(d.pc, pmode, 1)
			d.pc += 2
		case opJnz:
			p1, p2 := d.getValue(d.pc, pmode, 1), d.getValue(d.pc, pmode, 2) //d.memory[d.pc+2]
			if p1 != 0 {
				d.pc = p2
			} else {
				d.pc += 3
			}
		case opJz:
			p1, p2 := d.getValue(d.pc, pmode, 1), d.getValue(d.pc, pmode, 2) //d.memory[d.pc+2]
			if p1 == 0 {
				d.pc = p2
			} else {
				d.pc += 3
			}
		case opLt:
			p1, p2, p3 := d.getValue(d.pc, pmode, 1), d.getValue(d.pc, pmode, 2), d.memory[d.pc+3]
			if p1 < p2 {
				d.memory[p3] = 1
			} else {
				d.memory[p3] = 0
			}
			d.pc += 4
		case opE:
			p1, p2, p3 := d.getValue(d.pc, pmode, 1), d.getValue(d.pc, pmode, 2), d.memory[d.pc+3]
			if p1 == p2 {
				d.memory[p3] = 1
			} else {
				d.memory[p3] = 0
			}
			d.pc += 4
		case opRet:
			break loop
		default:
			return 0, fmt.Errorf("invalid opcode %d", opcode)
		}
	}
	return output, nil
}

func (d *Day5) Part1() (string, error) {
	result, err := d.runByteCode([]int{1})
	return strconv.Itoa(result), err
}

func (d *Day5) Part2() (string, error) {
	result, err := d.runByteCode([]int{5})
	return strconv.Itoa(result), err
}
