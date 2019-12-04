package day4

import (
	"errors"
	"os"
	"strconv"
)

type Day4 struct {
	minRange password
	maxRange password
}

func (d Day4) InputPath() string {
	return ""
}

type password []int8

func (d *Day4) Prepare(input *os.File) error {
	d.minRange = password{1, 6, 8, 8, 8, 8}
	d.maxRange = password{7, 1, 8, 0, 9, 8}
	return nil
}

func (p password) String() string {
	output := ""
	for _, b := range p {
		output += string(b + '0')
	}
	return output
}

func cmpPassword(lhs password, rhs password) int8 {
	for i := 0; i < len(lhs); i++ {
		diff := lhs[i] - rhs[i]
		if diff != 0 {
			return diff
		}
	}
	return 0
}

func recinc(password password, similarity int, i int) int {
	password[i]++
	if password[i] > 9 {
		similarity = recinc(password, similarity, i-1)
		password[i] = password[i-1]
	}
	if i > 0 {
		flag := 0
		if password[i] == password[i-1] {
			flag = 1
		}
		similarity ^= (-flag ^ similarity) & (1 << i)
	}
	return similarity
}

func computeSimilarity(password password) int {
	similarity := 0
	for i := len(password) - 1; i > 0; i-- {
		if password[i] == password[i-1] {
			similarity |= 1 << i
		}
	}
	return similarity
}

func (d *Day4) Part1() (string, error) {
	password := make(password, len(d.minRange))
	count := 0
	copy(password, d.minRange)
	similarity := computeSimilarity(password)
	if similarity != 0 {
		count += 1
	}
	for {
		similarity = recinc(password, similarity, len(password)-1)
		if cmpPassword(password, d.maxRange) > 0 {
			break
		}
		if similarity == 0 {
		} else {
			count += 1
		}
	}
	return strconv.Itoa(count), nil
}

func (d *Day4) Part2() (string, error) {
	return "", errors.New("todo")
}
