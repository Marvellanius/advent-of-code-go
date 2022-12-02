package main

import (
	_ "embed"
	"fmt"
	"flag"
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
	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (parsed []string) {
	for _, line := range strings.Split(input, "\n") {
		parsed = append(parsed, line)
	}

	return parsed
}
