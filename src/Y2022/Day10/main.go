package main

import (
	"advent-of-code-go/utils/stringUtil"
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
	// find sigstrength on 20, 60, 100, 140, 180, 220-th cycles
	// sum them
	x := 1
	cycle := 1
	var signalStrengths []int
	for _, line := range parseInput(input) {
		operation := strings.Split(line, " ")[0]
		switch operation {
		case "addx":
			signalStrengths, cycle = addCycles(x, signalStrengths, cycle, 2)
			x += stringUtil.ToInt(strings.Split(line, " ")[1])
		case "noop":
			signalStrengths, cycle = addCycles(x, signalStrengths, cycle, 1)
		}
		if cycle > 220 {
			break
		}
	}

	var sum int
	for _, signal := range signalStrengths {
		sum += signal
	}

	return sum
}

func addCycles(xVal int, signalStrengths []int, currentCycle int, cycles int) ([]int, int) {
	var i int
	for i = currentCycle; i < currentCycle+cycles; i++ {
		if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
			signalStrengths = append(signalStrengths, xVal*i)
		}
	}
	return signalStrengths, i
}

func part2(input string) string {
	// crt = 40 wide by 6 high
	crt := make([][]string, 6)
	for rowIndex := range crt {
		crt[rowIndex] = make([]string, 40)
		for colIndex := range crt[rowIndex] {
			crt[rowIndex][colIndex] = "."
		}
	}

	registerX := 1
	// sprite positions surround registerX position
	cycle := 1
	for _, line := range parseInput(input) {
		operation := strings.Split(line, " ")[0]
		switch operation {
		case "addx":
			crt, cycle = drawPixel(crt, registerX, cycle, 2)
			registerX += stringUtil.ToInt(strings.Split(line, " ")[1])
		case "noop":
			crt, cycle = drawPixel(crt, registerX, cycle, 1)
		}
		if cycle > 240 {
			break
		}
	}

	var crtString string
	for rowIndex := range crt {
		fmt.Println(crt[rowIndex])
		crtString += strings.Join(crt[rowIndex], "")
		crtString += "\n"
	}

	return crtString
}

func drawPixel(input [][]string, registerX int, currentCycle int, cycles int) ([][]string, int) {
	crt := input
	var i int
	spritePosition := []int{registerX - 1, registerX, registerX + 1}

	// fmt.Println(spritePosition)
	for i = currentCycle; i < currentCycle+cycles; i++ {
		rowIndex := int(i / 40)
		if i%40 == 0 {
			rowIndex--
		}
		colIndex := ((i - 1) % 40)
		for _, position := range spritePosition {
			if colIndex == position {
				crt[rowIndex][colIndex] = "#"
			}
		}
	}
	return crt, i
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)

	return parsed
}
