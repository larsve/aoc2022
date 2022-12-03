package main

import (
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

func Test_part1Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", want: 24000},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_part2Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "1000\n2000\n3000\n\n4000\n\n5000\n6000\n\n7000\n8000\n9000\n\n10000", want: 45000},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Sum(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want %d", got, tt.want)
			}
		})
	}
}
