package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"

	"advent-of-code-go/utils/slice"
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
	matrix := createTreeMatrix(input)

	amountOfTrees := 0
	for i := range matrix {
		if i == 0 || i == len(matrix)-1 {
			amountOfTrees += len(matrix[i])
			continue
		}
		for j := range matrix[i] {
			if j == 0 || j == len(matrix[i])-1 {
				amountOfTrees++
				continue
			}

			// check cells around (col, row)
			// check last row first
			visibleFrom := 0

			var rowsBottom []int
			for _, row := range matrix[i+1:] {
				rowsBottom = append(rowsBottom, row[j])
			}

			var rowsTop []int
			for _, row := range matrix[:i] {
				rowsTop = append(rowsTop, row[j])
			}

			if slice.Max(rowsBottom) < matrix[i][j] {
				visibleFrom++
			}

			if slice.Max(rowsTop) < matrix[i][j] {
				visibleFrom++
			}

			if slice.Max(matrix[i][:j]) < matrix[i][j] {
				visibleFrom++
			}

			if slice.Max(matrix[i][j+1:]) < matrix[i][j] {
				visibleFrom++
			}

			if visibleFrom > 0 {
				amountOfTrees++
			}
		}
	}

	return amountOfTrees
}

func part2(input string) int {
	matrix := createTreeMatrix(input)

	maxScenicScore := 0
	for i := range matrix {
		for j := range matrix[i] {

			scenicScore := 0
			// check cells around (col, row)
			var treesBottom []int
			for _, row := range matrix[i+1:] {
				treesBottom = append(treesBottom, row[j])
			}

			var treesTop []int
			for _, row := range matrix[:i] {
				treesTop = slice.PrependIntSlice(treesTop, row[j])
			}

			var treesLeft []int
			for _, tree := range matrix[i][:j] {
				treesLeft = slice.PrependIntSlice(treesLeft, tree)
			}

			t := calculateVisibleTrees(matrix[i][j], treesTop)
			l := calculateVisibleTrees(matrix[i][j], treesLeft)
			b := calculateVisibleTrees(matrix[i][j], treesBottom)
			r := calculateVisibleTrees(matrix[i][j], matrix[i][j+1:])

			scenicScore = b * t * l * r

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)

	return parsed
}

func calculateVisibleTrees(source int, targets []int) (visibleTrees int) {
	for _, trees := range targets {
		if trees < source {
			visibleTrees++
		}
		if trees >= source {
			visibleTrees++
			break
		}
	}
	return
}

func createTreeMatrix(input string) (matrix [][]int) {
	for _, line := range parseInput(input) {
		var row []int
		for _, c := range line {
			tree, _ := strconv.Atoi(string(c))
			row = append(row, tree)
		}
		matrix = append(matrix, row)
	}
	return
}
