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
		{input: "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9", want: 24},
		{input: benchInput, want: 644},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Sum(ilr.New(benchInput)); c != 644 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_part2Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9", want: 93},
		{input: benchInput, want: 27324},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Sum(ilr.New(benchInput)); c != 27324 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}
