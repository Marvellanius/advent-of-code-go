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
	// 1pt Rock A||X
	// 2pt Paper B||Y
	// 3pt Scissors C||Z
	// 0pt loss, 3pt draw, 6pt win

	return determineTotalScore(runGame(parseInput(input, 1)))
}

func part2(input string) int {
	// 1pt Rock A
	// 2pt Paper B
	// 3pt Scissors C
	// X -> lose
	// Y -> draw
	// Z -> win
	// 0pt loss, 3pt draw, 6pt win
	return determineTotalScore(runGame(parseInput(input, 2)))
}

func parseInput(input string, part int) (game []gameRound) {
	var currentRound gameRound
	for _, line := range strings.Split(input, "\n") {
		round := strings.Split(line, " ")
		currentRound.opponent = convertToPoints(round[0])
		if part == 1 {
			currentRound.player = convertToPoints(round[1])
		} else {
			currentRound.player = determinePlay(round[0], round[1])
		}

		game = append(game, currentRound)
	}

	return game
}

func determineTotalScore(game []gameRound) (points int) {
	for _, round := range game {
		points += (round.player + round.outcome)
	}
	return
}

// Need to remember moves by both players
func runGame(game []gameRound) (solvedGames []gameRound) {
	for _, round := range game {
		solvedGames = append(solvedGames, determineWin(round))
	}
	return
}

func determineWin(round gameRound) (solvedRound gameRound) {
	solvedRound = round

	// 1 = Rock
	// 2 = Paper
	// 3 = Scissor
	switch true {
	case round.opponent == 1 && round.player == 3:
		solvedRound.outcome = 0
	case round.player == 1 && round.opponent == 3:
		solvedRound.outcome = 6
	case round.opponent > round.player:
		solvedRound.outcome = 0
	case round.opponent == round.player:
		solvedRound.outcome = 3
	case round.opponent < round.player:
		solvedRound.outcome = 6
	}

	return
}

func determinePlay(opponent string, strategy string) (points int) {
	switch strategy {
	case "X":
		switch opponent {
		case "A":
			points = 3
		case "B":
			points = 1
		case "C":
			points = 2
		}
	case "Y":
		switch opponent {
		case "A":
			points = 1
		case "B":
			points = 2
		case "C":
			points = 3
		}
	case "Z":
		switch opponent {
		case "A":
			points = 2
		case "B":
			points = 3
		case "C":
			points = 1
		}
	}
	return
}

func convertToPoints(input string) (points int) {
	switch input {
	case "A", "X":
		points = 1
	case "B", "Y":
		points = 2
	case "C", "Z":
		points = 3
	}
	return points
}

type gameRound struct {
	opponent int
	player   int
	outcome  int
}
