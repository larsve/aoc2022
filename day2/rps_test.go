package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed input.txt
var benchInput string

func Test_part1Score(t *testing.T) {
	tests := []struct {
		input     string
		wantScore int
	}{
		{input: "A Y\nB X\nC Z", wantScore: 15},
		{input: "A Z\nB X\nC Y", wantScore: 6},
		{input: "A X\nB Y\nC Z", wantScore: 15},
		{input: "A Y\nB Z\nC X", wantScore: 24},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Score(ilr.New(tt.input)); got != tt.wantScore {
				t.Errorf("ERROR: got score: %v, want %v", got, tt.wantScore)
			}
		})
	}
}

func Benchmark_part1Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if sum := part1Score(ilr.New(benchInput)); sum != 9651 {
			b.Fatalf("ERROR: wrong sum")
		}
	}
}

func Test_part2Score(t *testing.T) {
	tests := []struct {
		input     string
		wantScore int
	}{
		{input: "A Y\nB X\nC Z", wantScore: 12},
		{input: "A X\nB X\nC X", wantScore: 6},
		{input: "A Y\nB Y\nC Y", wantScore: 15},
		{input: "A Z\nB Z\nC Z", wantScore: 24},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Score(ilr.New(tt.input)); got != tt.wantScore {
				t.Errorf("ERROR: got score: %v, want %v", got, tt.wantScore)
			}
		})
	}
}

func Benchmark_part2Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if sum := part2Score(ilr.New(benchInput)); sum != 10560 {
			b.Fatalf("ERROR: wrong sum")
		}
	}
}
