package main

import (
	"testing"
)

var example = ``

func TestPart1(t *testing.T) {
	input := example
	want := 0

	got := part1(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := example
	want := 0

	got := part2(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
