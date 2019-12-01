package code_advent

import "os"

type Challengeable interface {
	InputPath() string
	Prepare(input *os.File) error
	Part1() (string, error)
	Part2() (string, error)
}
