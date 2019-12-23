package day13

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

type Day13 struct {
	computer intcodeComputer
}

func (d Day13) InputPath() string {
	return "/calendar/day13/input"
}

func (d *Day13) Prepare(input io.Reader) error {
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

type insType int

const (
	emptyTile insType = iota
	wallTile
	blockTile
	paddleTile
	ballTile
	setScore
)

type output struct {
	x    int
	y    int
	tile insType
}

func runGame(computer intcodeComputer, inC <-chan int) <-chan output {
	inputsC := make(chan output, 1)
	go func() {
		outC := make(chan int, 3)
		go func() {
			computer.exec(inC, outC)
			close(outC)
		}()
		var input output
		offset := 0
		for ins := range outC {
			switch offset & 0x03 {
			case 0:
				input.x = ins
				offset = 1
			case 1:
				input.y = ins
				offset = 2
			case 2:
				if input.x == -1 && input.y == 0 {
					input.x = ins
					input.tile = setScore
				} else {
					input.tile = insType(ins)
				}
				inputsC <- input
				offset = 0
			}
		}
		close(inputsC)
	}()
	return inputsC
}

func (d Day13) Part1() (string, error) {
	outputC := runGame(d.computer, nil)
	blockTileCount := 0
	for inst := range outputC {
		if inst.tile == blockTile {
			blockTileCount++
		}
	}
	return strconv.Itoa(blockTileCount), nil
}

func clamp(n int, min int, max int) int {
	if n < min {
		return min
	} else if n > max {
		return max
	}
	return n
}

func (d Day13) Part2() (string, error) {
	computer := d.computer
	computer.intcode[0] = 2
	inC := make(chan int, 10)
	outputC := runGame(d.computer, inC)
	score := 0
	step := 0
	var ball, paddle output
	for ins := range outputC {
		switch ins.tile {
		case ballTile:
			ball = ins
			step |= 0b01
		case paddleTile:
			paddle = ins
			step |= 0b10
		case setScore:
			score = ins.x
		}
		if step == 0b11 {
			input := clamp(ball.x-paddle.x, -1, 1)
			if input == 0 {
				step = 0b10
			} else {
				step = 0
			}
			inC <- input
		}
	}
	return strconv.Itoa(score), nil
}
