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
		{
			name:  "actual",
			input: input,
			want:  6175,
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
			want:  36,
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

func Test_updateSegmentLocation(t *testing.T) {
	tests := []struct {
		name    string
		current location
		parent  location
		want    location
	}{
		{
			name:    "diagonalExample",
			current: location{3, 0},
			parent:  location{4, 2},
			want:    location{4, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := updateSegmentLocation(test.current, test.parent); got.x != test.want.x || got.y != test.want.y {
				t.Errorf("updateSegmentLocation() got x = %d, y = %d want x = %d, y = %d", got.x, got.y, test.want.x, test.want.y)
			}
		})
	}
}

func Test_moveTailAndHead(t *testing.T) {
	tests := []struct {
		name      string
		head      []location
		tail      []location
		direction rune
		steps     int
		tailWant  []location
		headWant  []location
	}{
		{
			name:      "example",
			tail:      []location{{0, 0}},
			head:      []location{{0, 0}},
			direction: RIGHT,
			steps:     4,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}},
			direction: UP,
			steps:     4,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}},
			direction: LEFT,
			steps:     3,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}},
			direction: DOWN,
			steps:     1,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}},
			direction: RIGHT,
			steps:     4,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}},
			direction: DOWN,
			steps:     1,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}},
			direction: LEFT,
			steps:     5,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}, {3, 2}, {2, 2}, {1, 2}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {4, 2}, {3, 2}, {2, 2}, {1, 2}, {0, 2}},
		},
		{
			name:      "example",
			tail:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}, {3, 2}, {2, 2}, {1, 2}},
			head:      []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {4, 2}, {3, 2}, {2, 2}, {1, 2}, {0, 2}},
			direction: RIGHT,
			steps:     2,
			tailWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 1}, {4, 2}, {4, 3}, {3, 4}, {2, 4}, {3, 3}, {4, 3}, {3, 2}, {2, 2}, {1, 2}},
			headWant:  []location{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4}, {3, 4}, {2, 4}, {1, 4}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}, {4, 2}, {3, 2}, {2, 2}, {1, 2}, {0, 2}, {1, 2}, {2, 2}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if gotTail, gotHead := moveTailAndHead(test.tail, test.head, test.direction, test.steps); len(gotTail) != len(test.tailWant) && len(gotHead) != len(test.headWant) {
				t.Errorf("moveTailAndHead() tailLength = %d, want %d, headLength = %d, want %d", len(gotTail), len(test.tailWant), len(gotHead), len(test.headWant))
			}
		})
	}
}
