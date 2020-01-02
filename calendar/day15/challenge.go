package day15

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/gdamore/tcell"
	"io"
	"os"
	"strconv"
	"time"
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

/*

 */

const (
	offsetX = 22
	offsetY = 22
)

func drawDrone(s tcell.Screen, c coordinate, o orientation) {
	p := 'o'
	switch o {
	case 0:
		p = '⬆'
	case 1:
		p = '⮕'
	case 2:
		p = '⬇'
	case 3:
		p = '⬅'

	}
	s.SetContent(int(c.x())+offsetX, int(c.y())+offsetY, p, nil, tcell.StyleDefault)
}

func drawWall(s tcell.Screen, c coordinate) {
	s.SetContent(int(c.x())+offsetX, int(c.y())+offsetY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGray))
}

func drawEmpty(s tcell.Screen, c coordinate) {
	s.SetContent(int(c.x())+offsetX, int(c.y())+offsetY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorLightBlue))
}

func drawBack(s tcell.Screen, c coordinate, d int) {
	if d == 0 {
		s.SetContent(int(c.x())+offsetX, int(c.y())+offsetY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorRed))
	}
	if d == 1 {
		s.SetContent(int(c.x())+offsetX, int(c.y())+offsetY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorGreen))
	}
	if d > 1 {
		s.SetContent(int(c.x())+offsetX, int(c.y())+offsetY, ' ', nil, tcell.StyleDefault.Background(tcell.ColorDarkGreen))
	}
}

func (d *Day15) Part1() (string, error) {
	screen, _ := tcell.NewScreen()
	if err := screen.Init(); err != nil {
		panic(err)
	}
	inC, outC := make(chan int), make(chan int)
	go d.computer.exec(inC, outC)
	var navigate func(c coordinate, o orientation) int
	maze := make(mazeMap)
	go func() {
		for {
			ev := screen.PollEvent()
			e, ok := ev.(*tcell.EventKey)
			if !ok {
				continue
			}
			if ok && e.Key() == tcell.KeyESC {
				screen.Clear()
				screen.Fini()
				os.Exit(1)
			}
			break
		}
	}()
	navigate = func(c coordinate, o orientation) int {
		inC <- o.toInput()
		r := <-outC
		maze[c] = intToBool(r)

		//time.Sleep(time.Millisecond * 10)
		distance := 0
		if r == 0 {
			drawWall(screen, c)
			screen.Show()
			return 0
		} else if r == 2 {
			distance++
		}
		drawEmpty(screen, c.move(o.backward()))
		drawDrone(screen, c, o)
		screen.Show()
		for d, i := o, 0; i < 4 && distance == 0; i++ {
			newC := c.move(d)
			_, ok := maze[c.move(d)]
			if !ok {
				distance += navigate(newC, d)
			}
			d = d.left()
		}

		inC <- o.backward().toInput()
		<-outC
		drawDrone(screen, c.move(o.backward()), o.backward())
		drawBack(screen, c, distance)
		screen.Show()
		return distance
	}
	distance := 0
	for o, i := orientation(0), 0; i < 4; i++ {
		distance += navigate(toCoordinate(0, 0).move(o), o)
		o = o.right()
	}
	time.Sleep(time.Minute)
	return strconv.Itoa(distance), nil
}

func (d *Day15) Part2() (string, error) {
	return "", nil
}
