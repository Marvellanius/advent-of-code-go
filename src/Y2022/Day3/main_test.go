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
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  157,
		},
		{
			name:  "actual",
			input: input,
			want:  7848,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := part1(test.input); got != test.want {
				t.Errorf("part1() = %d, want %d", got, test.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  70,
		},
		{
			name:  "actual",
			input: input,
			want:  2616,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := part2(test.input); got != test.want {
				t.Errorf("part2() = %d, want %d", got, test.want)
			}
		})
	}
}
