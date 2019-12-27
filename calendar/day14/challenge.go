package day14

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type ingredient struct {
	chemical *chemical
	count    int
}

type chemical struct {
	name    string
	recipe  []ingredient
	produce int
}

type Day14 struct {
	chemicals map[string]*chemical
}

func (d Day14) InputPath() string {
	return "calendar/day14/input"
}

func (d *Day14) Prepare(input io.Reader) error {
	d.chemicals = make(map[string]*chemical)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		reaction := strings.SplitN(scanner.Text(), "=>", 2)
		chems := strings.Split(reaction[0], ",")
		var name string
		var produce int
		fmt.Sscanf(reaction[1], "%d%s", &produce, &name)
		output, ok := d.chemicals[name]
		if !ok {
			output = &chemical{name: name}
			d.chemicals[name] = output
		}
		output.produce = produce
		for _, def := range chems {
			fmt.Sscanf(def, "%d%s", &produce, &name)
			chem, ok := d.chemicals[name]
			if !ok {
				chem = &chemical{name: name}
				d.chemicals[name] = chem
			}
			output.recipe = append(output.recipe, ingredient{chem, produce})
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (d *Day14) Part1() (string, error) {
	storage := make(map[string]int)
	var buildChemical func(chemical *chemical, indent int)
	buildChemical = func(chemical *chemical, indent int) {
		for _, ingredient := range chemical.recipe {
			if ingredient.chemical.name != "ORE" {
				for storage[ingredient.chemical.name] < ingredient.count {
					buildChemical(ingredient.chemical, indent+1)
				}
			}
			storage[ingredient.chemical.name] -= ingredient.count
		}
		storage[chemical.name] += chemical.produce
	}
	buildChemical(d.chemicals["FUEL"], 0)
	return strconv.Itoa(storage["ORE"] * -1), nil
}

func (d *Day14) Part2() (string, error) {
	return "", errors.New("todo")
}
