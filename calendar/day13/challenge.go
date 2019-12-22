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

type tileType int

const (
	emptyTile tileType = iota
	wallTile
	blockTile
	horizontalPaddleTile
	ballTile
)

type input struct {
	x    int
	y    int
	tile tileType
}

func (d Day13) Part1() (string, error) {
	inputsC := make(chan input, 1)
	go func() {
		outC := make(chan int, 3)
		go func() {
			d.computer.exec(nil, outC)
			close(outC)
		}()
		var input input
		offset := 0
		for inst := range outC {
			switch offset & 0x03 {
			case 0:
				input.x = inst
				offset = 1
			case 1:
				input.y = inst
				offset = 2
			case 2:
				input.tile = tileType(inst)
				inputsC <- input
				offset = 0
			}
		}
		close(inputsC)
	}()
	blockTileCount := 0
	for in := range inputsC {
		if in.tile == blockTile {
			blockTileCount++
		}
	}
	return strconv.Itoa(blockTileCount), nil
}

func (d Day13) Part2() (string, error) {
	return "", nil
}
