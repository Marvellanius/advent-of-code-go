package main

import (
	_ "embed"
	"fmt"
	"flag"
	"strings"
	"strconv"
	"sort"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int

	flag.IntVar(&part, "part", 1, "1 || 2")
	flag.Parse()
	fmt.Println("Running part: ", part)

	if part == 1 {
		answer := part1(input)
		fmt.Println("Answer:", answer)
	} else {
		answer := part2(input)
		fmt.Println("Answer:", answer)
	}
}


func part1(input string) int {
	parsed := parseInput(input)

	highestValue := 0

	var currentGoblinValue int
	for i, line := range parsed {
		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error during conversion")
				return 0
			}
			currentGoblinValue = currentGoblinValue + value
		}
		if i < len(parsed) - 1 && parsed[i+1] == "" && currentGoblinValue > highestValue {
			highestValue = currentGoblinValue
		}
		if line == "" {
			currentGoblinValue = 0
		}
	}

	return highestValue
}

func part2(input string) int {
	goblinValues := aggregateGoblins(input)

	// Sort goblinValues DESCENDING
	sort.Slice(goblinValues, func(i, j int) bool {
		return goblinValues[i] > goblinValues[j]
	})

	// Sum the first 3 values and save total
	var topThreeGoblins int
	for _, goblin := range goblinValues[0:3] {
		topThreeGoblins += goblin
	}
	return topThreeGoblins
}

func parseInput(input string) (parsed []string) {
	for _, line := range strings.Split(input, "\n") {
		parsed = append(parsed, line)
	}

	return parsed
}

func aggregateGoblins(input string) []int {
	parsed := parseInput(input)

	var currentGoblinValue int
	var goblinSlice []int
	for i, line := range parsed {
		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				panic("Error during conversion")
			}
			currentGoblinValue = currentGoblinValue + value
		}
		if i == len(parsed) - 1 || (i < len(parsed) - 1 && parsed[i+1] == "") {
			goblinSlice = append(goblinSlice, currentGoblinValue)
		}
		if line == "" {
			currentGoblinValue = 0
		}
	}

	return goblinSlice
}