package day03

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"sort"
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
	x1        int
	y1        int
	x2        int
	y2        int
	totalSize int
	distance  int
	dir       direction
}

func (w wire) Less(i, j int) bool {
	return w[i].distance < w[j].distance
}

func (w wire) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w wire) Len() int {
	return len(w)
}

func (d direction) String() string {
	return string(d)
}

type wire []segment

type Day3 struct {
	wires []wire
}

func (d Day3) InputPath() string {
	return "calendar/day03/input"
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

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return n * -1
}

func (d *Day3) Prepare(input io.Reader) error {
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
		totalSize := 0
		for scanner.Scan() {
			var dir direction
			var size int
			segmentSchema := scanner.Text()
			if _, err := fmt.Sscanf(segmentSchema, "%c%d", &dir, &size); err != nil {
				return err
			}
			segment := segment{x, y, x, y, totalSize, x + y, dir}
			totalSize += size
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
		d.wires = append(d.wires, wire)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Day3) Part1() (string, error) {
	wire1 := make(wire, len(d.wires[0]))
	wire2 := make(wire, len(d.wires[1]))
	copy(wire1, d.wires[0])
	copy(wire2, d.wires[1])
	minDistance := math.MaxInt64
	sort.Sort(wire1)
	sort.Sort(wire2)
	for _, seg1 := range wire1 {
		if seg1.distance > minDistance {
			break
		}
		for _, seg2 := range wire2 {
			if seg2.distance > minDistance {
				break
			}
			if (seg1.dir&horizontal != 0) && (seg2.dir&vertical != 0) &&
				(seg2.x1 > min(seg1.x1, seg1.x2) && seg2.x1 < max(seg1.x1, seg1.x2)) &&
				(seg1.y1 > min(seg2.y1, seg2.y2) && seg1.y1 < max(seg2.y1, seg2.y2)) {
				distance := seg2.x1 + seg1.y1
				minDistance = min(minDistance, distance)
			} else if (seg1.dir&vertical != 0) && (seg2.dir&horizontal != 0) &&
				(seg1.x1 > min(seg2.x1, seg2.x2) && seg1.x1 < max(seg2.x1, seg2.x2)) &&
				(seg2.y1 > min(seg1.y1, seg1.y2) && seg2.y1 < max(seg1.y1, seg1.y2)) {
				distance := seg1.x1 + seg2.y1
				minDistance = min(minDistance, distance)
			}
		}
	}
	return strconv.Itoa(minDistance), nil
}

func (d *Day3) Part2() (string, error) {
	wire1, wire2 := d.wires[0], d.wires[1]
	minDistance := math.MaxInt64
	for _, seg1 := range wire1 {
		for _, seg2 := range wire2 {
			if seg1.totalSize+seg2.totalSize > minDistance {
				break
			}
			if (seg1.dir&horizontal != 0) && (seg2.dir&vertical != 0) &&
				(seg2.x1 > min(seg1.x1, seg1.x2) && seg2.x1 < max(seg1.x1, seg1.x2)) &&
				(seg1.y1 > min(seg2.y1, seg2.y2) && seg1.y1 < max(seg2.y1, seg2.y2)) {
				distance := seg2.totalSize + seg1.totalSize + abs(seg1.x1-seg2.x1) + abs(seg2.y1-seg1.y1)
				minDistance = min(minDistance, distance)
			} else if (seg1.dir&vertical != 0) && (seg2.dir&horizontal != 0) &&
				(seg1.x1 > min(seg2.x1, seg2.x2) && seg1.x1 < max(seg2.x1, seg2.x2)) &&
				(seg2.y1 > min(seg1.y1, seg2.y2) && seg2.y1 < max(seg1.y1, seg1.y2)) {
				distance := seg2.totalSize + seg1.totalSize + abs(seg2.x1-seg1.x1) + abs(seg1.y1-seg2.y1)
				minDistance = min(minDistance, distance)
			}
		}
	}
	return strconv.Itoa(minDistance), nil
}
