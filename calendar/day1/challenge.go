package day1

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type Day1 struct {
	fuelRqmts []int
}

func (d *Day1) InputPath() string {
	return "/calendar/day1/input"
}

func (d *Day1) Prepare(input *os.File) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fuelRqmt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d.fuelRqmts = append(d.fuelRqmts, fuelRqmt)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Day1) Part1() (string, error) {
	fuelRqmtSum := 0
	for _, fuelRqmt := range d.fuelRqmts {
		fuelRqmtSum += (fuelRqmt / 3) - 2
	}
	return strconv.Itoa(fuelRqmtSum), nil
}

func (d *Day1) Part2() (string, error) {
	return "", errors.New("not implemented")
}
