package main

import (
	"fmt"

	"aoc2022/pkg/ilr"
)

func run(part string, size func(*ilr.LineReader) int) {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("%s size: %d\n", part, size(f))
}

func main() {
	run("Part1", part1Size)
	run("Part2", part2Size)
}
