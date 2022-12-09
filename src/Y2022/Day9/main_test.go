package main

import (
	"testing"
)

var example = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  13,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  483,
		// },
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
			want:  13,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  483,
		// },
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := part2(test.input); got != test.want {
				t.Errorf("part2() = %d, want %d", got, test.want)
			}
		})
	}
}
