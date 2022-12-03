package main

import (
	"sort"
	"strconv"

	"aoc2022/pkg/ilr"
)

func part1Sum(input *ilr.LineReader) int {
	var sum, es int
	input.ForEach(func(line string) {
		if line == "" {
			if es > sum {
				sum = es
			}
			es = 0
			return
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		es += n
	})
	return sum
}

func part2Sum(input *ilr.LineReader) int {
	var perElf []int
	var es int
	input.ForEach(func(line string) {
		if line == "" {
			perElf = append(perElf, es)
			es = 0
			return
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		es += n

	})
	if es > 0 {
		perElf = append(perElf, es)
	}

	// Pick top 3
	sort.Sort(sort.Reverse(sort.IntSlice(perElf)))
	perElf = perElf[:3]

	var sum int
	for _, cal := range perElf {
		sum += cal
	}

	return sum
}
