package day15

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Day15 struct {
	computer intcodeComputer
}

func (d *Day15) InputPath() string {
	return "calendar/day15/input"
}

func (d *Day15) Prepare(input io.Reader) error {
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

type coordinate uint64

type orientation int

func (o orientation) backward() orientation {
	return (o - 2) & 0b11
}

func (o orientation) left() orientation {
	return (o - 1) & 0b11
}

func (o orientation) right() orientation {
	return (o + 1) & 0b11
}

func (o orientation) toInput() int {
	switch o {
	case 0:
		return 1
	case 1:
		return 3
	case 2:
		return 2
	case 3:
		return 4
	}
	panic(errors.New("unknown orientation"))
}

type mazeMap map[coordinate]bool

func toCoordinate(x, y int32) coordinate {
	return coordinate(uint64(x)<<32 | (uint64(y) & 0xffffffff))
}

func (c coordinate) String() string {
	return fmt.Sprintf("x: %d, y:%d", c.x(), c.y())
}
func (c coordinate) x() int32 {
	return int32(c >> 32)
}

func (c coordinate) y() int32 {
	return int32(c & 0xffffffff)
}

func (c coordinate) prev(o orientation) coordinate {
	return c.move((o - 2) & 0b11)
}

func (c coordinate) move(o orientation) coordinate {
	x, y := c.x(), c.y()
	x -= int32(((o & 0b10) - 1) & (0 - (o & 0b1)))
	y += int32(((o & 0b10) - 1) & (0 - ((o & 0b1) ^ 0b1)))
	return toCoordinate(x, y)
}

func intToBool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func (d *Day15) Part1() (string, error) {
	inC, outC := make(chan int, 1), make(chan int, 1)
	go d.computer.exec(inC, outC)
	var navigate func(c coordinate, o orientation) int
	maze := make(mazeMap)
	navigate = func(c coordinate, o orientation) int {
		inC <- o.toInput()
		r := <-outC
		maze[c] = intToBool(r)
		if r == 0 {
			return 0
		} else if r == 2 {
			return 1
		}
		for d, i := o, 0; i < 4; i++ {
			newC := c.move(d)
			_, ok := maze[c.move(d)]
			if !ok {
				dis := navigate(newC, d)
				if dis != 0 {
					return dis + 1
				}
			}
			d = d.left()
		}
		inC <- o.backward().toInput()
		<-outC
		return 0
	}
	distance := 0
	for o, i := orientation(0), 0; i < 4; i++ {
		distance += navigate(toCoordinate(0, 0).move(o), o)
		o = o.right()
	}
	return strconv.Itoa(distance), nil
}

func (d *Day15) Part2() (string, error) {
	return "", nil
}
