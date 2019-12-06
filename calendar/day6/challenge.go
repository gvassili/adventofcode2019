package day6

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type system map[string]*planet

type planet struct {
	orbit *planet
}

type Day6 struct {
	system system
}

func (d Day6) InputPath() string {
	return "/calendar/day6/input"
}

func (d *Day6) Prepare(input *os.File) error {
	scanner := bufio.NewScanner(input)
	d.system = make(system, 256)
	for scanner.Scan() {
		tokens := strings.SplitN(scanner.Text(), ")", 2)
		orbitedName, orbiterName := tokens[0], tokens[1]
		orbited := d.system[orbitedName]
		orbiter := d.system[orbiterName]
		if orbited == nil {
			orbited = &planet{}
			d.system[orbitedName] = orbited
		}
		if orbiter == nil {
			orbiter = &planet{orbited}
			d.system[orbiterName] = orbiter
		} else {
			orbiter.orbit = orbited
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func planetOrbitCount(planet *planet) int {
	if planet.orbit != nil {
		return planetOrbitCount(planet.orbit) + 1
	}
	return 0
}

func (d *Day6) Part1() (string, error) {
	checkSum := 0
	for _, value := range d.system {
		checkSum += planetOrbitCount(value)
	}
	return strconv.Itoa(checkSum), nil
}

func (d *Day6) Part2() (string, error) {
	return "", errors.New("todo")
}
