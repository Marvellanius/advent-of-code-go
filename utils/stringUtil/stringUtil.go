package stringUtil

import "strconv"

func ToInt(input string) int {
	output, _ := strconv.Atoi(input)

	return output
}

func ToRune(input string) rune {
	return []rune(input)[0]
}
