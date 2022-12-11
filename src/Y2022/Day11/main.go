package main

import (
	"advent-of-code-go/utils/slice"
	"advent-of-code-go/utils/stringUtil"
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

func part1(input string) int {
	// Monkeys []monkey have Items []int, an operation string, a test {int, operation_if_true, operation_if_false}
	monkeys := []monkey{}

	monkeyStrings := regexp.MustCompile(`\n\n`).Split(input, -1)

	for _, line := range monkeyStrings {
		separatedMonkeys := regexp.MustCompile(`\r?\n`).Split(line, -1)

		currentMonkey := monkey{}
		currentMonkey.test = test{}
		for index, monkeyLine := range separatedMonkeys {
			switch index {
			case 1:
				currentMonkey.items = addStartingItems(separatedMonkeys[index])
			case 2:
				currentMonkey.operation = addOperation(monkeyLine)
			case 3:
				currentMonkey.test.division = stringUtil.ToInt(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1)[0])
			case 4:
				currentMonkey.test.operationTrue = stringUtil.ToInt(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1)[0])
			case 5:
				currentMonkey.test.operationFalse = stringUtil.ToInt(regexp.MustCompile(`\d+`).FindAllString(monkeyLine, -1)[0])
			}

		}
		monkeys = append(monkeys, currentMonkey)
	}

	for i := 0; i < 20; i++ {
		monkeys = playRound(monkeys)
	}

	var activity []int
	for _, monkey := range monkeys {
		activity = append(activity, monkey.activity)
	}

	sort.Ints(activity)

	return activity[len(activity)-1] * activity[len(activity)-2]
}

func part2(input string) int {
	return 0
}

type test struct {
	division       int
	operationTrue  int
	operationFalse int
}

type monkey struct {
	items     []int
	operation string
	test      test
	activity  int
}

func playRound(monkeys []monkey) []monkey {
	for index, monkey := range monkeys {
		initialLength := len(monkey.items)
		monkeys[index].activity += initialLength
		for itemIndex := 0; itemIndex < initialLength; itemIndex++ {
			var item int
			item, monkeys[index].items = slice.PopFirstElement(monkeys[index].items)
			worry := doOperation(item, monkey.operation)
			item = worry / 3

			var throwIndex int
			if item%monkey.test.division == 0 {
				throwIndex = monkey.test.operationTrue
			} else {
				throwIndex = monkey.test.operationFalse
			}

			monkeys[throwIndex].items = append(monkeys[throwIndex].items, item)
		}
	}

	return monkeys
}

func doOperation(old int, operation string) (new int) {
	parsed := strings.Replace(operation, "new = old ", "", -1)

	operator := parsed[0]
	value := strings.Replace(string(parsed[1:]), " ", "", -1)
	fmt.Println(operator)
	switch operator {
	case '*':
		if value == "old" {
			new = old * old
		} else {
			new = old * stringUtil.ToInt(value)
		}
	default:
		if value == "old" {
			new = old + old
		} else {
			new = old + stringUtil.ToInt(value)
		}
	}

	return
}

func addOperation(input string) (output string) {
	output = strings.Replace(input, "  Operation: ", "", -1)
	return
}

func addStartingItems(input string) (output []int) {
	stringItems := regexp.MustCompile(`\d+`).FindAllString(input, -1)
	for _, item := range stringItems {
		output = append(output, stringUtil.ToInt(item))
	}

	return
}

func parseInput(input string) (parsed []string) {
	parsed = append(parsed, strings.Split(input, "\n")...)

	return parsed
}
