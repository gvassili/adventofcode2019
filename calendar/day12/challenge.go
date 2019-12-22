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

func boolToInt(cond bool) int {
	if cond == false {
		return 0
	}
	return 1
}

func (b *body) alignment(other body) int {
	return boolToInt(b.x == other.x) |
		(boolToInt(b.y == other.y) << 1) |
		(boolToInt(b.z == other.z) << 2)
}

func (b *body) energy() int {
	return (abs(b.x) + abs(b.y) + abs(b.z)) * (abs(b.vx) + abs(b.vy) + abs(b.vz))
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func (d *Day12) Part1() (string, error) {
	system := make([]body, len(d.system))
	copy(system, d.system)
	for step := 0; step < 1000; step++ {
		for i := 0; i < len(system)-1; i++ {
			for n := i + 1; n < len(system); n++ {
				system[i].applyGravity(system[n])
				system[n].applyGravity(system[i])
			}
		}
		for i := range system {
			system[i].move()
		}
	}
	sumEnergy := 0
	for _, body := range system {
		sumEnergy += body.energy()
	}
	return strconv.Itoa(sumEnergy), nil
}

func (d *Day12) Part2() (string, error) {
	system := make([]body, len(d.system))
	ax, ay, az, afound := 0, 0, 0, 0
	copy(system, d.system)
	for step := 2; afound != 0x7; step++ {
		for i := 0; i < len(system)-1; i++ {
			for n := i + 1; n < len(system); n++ {
				system[i].applyGravity(system[n])
				system[n].applyGravity(system[i])
			}
		}
		alignment := 0x7
		for i := range system {
			system[i].move()
			alignment &= system[i].alignment(d.system[i])
		}
		if ax == 0 && (alignment&1) != 0 {
			ax = step
		}
		if ay == 0 && (alignment&(1<<1)) != 0 {
			ay = step
		}
		if az == 0 && (alignment&(1<<2)) != 0 {
			az = step
		}
		afound |= alignment
	}
	return strconv.Itoa(lcm(lcm(ax, ay), az)), nil
}
