package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed input.txt
var benchInput string

func Test_part1Visible(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "30373\n25512\n65332\n33549\n35390", want: 21},
		{input: "999\n555\n111", want: 9},
		{input: "9999\n6666\n3333\n1111", want: 16},
		{input: "000010000\n011121110\n222222222\n011121110\n000010000", want: 9*5 - 1},
		{input: "000010000\n011121110\n222232222\n011121110\n000010000", want: 9 * 5},
		{input: benchInput, want: 1843},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Visible(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Visible(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Visible(ilr.New(benchInput)); c != 1843 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_part2Score(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "30373\n25512\n65332\n33549\n35390", want: 8},
		{input: "000010000\n011121110\n222222222\n011121110\n000010000", want: 16},
		{input: "000010000\n011121110\n222232222\n011121110\n000010000", want: 64},
		{input: benchInput, want: 180000},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Score(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Score(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Score(ilr.New(benchInput)); c != 180000 {
			b.Fatalf("ERROR: wrong score: %d", c)
		}
	}
}
