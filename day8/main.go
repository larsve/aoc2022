package main

import (
	"fmt"

	"aoc2022/pkg/ilr"
)

func run(part string, pf func(*ilr.LineReader) int) {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("%s: %d\n", part, pf(f))
}

func main() {
	run("Part1", part1Visible)
	run("Part2", part2Score)
}
