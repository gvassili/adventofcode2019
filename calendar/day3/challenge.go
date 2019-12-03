package day3

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type direction byte

const (
	up         direction = 1
	down       direction = 1 << 2
	right      direction = 1 << 1
	left       direction = 1 << 3
	vertical             = up | down
	horizontal           = right | left
)

type segment struct {
	x1   int
	y1   int
	x2   int
	y2   int
	size int
	dir  direction
}

func (d direction) String() string {
	return string(d)
}

type wire []segment

type Day3 struct {
	wires []wire
}

func (d Day3) InputPath() string {
	return "/calendar/day3/input"
}

func min(rhs int, lhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}

func max(rhs int, lhs int) int {
	if rhs > lhs {
		return rhs
	}
	return lhs
}

func (d *Day3) Prepare(input *os.File) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		wireSchema := scanner.Text()
		scanner := bufio.NewScanner(strings.NewReader(wireSchema))
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
		wire := make(wire, 0, 64)
		x, y := 0, 0
		for scanner.Scan() {
			var dir direction
			var size int
			segmentSchema := scanner.Text()
			if _, err := fmt.Sscanf(segmentSchema, "%c%d", &dir, &size); err != nil {
				return err
			}
			segment := segment{x, y, x, y, size, dir}
			switch dir {
			case 'R':
				segment.x2 += size
			case 'U':
				segment.y2 += size
			case 'L':
				segment.x2 -= size
			case 'D':
				segment.y2 -= size
			}
			x, y = segment.x2, segment.y2
			wire = append(wire, segment)
		}
		if err := scanner.Err(); err != nil {
			return err
		}
		for i, segment := range wire {
			wire[i].x1 = min(segment.x1, segment.x2)
			wire[i].x2 = max(segment.x1, segment.x2)
			wire[i].y1 = min(segment.y1, segment.y2)
			wire[i].y2 = max(segment.y1, segment.y2)
		}
		d.wires = append(d.wires, wire)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Day3) Part1() (string, error) {
	wire1, wire2 := d.wires[0], d.wires[1]
	minDistance := math.MaxInt64
	for _, seg1 := range wire1 {
		for _, seg2 := range wire2 {
			if (seg1.dir&horizontal != 0) && (seg2.dir&vertical != 0) &&
				(seg2.x1 > seg1.x1 && seg2.x1 < seg1.x2) &&
				(seg1.y1 > seg2.y1 && seg1.y1 < seg2.y2) {
				distance := seg2.x1 + seg1.y1
				minDistance = min(minDistance, distance)
			} else if (seg1.dir&vertical != 0) && (seg2.dir&horizontal != 0) &&
				(seg1.x1 > seg2.x1 && seg1.x1 < seg2.x2) &&
				(seg2.y1 > seg1.y1 && seg2.y1 < seg1.y2) {
				distance := seg1.x1 + seg2.y1
				minDistance = min(minDistance, distance)
			}
		}
	}
	return strconv.Itoa(minDistance), nil
}

func (d *Day3) Part2() (string, error) {
	return "", errors.New("todo")
}
