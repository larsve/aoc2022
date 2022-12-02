package main

import (
	"fmt"
	"os"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer f.Close()

	s := scorePart1(f)

	fmt.Printf("Part1: RPS score: %d", s)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer f.Close()

	s := scorePart2(f)

	fmt.Printf("Part2: RPS score: %d", s)
}

func main() {
	part1()
	part2()
}
