package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed input.txt
var benchInput string

func Test_part1Positions(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2", want: 13},
		{input: benchInput, want: 6044},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Positions(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Positions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Positions(ilr.New(benchInput)); c != 6044 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_part2Positions(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "R 4\nU 4\nL 3\nD 1\nR 4\nD 1\nL 5\nR 2", want: 1},
		{input: "R 5\nU 8\nL 8\nD 3\nR 17\nD 10\nL 25\nU 20", want: 36},
		{input: benchInput, want: 2384},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Positions(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Positions(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Positions(ilr.New(benchInput)); c != 2384 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}
