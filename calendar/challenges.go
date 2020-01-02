package calendar

import (
	"errors"
	"github.com/gvassili/adventofcode2019/calendar/day01"
	"github.com/gvassili/adventofcode2019/calendar/day02"
	"github.com/gvassili/adventofcode2019/calendar/day03"
	"github.com/gvassili/adventofcode2019/calendar/day04"
	"github.com/gvassili/adventofcode2019/calendar/day05"
	"github.com/gvassili/adventofcode2019/calendar/day06"
	"github.com/gvassili/adventofcode2019/calendar/day07"
	"github.com/gvassili/adventofcode2019/calendar/day08"
	"github.com/gvassili/adventofcode2019/calendar/day09"
	"github.com/gvassili/adventofcode2019/calendar/day10"
	"github.com/gvassili/adventofcode2019/calendar/day11"
	"github.com/gvassili/adventofcode2019/calendar/day12"
	"github.com/gvassili/adventofcode2019/calendar/day13"
	"github.com/gvassili/adventofcode2019/calendar/day14"
	"github.com/gvassili/adventofcode2019/calendar/day15"
	"github.com/gvassili/adventofcode2019/code_advent"
	"sort"
)

var challenges = map[string]func() code_advent.Challenger{
	"day01": func() code_advent.Challenger { return &day01.Day1{} },
	"day02": func() code_advent.Challenger { return &day02.Day2{} },
	"day03": func() code_advent.Challenger { return &day03.Day3{} },
	"day04": func() code_advent.Challenger { return &day04.Day4{} },
	"day05": func() code_advent.Challenger { return &day05.Day5{} },
	"day06": func() code_advent.Challenger { return &day06.Day6{} },
	"day07": func() code_advent.Challenger { return &day07.Day7{} },
	"day08": func() code_advent.Challenger { return &day08.Day8{} },
	"day09": func() code_advent.Challenger { return &day09.Day9{} },
	"day10": func() code_advent.Challenger { return &day10.Day10{} },
	"day11": func() code_advent.Challenger { return &day11.Day11{} },
	"day12": func() code_advent.Challenger { return &day12.Day12{} },
	"day13": func() code_advent.Challenger { return &day13.Day13{} },
	"day14": func() code_advent.Challenger { return &day14.Day14{} },
	"day15": func() code_advent.Challenger { return &day15.Day15{} },
}

func LoadChallenge(name string) (DailyChallenge, error) {
	loader, ok := challenges[name]
	if !ok {
		return DailyChallenge{}, errors.New("could not find challenge " + name)
	}
	return DailyChallenge{loader(), name}, nil
}

type DailyChallenge struct {
	Challenge code_advent.Challenger
	Name      string
}

func LoadAllChallenges() []DailyChallenge {
	challengeNames := make([]string, 0, len(challenges))
	for name := range challenges {
		challengeNames = append(challengeNames, name)
	}
	sort.Strings(challengeNames)
	result := make([]DailyChallenge, 0, len(challenges))
	for _, name := range challengeNames {
		result = append(result, DailyChallenge{challenges[name](), name})
	}
	return result
}
