package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed input.txt
var benchInput string

func Test_part1Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8", want: 2},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if sum := part1Sum(ilr.New(benchInput)); sum != 494 {
			b.Fatalf("ERROR: wrong sum")
		}
	}
}

func Test_part2Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "2-4,6-8\n2-3,4-5\n5-7,7-9\n2-8,3-7\n6-6,4-6\n2-6,4-8", want: 4},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Sum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if sum := part2Sum(ilr.New(benchInput)); sum != 833 {
			b.Fatalf("ERROR: wrong sum")
		}
	}
}
