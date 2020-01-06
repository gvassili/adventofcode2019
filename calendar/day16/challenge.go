package day16

import (
	"bufio"
	"io"
)

type Day16 struct {
	sequence []byte
}

func (d Day16) InputPath() string {
	return "calendar/day16/input"
}

func (d *Day16) Prepare(input io.Reader) error {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		d.sequence = append(d.sequence, scanner.Bytes()[0]-'0')
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

var pattern = []int{0, 1, 0, -1}

func abs(nb int) int {
	if nb < 0 {
		return nb * -1
	}
	return nb
}

func (d *Day16) Part1() (string, error) {
	prevPhase := make([]byte, len(d.sequence))
	copy(prevPhase, d.sequence)
	for phase := 0; phase < 100; phase++ {
		newPhase := make([]byte, len(d.sequence))
		for i := range prevPhase {
			sum := 0
			for n, digit := range prevPhase {
				sum += int(digit) * pattern[((n+1)/(i+1))&0b11]
			}
			newPhase[i] = byte(abs(sum) % 10)
		}
		prevPhase = newPhase
	}
	for i := range prevPhase {
		prevPhase[i] += '0'
	}
	return string(prevPhase[:8]), nil
}

func (d *Day16) Part2() (string, error) {
	return "", nil
}
