package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed testinput.txt
var testInput string

//go:embed input.txt
var benchInput string

func Test_part1Count(t *testing.T) {
	tests := []struct {
		input string
		at    int
		want  int
	}{
		{input: testInput, at: 10, want: 26},
		{input: benchInput, at: 2000000, want: 5508234},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Count(ilr.New(tt.input), tt.at); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Count(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Count(ilr.New(benchInput), 2000000); c != 5508234 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_part2Frequency(t *testing.T) {
	tests := []struct {
		input  string
		maxPos int
		want   int
	}{
		{input: testInput, maxPos: 20, want: 56000011},
		{input: benchInput, maxPos: 4000000, want: 10457634860779},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Frequency(ilr.New(tt.input), tt.maxPos); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Frequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Frequency(ilr.New(benchInput), 4000000); c != 10457634860779 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}
