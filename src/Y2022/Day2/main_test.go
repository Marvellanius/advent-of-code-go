package main

import (
	"testing"
)

var example = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  15,
		},
		{
			name:  "actual",
			input: input,
			want:  9651,
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
			want:  12,
		},
		{
			name:  "actual",
			input: input,
			want:  10560,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := part2(test.input); got != test.want {
				t.Errorf("part2() = %d, want %d", got, test.want)
			}
		})
	}
	input := example
	want := 15

	got := part1(input)
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func Test_convertToPoints(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "A",
			input: "A",
			want:  1,
		},
		{
			name:  "B",
			input: "B",
			want:  2,
		},
		{
			name:  "C",
			input: "C",
			want:  3,
		},
		{
			name:  "X",
			input: "X",
			want:  1,
		},
		{
			name:  "Y",
			input: "Y",
			want:  2,
		},
		{
			name:  "Z",
			input: "Z",
			want:  3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := convertToPoints(test.input); got != test.want {
				t.Errorf("convertToPoints(%s) = %d, want %d", test.input, got, test.want)
			}
		})
	}
}

func Test_determineWin(t *testing.T) {
	tests := []struct {
		input gameRound
		want  int
	}{
		{
			input: gameRound{1, 1, 0},
			want:  3,
		},
		{
			input: gameRound{2, 2, 0},
			want:  3,
		},
		{
			input: gameRound{3, 3, 0},
			want:  3,
		},
		{
			input: gameRound{1, 2, 0},
			want:  6,
		},
		{
			input: gameRound{2, 3, 0},
			want:  6,
		},
		{
			input: gameRound{1, 3, 0},
			want:  0,
		},
		{
			input: gameRound{3, 1, 0},
			want:  6,
		},
		{
			input: gameRound{2, 1, 0},
			want:  0,
		},
		{
			input: gameRound{3, 2, 0},
			want:  0,
		},
	}

	for _, test := range tests {
		t.Run("example", func(t *testing.T) {
			if got := determineWin(test.input).outcome; got != test.want {
				t.Errorf("determineWin(%d\n) = %d, want %d", test.input, got, test.want)
			}
		})
	}
}
