package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

	var containedAssignments int

	for _, pair := range parsed {
		regex := regexp.MustCompile(`(\d+)`).FindAllString(pair, -1)

		intAssignments := convertToIntAssignments(regex)

		if intAssignments[0] >= intAssignments[2] && intAssignments[1] <= intAssignments[3] || intAssignments[2] >= intAssignments[0] && intAssignments[3] <= intAssignments[1] {
			containedAssignments++
		}
	}
	return containedAssignments
}

func part2(input string) int {
	parsed := parseInput(input)

	var containedAssignments int

	for _, pair := range parsed {

		regex := regexp.MustCompile(`(\d+)`).FindAllString(pair, -1)
		intAssignments := convertToIntAssignments(regex)

		if intAssignments[0] >= intAssignments[2] && intAssignments[0] <= intAssignments[3] ||
			intAssignments[2] >= intAssignments[0] && intAssignments[2] <= intAssignments[1] ||
			intAssignments[1] >= intAssignments[2] && intAssignments[1] <= intAssignments[3] ||
			intAssignments[3] >= intAssignments[0] && intAssignments[3] <= intAssignments[1] {
			containedAssignments++
		}
	}
	return containedAssignments
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)
	return parsed
}

func convertToIntAssignments(pair []string) (intAssignments []int) {
	for _, item := range pair {
		i, err := strconv.Atoi(item)
		if err != nil {
			panic("Shit broke")
		}
		intAssignments = append(intAssignments, i)
	}
	return
}
