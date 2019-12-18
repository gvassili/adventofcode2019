package day09

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

type Day9 struct {
	computer intcodeComputer
}

func (d *Day9) InputPath() string {
	return "calendar/day09/input"
}

func (d *Day9) Prepare(input io.Reader) error {
	var program []int
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
		program = append(program, opCode)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	d.computer.init(program)
	return nil
}

func (d *Day9) Part1() (string, error) {
	inC, outC := make(chan int, 1), make(chan int, 1)
	inC <- 1
	d.computer.exec(inC, outC)
	return strconv.Itoa(<-outC), nil
}

func (d *Day9) Part2() (string, error) {
	inC, outC := make(chan int, 1), make(chan int, 1)
	inC <- 2
	d.computer.exec(inC, outC)
	return strconv.Itoa(<-outC), nil
}
