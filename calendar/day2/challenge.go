package day2

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
)

type Day2 struct {
	program []int
	memory  []int
}

func (d *Day2) InputPath() string {
	return "calendar/day2/input"
}

func (d *Day2) Prepare(input io.Reader) error {
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

func (d *Day2) runByteCode(noun int, verb int) int {
	copy(d.memory, d.program)
	d.memory[1] = noun
	d.memory[2] = verb
loop:
	for pc := 0; ; pc += 4 {
		switch d.memory[pc] {
		case 1:
			d.memory[d.memory[pc+3]] = d.memory[d.memory[pc+1]] + d.memory[d.memory[pc+2]]
		case 2:
			d.memory[d.memory[pc+3]] = d.memory[d.memory[pc+1]] * d.memory[d.memory[pc+2]]
		default:
			break loop
		}
	}
	return d.memory[0]
}

func (d *Day2) Part1() (string, error) {
	d.runByteCode(12, 2)
	return strconv.Itoa(d.memory[0]), nil
}

func (d *Day2) Part2() (string, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			res := d.runByteCode(noun, verb)
			if res == 19690720 {
				return strconv.Itoa(noun) + strconv.Itoa(verb), nil
			}
		}
	}
	return "", errors.New("could not find result")
}
