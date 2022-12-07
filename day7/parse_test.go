package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed input.txt
var benchInput string

func Test_part1Size(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k", want: 95437},
		{input: benchInput, want: 1334506},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Size(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Size(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Size(ilr.New(benchInput)); c != 1334506 {
			b.Fatalf("ERROR: wrong size: %d", c)
		}
	}
}

func Test_part2Size(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: "$ cd /\n$ ls\ndir a\n14848514 b.txt\n8504156 c.dat\ndir d\n$ cd a\n$ ls\ndir e\n29116 f\n2557 g\n62596 h.lst\n$ cd e\n$ ls\n584 i\n$ cd ..\n$ cd ..\n$ cd d\n$ ls\n4060174 j\n8033020 d.log\n5626152 d.ext\n7214296 k", want: 24933642},
		{input: benchInput, want: 7421137},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Size(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d, want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Size(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Size(ilr.New(benchInput)); c != 7421137 {
			b.Fatalf("ERROR: wrong size: %d", c)
		}
	}
}
