package main

import (
	"fmt"
	"os"
	"sort"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer f.Close()
	var max int
	err = forEachElf(f, func(cals []int) {
		var sum int
		for _, cal := range cals {
			sum += cal
		}

		if sum > max {
			max = sum
		}
	})
	if err != nil {
		fmt.Printf("Failed to read inventory, error: %v", err)
	}

	fmt.Printf("Part1: Elf carrying calories = %d\n", max)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer f.Close()

	var perElf []int
	err = forEachElf(f, func(cals []int) {
		var sum int
		for _, cal := range cals {
			sum += cal
		}
		perElf = append(perElf, sum)
	})
	if err != nil {
		fmt.Printf("Failed to read inventory, error: %v", err)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(perElf)))
	perElf = perElf[:3]

	var sum int
	for _, cal := range perElf {
		sum += cal
	}

	fmt.Printf("Part2: Top three Elfs: %v, total: %d", perElf, sum)
}

func main() {
	part1()
	part2()
}
