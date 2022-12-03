package main

import (
	_ "embed"
	"flag"
	"fmt"
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

	var prioritySum int
	for _, line := range parsed {
		backpackOne := []rune(line[:len(line)/2])
		backpackTwo := []rune(line[len(line)/2:])

		prioritySum += getPriorityForType(determineDuplicateType(backpackOne, backpackTwo))
	}
	return prioritySum
}

func part2(input string) int {
	parsed := parseInput(input)

	var prioritySum int
	for i := 0; i < len(parsed); i += 3 {
		var groupRucksacks [][]rune
		for _, rucksack := range parsed[i : i+3] {
			groupRucksacks = append(groupRucksacks, []rune(rucksack))
		}
		prioritySum += getPriorityForType(determineBadgeItem(groupRucksacks))
	}

	return prioritySum
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)
	return parsed
}

func determineDuplicateType(backpackOne, backpackTwo []rune) rune {

	out := intersect(backpackOne, backpackTwo)

	// there really only is one useful result according to the assignment
	return out[0]
}

func determineBadgeItem(rucksacks [][]rune) rune {
	out := intersect(rucksacks[2], intersect(rucksacks[0], rucksacks[1]))

	return out[0]
}

func intersect(first, second []rune) []rune {
	out := []rune{}
	bucket := map[rune]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
}

func getPriorityForType(itemType rune) int {

	m := map[rune]int{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
		'A': 27,
		'B': 28,
		'C': 29,
		'D': 30,
		'E': 31,
		'F': 32,
		'G': 33,
		'H': 34,
		'I': 35,
		'J': 36,
		'K': 37,
		'L': 38,
		'M': 39,
		'N': 40,
		'O': 41,
		'P': 42,
		'Q': 43,
		'R': 44,
		'S': 45,
		'T': 46,
		'U': 47,
		'V': 48,
		'W': 49,
		'X': 50,
		'Y': 51,
		'Z': 52,
	}

	return m[itemType]
}
