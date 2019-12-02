package day2

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strconv"
)

type Day2 struct {
	byteCode []int
}

func (d *Day2) InputPath() string {
	return "/calendar/day2/input"
}

func (d *Day2) Prepare(input *os.File) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(func (data []byte, atEOF bool) (advance int, token []byte, err error) {
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
		d.byteCode = append(d.byteCode, opCode)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
d.byteCode[1] = 12
d.byteCode[2] = 2
	return nil
}

func (d *Day2) Part1() (string, error) {
	loop: for pc := 0; ; pc+=4 {
		switch d.byteCode[pc] {
		case 1:
			d.byteCode[d.byteCode[pc + 3]] = d.byteCode[d.byteCode[pc + 1]] + d.byteCode[d.byteCode[pc + 2]]
		case 2:
			d.byteCode[d.byteCode[pc + 3]] = d.byteCode[d.byteCode[pc + 1]] * d.byteCode[d.byteCode[pc + 2]]
		default:
			break loop
		}
		pc += 4
	}
	return strconv.Itoa(d.byteCode[0]), nil
}

func (d *Day2) Part2() (string, error) {
	return "", errors.New("todo")
}


