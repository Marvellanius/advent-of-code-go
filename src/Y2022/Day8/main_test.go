package main

import (
	"testing"
)

var example = `30373
25512
65332
33549
35390`

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  21,
		},
		{
			name:  "actual",
			input: input,
			want:  1693,
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
			want:  8,
		},
		{
			name:  "actual",
			input: input,
			want:  422059,
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

func Test_calculateVisibleTrees(t *testing.T) {
	tests := []struct {
		name   string
		source int
		trees  []int
		want   int
	}{
		{
			name:   "example",
			source: 5,
			trees:  []int{6},
			want:   1,
		},
		{
			name:   "example",
			source: 5,
			trees:  []int{5, 0},
			want:   1,
		},
		{
			name:   "example",
			source: 5,
			trees:  []int{3, 3, 2},
			want:   3,
		},
		{
			name:   "example",
			source: 5,
			trees:  []int{3, 5},
			want:   2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := calculateVisibleTrees(test.source, test.trees); got != test.want {
				t.Errorf("calculateVisibleTrees() = %d, want %d", got, test.want)
			}
		})
	}
}
