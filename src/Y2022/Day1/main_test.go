package main

import (
	"testing"
)

var example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	input := example
	want := 24000

	got := part1(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := example
	want := 45000

	got := part2(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
