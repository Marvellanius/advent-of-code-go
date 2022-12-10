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
	snake := make([]location, 10)

	tailLocations := map[location]bool{}
	tailLocations[snake[9]] = true

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
			for index := range snake[:len(snake)-1] {
				snake[index+1] = updateSegmentLocation(snake[index+1], snake[index])
			}
			tailLocations[snake[9]] = true
		}
	}

	return len(tailLocations)
}

func updateSegmentLocation(current, parent location) (newSegment location) {
	newSegment = current
	/*
		all possible configurations for (at least) vertical movement, convert all Hs to locations to switch on
		HHHHH
		H...H
		..T..
		H...H
		HHHHH

		all possible configurations for (at least) horizontal movement, convert all Hs to locations to switch on
		HH.HH
		H...H
		H.T.H
		H...H
		HH.HH
	*/

	switch (location{parent.x - current.x, parent.y - current.y}) {
	case location{-2, 1}, location{-1, 2}, location{0, 2}, location{1, 2}, location{2, 1}, location{2, 2}, location{-2, 2}:
		newSegment.y++
	}
	switch (location{parent.x - current.x, parent.y - current.y}) {
	case location{2, -1}, location{1, -2}, location{2, 0}, location{2, 1}, location{1, 2}, location{2, 2}, location{2, -2}:
		newSegment.x++
	}
	switch (location{parent.x - current.x, parent.y - current.y}) {
	case location{-2, -2}, location{2, -1}, location{1, -2}, location{0, -2}, location{-1, -2}, location{-2, -1}, location{2, -2}:
		newSegment.y--
	}
	switch (location{parent.x - current.x, parent.y - current.y}) {
	case location{-2, -2}, location{-1, -2}, location{-2, -1}, location{-2, 0}, location{-2, 1}, location{-1, 2}, location{-2, 2}:
		newSegment.x--
	}

	return newSegment
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
		newTail = updateSegmentLocation(newTail, newHead)

		tail = append(tail, newTail)
		head = append(head, newHead)

	}

	return tail, head
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)

	return parsed
}

type location struct {
	x int
	y int
}
