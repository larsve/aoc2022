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

	fmt.Printf("Part1 signal strength: %d\n", part1Signal(f))
}

func part2() {
	f, close, err := ilr.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer close()

	fmt.Printf("Part2 CRT output: \n%s\n", part2Lines(f))
}

func main() {
	part1()
	part2()
}
