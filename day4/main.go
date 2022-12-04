package main

import (
	"fmt"

	"aoc2022/pkg/ilr"
)

func run(part string, sum func(*ilr.LineReader) int) {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("%s sum: %d\n", part, sum(f))
}

func main() {
	run("Part1", part1Sum)
	run("Part2", part2Sum)
}
