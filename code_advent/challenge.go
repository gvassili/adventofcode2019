package code_advent

import (
	"io"
)

type Challenger interface {
	InputPath() string
	Prepare(input io.Reader) error
	Part1() (string, error)
	Part2() (string, error)
}
