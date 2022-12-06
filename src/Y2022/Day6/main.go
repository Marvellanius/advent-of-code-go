package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"advent-of-code-go/utils/advent"
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
	return findStart(input, 4)
}

func part2(input string) int {
	return findStart(input, 14)
}

func findStart(input string, distinctRunes int) (answer int) {
	runeInput := []rune(input)

	for i := range input {
		var window []rune
		for j := 0; j < distinctRunes; j++ {
			window = append(window, runeInput[i+j])
		}

		if len(advent.RuneIntersect(window, window)) == distinctRunes {
			answer = i + distinctRunes
			break
		}
	}

	return
}
