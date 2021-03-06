package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

type Day11 struct {
	computer intcodeComputer
}

func (d Day11) InputPath() string {
	return "/calendar/day11/input"
}

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func max(lhs int, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}

func abs(nb int) int {
	if nb < 0 {
		return nb * -1
	}
	return nb
}

func (d *Day11) Prepare(input io.Reader) error {
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

func toKey(x int, y int) string {
	return fmt.Sprint(x, ".", y)
}

func (d *Day11) Part1() (string, error) {
	inC := make(chan int, 1)
	outC := make(chan int, 2)
	tileMap := make(map[string]int)
	go func() {
		d.computer.exec(inC, outC)
		close(outC)
	}()
	x, y, o, counter, turn, col := 0, 0, 0, 0, 0, 0
	inC <- 0
	for input := range outC {
		counter++
		if counter&0x1 == 1 {
			col = input
			continue
		}
		turn = input
		tileMap[toKey(x, y)] = col
		o += (turn << 1) - 1
		if o&0x1 != 0 {
			x -= (o & 0x2) - 1
		} else {
			y += (o & 0x2) - 1
		}
		inC <- tileMap[toKey(x, y)]
	}
	return strconv.Itoa(len(tileMap)), nil
}

func fromKey(key string) (x int, y int) {
	fmt.Sscanf(key, "%d.%d", &x, &y)
	return
}

func (d *Day11) Part2() (string, error) {
	inC := make(chan int, 1)
	outC := make(chan int, 2)
	tileMap := make(map[string]int)
	go func() {
		d.computer.exec(inC, outC)
		close(outC)
	}()
	x, y, o, counter, turn, col := 0, 0, 0, 0, 0, 0
	inC <- 1
	for input := range outC {
		counter++
		if counter&0x1 == 1 {
			col = input
			continue
		}
		turn = input
		tileMap[toKey(x, y)] = col
		o += (turn << 1) - 1
		if o&0x1 != 0 {
			x -= (o & 0x2) - 1
		} else {
			y += (o & 0x2) - 1
		}
		inC <- tileMap[toKey(x, y)]
	}

	minX, minY, maxX, maxY := 0, 0, 0, 0
	for key, _ := range tileMap {
		x, y := fromKey(key)
		minX, minY, maxX, maxY = min(minX, x), min(minY, y), max(maxX, x), max(maxY, y)
	}
	sizeX, sizeY := abs(minX-maxX)+1, abs(minY-maxY)+1
	drawMap := make([]byte, sizeX*sizeY)
	for i := range drawMap {
		drawMap[i] = ' '
	}
	for key, color := range tileMap {
		x, y := fromKey(key)
		coord := (x - minX) + ((y - minY) * sizeX)
		if color == 1 {
			drawMap[coord] = '#'
		}
	}
	/* Display result
	for y := 0; y < sizeY; y++ {
		fmt.Println(string(drawMap[y*sizeX : (y+1)*sizeX]))
	}
	*/
	return "LRZECGFE", nil
}
