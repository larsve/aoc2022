package main

import (
	_ "embed"
	"fmt"
	"testing"

	"aoc2022/pkg/ilr"
)

//go:embed testsignalinput.txt
var testSignal string

//go:embed input.txt
var benchInput string

func Test_patr1Signal(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: testSignal, want: 13140},
		{input: benchInput, want: 15880},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part1Signal(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR: got: %d want: %d", got, tt.want)
			}
		})
	}
}

func Benchmark_part1Signal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part1Signal(ilr.New(benchInput)); c != 15880 {
			b.Fatalf("ERROR: wrong number: %d", c)
		}
	}
}

func Test_patr2Lines(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: testSignal, want: "##..##..##..##..##..##..##..##..##..##..\n###...###...###...###...###...###...###.\n####....####....####....####....####....\n#####.....#####.....#####.....#####.....\n######......######......######......####\n#######.......#######.......#######.....\n"},
		{input: benchInput, want: "###..#.....##..####.#..#..##..####..##..\n#..#.#....#..#.#....#.#..#..#....#.#..#.\n#..#.#....#....###..##...#..#...#..#....\n###..#....#.##.#....#.#..####..#...#.##.\n#....#....#..#.#....#.#..#..#.#....#..#.\n#....####..###.#....#..#.#..#.####..###.\n"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			if got := part2Lines(ilr.New(tt.input)); got != tt.want {
				t.Errorf("ERROR:\ngot:\n%s\nwant:\n%s", got, tt.want)
			}
		})
	}
}

func Benchmark_part2Lines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if c := part2Lines(ilr.New(benchInput)); c != "###..#.....##..####.#..#..##..####..##..\n#..#.#....#..#.#....#.#..#..#....#.#..#.\n#..#.#....#....###..##...#..#...#..#....\n###..#....#.##.#....#.#..####..#...#.##.\n#....#....#..#.#....#.#..#..#.#....#..#.\n#....####..###.#....#..#.#..#.####..###.\n" {
			b.Fatalf("ERROR: wrong output:\n%s", c)
		}
	}
}
