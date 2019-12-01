package calendar

import (
	"errors"
	"github.com/gvassili/adventofcode2019/code_advent"
)

var challenges = map[string]func() code_advent.Challengeable{}

func LoadChallenge(name string) (code_advent.Challengeable, error) {
	loader, ok := challenges[name]
	if !ok {
		return nil, errors.New("could not find challenge " + name)
	}
	return loader(), nil
}
