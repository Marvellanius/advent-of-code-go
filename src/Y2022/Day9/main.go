package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
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

	return len(tail)
}

func part2(input string) int {
	return 0
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
		if moreThanOneRemoved(newTail.x, newHead.x) && newTail.y == newHead.y {
			newTail.x = newHead.x - 1
		} else if moreThanOneRemoved(newTail.y, newHead.y) && newTail.x == newHead.x {
			// vertical movement?
			newTail.y = newHead.y - 1
		} else if (moreThanOneRemoved(newTail.y, newHead.y) && int(math.Abs(float64(newTail.x)-float64(newHead.x))) == 1) ||
			(moreThanOneRemoved(newTail.x, newHead.x) && int(math.Abs(float64(newTail.y)-float64(newHead.y))) == 1) {
			// diagonal movement?
			newTail = lastHeadLocation
		}

		if lastTailLocation != newTail {
			tail = append(tail, newTail)
		}
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
