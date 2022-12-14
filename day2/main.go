package main

import (
	"fmt"

	"aoc2022/pkg/ilr"
)

func part1() {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("Part1 score: %d\n", part1Score(f))
}

func part2() {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("Part2 score: %d\n", part2Score(f))
}

func main() {
	part1()
	part2()
}
