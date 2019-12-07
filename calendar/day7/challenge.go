package day7

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"runtime"
	"strconv"
	"sync"
)

type Day7 struct {
	program []int
}

func (d Day7) InputPath() string {
	return "/calendar/day7/input"
}

func (d *Day7) Prepare(input *os.File) error {
	scanner := bufio.NewScanner(input)
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

	for scanner.Scan() {
		opCode, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d.program = append(d.program, opCode)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func max(rhs int, lhs int) int {
	if rhs > lhs {
		return rhs
	}
	return lhs
}

type sequence [5]int

func permutate(seq sequence, output chan sequence) {
	output <- seq
	size := len(seq)
	p := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		p[i] = i
	}
	for i := 1; i < size; {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		seq[j], seq[i] = seq[i], seq[j]
		output <- seq
		for i = 1; p[i] == 0; i++ {
			p[i] = i
		}
	}
}

func (d *Day7) Part1() (string, error) {
	permsC := make(chan sequence, 32)
	resultC := make(chan int, 32)
	var wg sync.WaitGroup
	go func() {
		permutate(sequence{0, 1, 2, 3, 4}, permsC)
		close(permsC)
	}()
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for seq := range permsC {
				outSig := 0
				var computer intcodeComputer
				computer.init(d.program)
				for _, inSig := range seq {
					r, err := computer.exec([]int{inSig, outSig})
					if err != nil {
						panic(err)
					}
					outSig = r
				}
				resultC <- outSig
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(resultC)
	}()
	maxThrust := 0
	for r := range resultC {
		maxThrust = max(maxThrust, r)
	}
	return strconv.Itoa(maxThrust), nil
}

func (d *Day7) Part2() (string, error) {
	return "", errors.New("todo")
}
