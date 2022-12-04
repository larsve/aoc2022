package main

import (
	"strconv"
	"strings"

	"aoc2022/pkg/ilr"
)

func splitLine(line string) (s1, e1, s2, e2 int) {
	strPairs := strings.Split(line, ",")

	s1, e1 = splitRange(strPairs[0])
	s2, e2 = splitRange(strPairs[1])
	return
}

func splitRange(r string) (int, int) {
	p := strings.Split(r, "-")
	s, err := strconv.Atoi(p[0])
	if err != nil {
		panic(err)
	}
	e, err := strconv.Atoi(p[1])
	if err != nil {
		panic(err)
	}
	return s, e
}

func part1Sum(input *ilr.LineReader) int {
	var sum int
	input.ForEach(func(line string) {
		s1, e1, s2, e2 := splitLine(line)

		if (s1 <= s2 && e1 >= e2) || (s2 <= s1 && e2 >= e1) {
			sum++
		}
	})
	return sum
}

func part2Sum(input *ilr.LineReader) int {
	var sum int
	input.ForEach(func(line string) {
		s1, e1, s2, e2 := splitLine(line)

		if (s1 <= s2 && e1 >= s2) || (s1 <= e2 && e1 >= e2) ||
			(s2 <= s1 && e2 >= s1) || (s2 <= e1 && e2 >= e1) {
			sum++
		}
	})
	return sum
}
