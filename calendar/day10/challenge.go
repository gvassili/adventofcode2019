package day10

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

type asteroids []point

type point struct {
	x, y int
}

type Day10 struct {
	asteroids asteroids
}

func (p point) distance(origin point) int {
	return abs(p.y-origin.y) + abs(p.x-origin.x)
}

func (p point) angle(origin point) float64 {
	theta := math.Atan2(float64(p.y-origin.y), float64(p.x-origin.x))
	angle := (theta / (math.Pi / 180)) + 90
	angle = (angle) + math.Ceil(-angle/360)*360
	return angle
}

func (d Day10) InputPath() string {
	return "/calendar/day10/input"
}

func (d *Day10) Prepare(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	for y := 0; scanner.Scan(); y++ {
		belt := scanner.Bytes()
		for x, c := range belt {
			if c == '#' {
				d.asteroids = append(d.asteroids, point{x, y})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Day10) Part1() (string, error) {
	bestPosition := point{}
	mostSeen := 0
	for i, candidate := range d.asteroids {
		seenMap := make(map[float64]point)
		for n, asteroid := range d.asteroids {
			if i != n {
				seenMap[asteroid.angle(candidate)] = asteroid
			}
		}
		if mostSeen < len(seenMap) {
			mostSeen = len(seenMap)
			bestPosition = candidate
		}
	}
	// fmt.Printf("best station position %+v\n", bestPosition)
	return fmt.Sprintf("%d (x:%d, y:%d)", mostSeen, bestPosition.x, bestPosition.y), nil
}

type VisionVector struct {
	angle     float64
	asteroids asteroids
}
type VisionVectors []VisionVector

func (v VisionVectors) Len() int {
	return len(v)
}

func (v VisionVectors) Less(i, j int) bool {
	return v[j].angle > v[i].angle
}

func (v VisionVectors) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v asteroids) Len() int {
	return len(v)
}

func (v asteroids) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

var station = point{11, 13}

func (v asteroids) Less(i, j int) bool {
	return v[j].distance(station) < v[i].distance(station)
}
func (d *Day10) Part2() (string, error) {
	seenMap := make(map[float64]asteroids)
	for _, asteroid := range d.asteroids {
		if asteroid.x == station.x && asteroid.y == station.y {
			continue
		}
		angle := asteroid.angle(station)
		seenMap[angle] = append(seenMap[angle], asteroid)
	}
	seenArray := make(VisionVectors, 0, len(seenMap))
	for k, v := range seenMap {
		vector := VisionVector{k, v}
		sort.Sort(vector.asteroids)
		seenArray = append(seenArray, vector)
	}
	sort.Sort(seenArray)
	for i, destroyed := 0, 0; ; {
		vector := seenArray[i]
		if len(vector.asteroids) > 0 {
			destroyed++
			if destroyed == 200 {
				asteroid := vector.asteroids[len(vector.asteroids)-1]
				return strconv.Itoa(asteroid.x*100 + asteroid.y), nil
			}
			seenArray[i].asteroids = vector.asteroids[:len(vector.asteroids)-1]
		}
		i++
		if i == len(seenArray) {
			i = 0
		}
	}
}
