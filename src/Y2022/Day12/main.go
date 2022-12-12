package main

import (
	_ "embed"
	"flag"
	"fmt"
	"sort"
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

type location struct {
	row, col int
}

func part1(input string) int {
	return findShortestDistanceForInput(input, false)
}

func part2(input string) int {
	return findShortestDistanceForInput(input, true)
}

func findShortestDistanceForInput(input string, partTwo bool) int {
	start, end, matrix := parseInput(input)
	visited := make(map[location]bool)

	rowBound := len(matrix)
	columnBound := len(matrix[0])

	// start from the end, makes part 2 easier
	queue := []location{end}
	distances := map[location]int{end: 0}

	directions := [][]int{
		{0, 1}, {1, 0}, {-1, 0}, {0, -1},
	}

	for {
		currentLocation := queue[0]
		visited[currentLocation] = true
		queue = queue[1:]
		if currentLocation == start || (partTwo && matrix[currentLocation.row][currentLocation.col] == 'a') {
			start = currentLocation
			break
		}

		for _, neighbour := range directions {
			row, column := neighbour[1], neighbour[0]
			nextLocation := location{currentLocation.row + row, currentLocation.col + column}

			if !visited[nextLocation] && nextLocation.row >= 0 && nextLocation.col >= 0 &&
				nextLocation.col < columnBound && nextLocation.row < rowBound &&
				matrix[currentLocation.row][currentLocation.col]-matrix[nextLocation.row][nextLocation.col] <= 1 {
				if distances[nextLocation] == 0 {
					queue = append(queue, nextLocation)
					distances[nextLocation] = distances[currentLocation] + 1
				}
				if distances[nextLocation] >= distances[currentLocation]+1 {
					distances[nextLocation] = distances[currentLocation] + 1
				}
			}

		}
		sort.Slice(queue, func(a, b int) bool {
			return distances[queue[a]] < distances[queue[b]]
		})
	}

	return distances[start]
}

func parseInput(input string) (start, end location, matrix [][]rune) {
	var parsed []string
	parsed = append(parsed, strings.Split(input, "\n")...)

	matrix = make([][]rune, 0)
	// make it runes, consecutive letters correspond to consecutive ints
	for _, line := range parsed {
		var row []rune
		for i, char := range line {
			if char == 'S' {
				start = location{col: i, row: len(matrix)}
				char = 'a'
			}
			if char == 'E' {
				end = location{col: i, row: len(matrix)}
				char = 'z'
			}
			row = append(row, char)
		}
		matrix = append(matrix, row)
	}

	return
}
