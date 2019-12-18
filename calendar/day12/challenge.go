package day12

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type body struct {
	x, y, z    int
	vx, vy, vz int
}

type Day12 struct {
	system []body
}

func (d Day12) InputPath() string {
	return "/calendar/day12/input"
}

func (d *Day12) Prepare(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var body body
		_, err := fmt.Sscanf(scanner.Text(), "<x=%d, y=%d, z=%d>", &body.x, &body.y, &body.z)
		if err != nil {
			return err
		}
		d.system = append(d.system, body)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func abs(nb int) int {
	if nb < 0 {
		return nb * -1
	}
	return nb
}

func (b *body) applyGravity(other body) {
	if b.x < other.x {
		b.vx += 1
	} else if b.x > other.x {
		b.vx -= 1
	}
	if b.y < other.y {
		b.vy += 1
	} else if b.y > other.y {
		b.vy -= 1
	}
	if b.z < other.z {
		b.vz += 1
	} else if b.z > other.z {
		b.vz -= 1
	}
}

func (b *body) move() {
	b.x, b.y, b.z = b.x+b.vx, b.y+b.vy, b.z+b.vz
}

func (b *body) energy() int {
	return (abs(b.x) + abs(b.y) + abs(b.z)) * (abs(b.vx) + abs(b.vy) + abs(b.vz))
}

func (d *Day12) Part1() (string, error) {
	for step := 0; step < 1000; step++ {
		for i := 0; i < len(d.system)-1; i++ {
			for n := i + 1; n < len(d.system); n++ {
				d.system[i].applyGravity(d.system[n])
				d.system[n].applyGravity(d.system[i])
			}
		}
		for i := range d.system {
			d.system[i].move()
		}
	}
	sumEnergy := 0
	for _, body := range d.system {
		sumEnergy += body.energy()
	}
	return strconv.Itoa(sumEnergy), nil
}

func (d *Day12) Part2() (string, error) {
	return "", nil
}
