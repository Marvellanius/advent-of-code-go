package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"

	"advent-of-code-go/utils/stringUtil"
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

const UP = 'U'
const DOWN = 'D'
const RIGHT = 'R'
const LEFT = 'L'
const TAIL = 'T'
const HEAD = 'H'
const MARKER = '#'

func part1(input string) int {
	// dirs := []rune{UP, DOWN, RIGHT, LEFT}

	tail := []location{location{0, 0}}
	head := []location{location{0, 0}}
	// initialize start location

	for _, line := range parseInput(input) {
		splitLine := strings.Split(line, " ")
		// move head AND tail a certain amount of steps in a defined direction
		tail, head = moveTailAndHead(tail, head, stringUtil.ToRune(splitLine[0]), stringUtil.ToInt(splitLine[1]))
	}

	m := map[location]bool{}
	for _, t := range tail {
		m[t] = true
	}

	return len(m)
}

func part2(input string) int {
	snake := []location{}
	for i := 0; i < 10; i++ {
		snake = append(snake, location{0, 0})
	}

	tailLocations := map[location]bool{}
	tailLocations[location{0, 0}] = true

	for _, line := range parseInput(input) {
		splitLine := strings.Split(line, " ")

		for step := 0; step < stringUtil.ToInt(splitLine[1]); step++ {
			switch stringUtil.ToRune(splitLine[0]) {
			case UP:
				snake[0].y++
			case DOWN:
				snake[0].y--
			case LEFT:
				snake[0].x--
			case RIGHT:
				snake[0].x++
			}
			for index := 1; index < len(snake); index++ {
				snake[index] = updateSegmentLocation(snake[index], snake[index-1])

				if index == len(snake)-1 {
					tailLocations[snake[index]] = true
					fmt.Println(tailLocations)
				}
			}
		}
	}
	return len(tailLocations)
}

func updateSegmentLocation(current, parent location) location {
	if moreThanOneRemoved(current.x, parent.x) || moreThanOneRemoved(current.y, parent.y) {
		if parent.x-current.x > 0 {
			current.x++
		} else if parent.x-current.x < 0 {
			current.x--
		}

		if parent.y-current.y > 0 {
			current.y++
		} else if parent.x-current.y < 0 {
			current.y--
		}
	}

	return current
}

func moveTailAndHead(tail, head []location, direction rune, steps int) ([]location, []location) {
	// DIAGONALS mean TAIL will take up space HEAD left
	// H[2,2] T[1,1] -> H[2,3] T[1,1] -> H[2,3] T[2,2]
	// H[2,2] T[1,1] -> H[3,2] T[1,1] -> H[3,2] T[2,2]

	for i := 0; i < steps; i++ {
		lastTailLocation := tail[len(tail)-1]
		lastHeadLocation := head[len(head)-1]
		newHead := lastHeadLocation
		newTail := lastTailLocation
		switch direction {
		case UP:
			newHead.y++
		case DOWN:
			newHead.y--
		case LEFT:
			newHead.x--
		case RIGHT:
			newHead.x++
		}
		// horizontal movement?
		newTail = updateSegmentLocation(newTail, newHead)
		// if moreThanOneRemoved(newTail.x, newHead.x) && newTail.y == newHead.y {
		// 	if newTail.x > newHead.x {
		// 		newTail.x = newHead.x + 1
		// 	} else {
		// 		newTail.x = newHead.x - 1
		// 	}
		// } else if moreThanOneRemoved(newTail.y, newHead.y) && newTail.x == newHead.x {
		// 	// vertical movement?
		// 	if newTail.y > newHead.y {
		// 		newTail.y = newHead.y + 1
		// 	} else {
		// 		newTail.y = newHead.y - 1
		// 	}
		// } else if (moreThanOneRemoved(newTail.y, newHead.y) && int(math.Abs(float64(newTail.x)-float64(newHead.x))) == 1) ||
		// 	(moreThanOneRemoved(newTail.x, newHead.x) && int(math.Abs(float64(newTail.y)-float64(newHead.y))) == 1) {
		// 	// diagonal movement?
		// 	newTail = lastHeadLocation
		// }

		tail = append(tail, newTail)
		head = append(head, newHead)

	}

	return tail, head
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)

	return parsed
}

func moreThanOneRemoved(source, target int) bool {
	return source-target > 1 || source-target < -1
}

type location struct {
	x int
	y int
}
