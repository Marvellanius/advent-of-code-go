package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  7,
		},
		{
			name:  "example2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		{
			name:  "example3",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
		},
		{
			name:  "example4",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
		},
		{
			name:  "example5",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
		{
			name:  "actual",
			input: input,
			want:  1262,
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
			name:  "example1",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
		{
			name:  "example2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  23,
		},
		{
			name:  "example3",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  23,
		},
		{
			name:  "example4",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  29,
		},
		{
			name:  "example5",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  26,
		},
		{
			name:  "actual",
			input: input,
			want:  3444,
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
