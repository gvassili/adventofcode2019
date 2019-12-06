package day5

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Day5 struct {
	program []int
	memory  []int
}

const (
	OpAdd = 1
	OpMul = 2
	OpInp = 3
	OpOut = 4
	OpExt = 99
)

func (d Day5) InputPath() string {
	return "/calendar/day5/input"
}

func (d *Day5) Prepare(input *os.File) error {
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
	for i := offset; i > 1; i-- {
		pmode /= 10
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
	for pc := 0; ; {
		instruction := d.memory[pc]
		opcode := instruction % 100
		pmode := instruction / 100
		switch opcode {
		case OpAdd:
			p1, p2, p3 := d.getValue(pc, pmode, 1), d.getValue(pc, pmode, 2), d.memory[pc+3]
			d.memory[p3] = p1 + p2
			pc += 4
		case OpMul:
			p1, p2, p3 := d.getValue(pc, pmode, 1), d.getValue(pc, pmode, 2), d.memory[pc+3]
			d.memory[p3] = p1 * p2
			pc += 4
		case OpInp:
			d.memory[d.memory[pc+1]] = input[0]
			input = input[1:]
			pc += 2
		case OpOut:
			output = d.getValue(pc, pmode, 1)
			pc += 2
		case OpExt:
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
	return "", errors.New("todo")
}
