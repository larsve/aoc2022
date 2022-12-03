package main

import (
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

func Test_prio(t *testing.T) {
	tests := []struct {
		r    rune
		want int
	}{
		{r: 'p', want: 16},
		{r: 'L', want: 38},
		{r: 'P', want: 42},
		{r: 'v', want: 22},
		{r: 't', want: 20},
		{r: 's', want: 19},
		{r: '@', want: 0},
		{r: '[', want: 0},
		{r: '`', want: 0},
		{r: '{', want: 0},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := prio(tt.r); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Test_part1Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw",
			want:  157,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := part1Sum(ilr.New(tt.input))
			if got != tt.want {
				t.Errorf("ERROR: got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_part2Sum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "vJrwpWtwJgWrhcsFMMfFFhFp\njqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\nPmmdzqPrVvPwwTWBwg\nwMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\nttgJtRGJQctTZtZT\nCrZsJsPPZsGzwwsLwLmpwMDw",
			want:  70,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := part2Sum(ilr.New(tt.input))
			if got != tt.want {
				t.Errorf("ERROR: got: %v, want: %v", got, tt.want)
			}
		})
	}
}
