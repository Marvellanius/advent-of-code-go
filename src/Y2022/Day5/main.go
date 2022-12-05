package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

func part1(input string) string {
	crates, instructions := parseInput(input)

	for _, line := range instructions {
		instruction := parseInstruction(regexp.MustCompile(`\d+`).FindAllString(line, -1))
		crates = moveCrates(crates, instruction[0], instruction[1], instruction[2], false)
	}

	var cratesString string
	for i := range crates {
		cratesString += string(crates[i][len(crates[i])-1])
	}
	return cratesString
}

func part2(input string) string {
	crates, instructions := parseInput(input)

	for _, line := range instructions {
		instruction := parseInstruction(regexp.MustCompile(`\d+`).FindAllString(line, -1))
		crates = moveCrates(crates, instruction[0], instruction[1], instruction[2], true)
	}

	var cratesString string
	for i := range crates {
		cratesString += string(crates[i][len(crates[i])-1])
	}
	return cratesString
}

func parseInput(input string) (crates [][]rune, instructions []string) {

	regex := regexp.MustCompile(`\n\n`).Split(input, -1)

	crates = parseCrates(regex[0])
	instructions = append(instructions, strings.Split(regex[1], "\n")...)

	return
}

func parseCrates(input string) (crates [][]rune) {
	unparsedRows := strings.Split(input, "\n")
	rowLength := len(unparsedRows[0])

	numCols := int(math.Ceil(float64(rowLength) / 4))

	var rows [][]rune
	var cols []rune = nil
	for _, row := range unparsedRows {
		if row[1] == '1' {
			continue
		}
		cols = make([]rune, numCols)
		for j, col := range row {
			if col != ' ' && col != '[' && col != ']' {
				colIndex := int(float64(j) / 4)
				cols[colIndex] = col
			}
		}
		rows = append(rows, cols)
	}
	crates = make([][]rune, numCols)

	for i := range crates {
		crates[i] = make([]rune, len(rows))
	}

	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			crates[j][i] = rows[len(rows)-(1+i)][j]
		}
	}
	return
}

func moveCrates(input [][]rune, amount, source, destination int, multi bool) (crates [][]rune) {
	sourceCol := source - 1
	destinationCol := destination - 1
	var movingCrate rune
	var movingCrates []rune

	for i := 0; i < amount; i++ {
		for ok := true; ok; ok = (movingCrate == 0) {
			movingCrate, input[sourceCol] = input[sourceCol][len(input[sourceCol])-1], input[sourceCol][:len(input[sourceCol])-1]
			if movingCrate != 0 {
				movingCrates = append([]rune{movingCrate}, movingCrates...)
			}
		}
		if !multi {
			input[destinationCol] = append(input[destinationCol], movingCrate)
		}
	}
	if multi {
		input[destinationCol] = append(input[destinationCol], movingCrates...)
	}

	crates = append(crates, input...)

	return
}

func parseInstruction(input []string) (instructions []int) {
	for _, item := range input {
		i, err := strconv.Atoi(item)
		if err != nil {
			panic("Shit broke")
		}
		instructions = append(instructions, i)
	}
	return
}
