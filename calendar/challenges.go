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

func LoadChallenge(name string) (DailyChallenge, error) {
	loader, ok := challenges[name]
	if !ok {
		return DailyChallenge{}, errors.New("could not find challenge " + name)
	}
	return DailyChallenge{loader(), name}, nil
}

type DailyChallenge struct {
	Challenge code_advent.Challengeable
	Name      string
}

func LoadAllChallenges() []DailyChallenge {
	result := make([]DailyChallenge, 0, len(challenges))
	for name, loader := range challenges {
		result = append(result, DailyChallenge{loader(), name})
	}
	return result
}
