package calendar

import (
	"errors"
	"github.com/gvassili/adventofcode2019/calendar/day1"
	"github.com/gvassili/adventofcode2019/calendar/day2"
	"github.com/gvassili/adventofcode2019/code_advent"
)

var challenges = map[string]func() code_advent.Challengeable{
	"day1": func() code_advent.Challengeable { return &day1.Day1{} },
	"day2": func() code_advent.Challengeable { return &day2.Day2{} },
}

func LoadChallenge(name string) (code_advent.Challengeable, error) {
	loader, ok := challenges[name]
	if !ok {
		return nil, errors.New("could not find challenge " + name)
	}
	return loader(), nil
}
