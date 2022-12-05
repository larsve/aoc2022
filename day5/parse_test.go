package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed input.txt
var benchInput string

func Test_part1Crates(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "    [D]    \n[N] [C]    \n[Z] [M] [P]\n	1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", want: "CMZ"},
		{input: benchInput, want: "MQSHJMWNH"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Crates(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got crates: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Crates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Crates(ilr.New(benchInput)); c != "MQSHJMWNH" {
			b.Fatalf("ERROR: wrong top crates")
		}
	}
}

func Test_part2Crates(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "    [D]    \n[N] [C]    \n[Z] [M] [P]\n	1   2   3 \n\nmove 1 from 2 to 1\nmove 3 from 1 to 3\nmove 2 from 2 to 1\nmove 1 from 1 to 2", want: "MCD"},
		{input: benchInput, want: "LLWJRBHVZ"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Crates(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got crates: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Crates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Crates(ilr.New(benchInput)); c != "LLWJRBHVZ" {
			b.Fatalf("ERROR: wrong top crates")
		}
	}
}
