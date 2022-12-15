package main

import (
	"aoc2022/pkg/ilr"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var benchInput string

func Test_part1Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]\n",
			want:  13,
		},
		{input: benchInput, want: 6428},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Sum(ilr.New(benchInput)); c != 6428 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_part2Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]\n",
			want:  140,
		},
		{input: benchInput, want: 22464},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Sum(ilr.New(benchInput)); c != 22464 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}
