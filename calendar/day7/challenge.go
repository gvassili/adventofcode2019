package day7

import (
	"bufio"
	"bytes"
	"os"
	"runtime"
	"strconv"
	"sync"
)

type Day7 struct {
	program []int
}

func (d Day7) InputPath() string {
	return "calendar/day7/input"
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
	permsC := make(chan sequence, 1)
	resultC := make(chan int, 1)
	var wg sync.WaitGroup
	go func() {
		permutate(sequence{0, 1, 2, 3, 4}, permsC)
		close(permsC)
	}()
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for seq := range permsC {
				var computer intcodeComputer
				computer.init(d.program)
				ioC := make([]chan int, len(seq))
				for idx, sig := range seq {
					c := make(chan int, 2)
					c <- sig
					ioC[idx] = c
				}
				ioC[0] <- 0
				for idx := range seq {
					err := computer.exec(ioC[idx], ioC[(idx+1)%len(seq)])
					if err != nil {
						panic(err)
					}
				}
				lastSig := <-ioC[0]
				resultC <- lastSig
			}
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
	permsC := make(chan sequence, 1)
	resultC := make(chan int, 1)
	var wg sync.WaitGroup
	go func() {
		permutate(sequence{5, 6, 7, 8, 9}, permsC)
		close(permsC)
	}()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for seq := range permsC {
				ioC := make([]chan int, len(seq))
				for idx, sig := range seq {
					c := make(chan int, 16)
					c <- sig
					ioC[idx] = c
				}
				ioC[0] <- 0
				var wg sync.WaitGroup
				for idx := range seq {
					wg.Add(1)
					go func(pid int) {
						var computer intcodeComputer
						computer.init(d.program)
						defer wg.Done()
						err := computer.exec(ioC[pid], ioC[(pid+1)%len(seq)])
						if err != nil {
							panic(err)
						}
					}(idx)
				}
				wg.Wait()
				lastSig := <-ioC[0]
				resultC <- lastSig
			}
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
