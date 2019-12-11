package day10

import (
	"bufio"
	"container/list"
	"errors"
	"io"
	"strconv"
)

type point struct {
	x, y int
}

type asteroid point

type Day10 struct {
	asteroidBelt *list.List
}

func (d Day10) InputPath() string {
	return "/calendar/day10/input"
}

func (d *Day10) Prepare(input io.Reader) error {
	d.asteroidBelt = list.New()
	scanner := bufio.NewScanner(input)
	for y := 0; scanner.Scan(); y++ {
		belt := scanner.Bytes()
		for x, c := range belt {
			if c == '#' {
				d.asteroidBelt.PushBack(asteroid{x, y})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func findBestPosition(station asteroid, asteroidBeltSrc *list.List) int {
	asteroidBelt := list.New()
	asteroidBelt.PushBackList(asteroidBeltSrc)
	front := asteroidBelt.Front()
	for front != nil {
		asteroid1 := front.Value.(asteroid)
		a1 := asteroid{asteroid1.x - station.x, asteroid1.y - station.y}
		for e := front.Next(); e != nil; {
			a2 := e.Value.(asteroid)
			a2 = asteroid{a2.x - station.x, a2.y - station.y}
			//	fmt.Printf("%v %v\t(%v) (%v) =>> %v\n", a1, a2, (a1.x^a2.x) >= 0 || (a1.x*a2.x) == 0, (a1.y^a2.y) >= 0 || (a1.y*a2.y) == 0, a1.x*a2.y-a1.y*a2.x == 0)
			if ((a1.x^a2.x) >= 0 || (a1.x*a2.x) == 0) && ((a1.y^a2.y) >= 0 || (a1.y*a2.y) == 0) && a1.x*a2.y-a1.y*a2.x == 0 {
				if abs(a1.x)+abs(a1.y) < abs(a2.x)+abs(a2.y) {
					tmp := e.Next()
					asteroidBelt.Remove(e)
					e = tmp
				} else {
					goto skip
				}
			} else {
				e = e.Next()
			}
		}
		front = front.Next()
		continue
	skip:
		tmp := front.Next()
		asteroidBelt.Remove(front)
		front = tmp
	}
	return asteroidBelt.Len()
}

func (d *Day10) Part1() (string, error) {
	best := 0
	for e := d.asteroidBelt.Front(); e != nil; e = e.Next() {
		asteroid := e.Value.(asteroid)
		tmp := e.Prev()
		d.asteroidBelt.Remove(e)
		count := findBestPosition(asteroid, d.asteroidBelt)
		if tmp != nil {
			e = d.asteroidBelt.InsertAfter(asteroid, tmp)
		} else {
			e = d.asteroidBelt.PushFront(asteroid)
		}
		if count > best {
			best = count
		}
	}
	return strconv.Itoa(best), errors.New("todo")
}

func (d *Day10) Part2() (string, error) {
	return "", errors.New("todo")
}
