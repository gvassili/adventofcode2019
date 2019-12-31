package day14

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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

func (d *Day14) buildFuel(amount int) int {
	storage := make(map[string]int)
	var buildChemical func(chemical *chemical, amount int)
	buildChemical = func(chemical *chemical, amount int) {
		count := amount / chemical.produce
		if amount%chemical.produce != 0 {
			count++
		}
		for _, ingredient := range chemical.recipe {
			storage[ingredient.chemical.name] -= ingredient.count * count
			if ingredient.chemical.name != "ORE" {
				if storage[ingredient.chemical.name] < 0 {
					buildChemical(ingredient.chemical, storage[ingredient.chemical.name]*-1)
				}
			}
		}
		storage[chemical.name] += chemical.produce * count
	}
	buildChemical(d.chemicals["FUEL"], amount)
	return storage["ORE"] * -1
}

func (d *Day14) Part1() (string, error) {
	return strconv.Itoa(d.buildFuel(1)), nil
}

func (d *Day14) Part2() (string, error) {
	return strconv.Itoa(sort.Search(10_000_000, func(i int) bool {
		return d.buildFuel(i) >= 1_000_000_000_000
	}) - 1), nil
}
