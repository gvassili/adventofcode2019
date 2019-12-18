package day06

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type system map[string]*planet

type planet struct {
	orbit    *planet
	moons    []*planet
	distance int
}

type Day6 struct {
	system system
}

func (d Day6) InputPath() string {
	return "calendar/day06/input"
}

func (d *Day6) Prepare(input io.Reader) error {
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
			orbiter = &planet{}
			d.system[orbiterName] = orbiter
		}
		orbited.moons = append(orbited.moons, orbiter)
		orbiter.orbit = orbited
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

func planetDistance(planet *planet, distance int) {
	if planet.distance != 0 && planet.distance < distance {
		return
	}
	planet.distance = distance
	for _, moon := range planet.moons {
		planetDistance(moon, distance+1)
	}
	if planet.orbit != nil {
		planetDistance(planet.orbit, distance+1)
	}
}

func (d *Day6) Part2() (string, error) {
	planetDistance(d.system["YOU"], -1)
	sanPlanet := d.system["SAN"]
	return strconv.Itoa(sanPlanet.orbit.distance), nil
}
