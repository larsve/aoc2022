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

	fmt.Printf("Part1 sum: %d\n", part1Sum(f))
}

func part2() {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("Part2 sum: %d\n", part2Sum(f))
}

func main() {
	part1()
	part2()
}
