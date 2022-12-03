package main

import (
	"testing"
)

var example = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestPart1(t *testing.T) {
	input := example
	want := 157

	got := part1(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := example
	want := 70

	got := part2(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
