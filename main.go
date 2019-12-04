package main

import (
	"github.com/gvassili/adventofcode2019/calendar"
	"github.com/gvassili/adventofcode2019/code_advent"
	"github.com/olekukonko/tablewriter"
	"os"
	"time"
)

type result struct {
	label    string
	result   string
	error    error
	duration time.Duration
}

func runChallenge(challenge code_advent.Challengeable) []result {
	filepath := challenge.InputPath()
	var input *os.File
	if filepath != "" {
		file, err := os.Open("./" + challenge.InputPath())
		if err != nil {
			return []result{{"prepare", "", err, 0}}
		}
		defer file.Close()
		input = file
	}
	startTs := time.Now()
	if err := challenge.Prepare(input); err != nil {
		return []result{{"prepare", "", err, time.Now().Sub(startTs)}}
	}
	prepareTs := time.Now()
	result1, err1 := challenge.Part1()
	part1Ts := time.Now()
	result2, err2 := challenge.Part2()
	part2Ts := time.Now()
	return []result{
		{"prepare", "", nil, prepareTs.Sub(startTs)},
		{"part1", result1, err1, part1Ts.Sub(prepareTs)},
		{"part2", result2, err2, part2Ts.Sub(part1Ts)},
	}
}

func main() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetAutoMergeCells(true)
	table.SetHeader([]string{"Day", "Step", "Result", "Error", "Excl time", "Incl time"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgBlackColor, tablewriter.BgCyanColor},
	)
	var challenges []calendar.DailyChallenge
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			challenge, err := calendar.LoadChallenge(arg)
			if err != nil {
				panic(err)
			}
			challenges = append(challenges, challenge)
		}
	} else {
		challenges = calendar.LoadAllChallenges()
	}
	for idx, challenge := range challenges {
		blockFgColor := tablewriter.FgHiBlackColor
		blockMod := tablewriter.Normal
		if idx&1 == 1 {
			blockFgColor = tablewriter.FgWhiteColor
			blockMod = 2
		}
		var inclTime time.Duration
		for _, row := range runChallenge(challenge.Challenge) {
			inclTime += row.duration
			fgColor := blockFgColor
			errMsg := ""
			if row.error != nil {
				fgColor = tablewriter.FgRedColor
				errMsg = row.error.Error()
			}
			table.Rich([]string{challenge.Name, row.label, row.result, errMsg, row.duration.String(), inclTime.String()}, []tablewriter.Colors{
				{blockMod, blockFgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
				{blockMod, fgColor},
			})
		}
	}
	table.Render()
}
