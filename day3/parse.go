package main

import (
	"strings"

	"aoc2022/pkg/ilr"
)

func prio(r rune) int {
	switch {
	case r >= 'a' && r <= 'z':
		return int(r) - 96
	case r >= 'A' && r <= 'Z':
		return int(r) - 38
	default:
		return 0
	}
}

func part1Sum(input *ilr.LineReader) int {
	var sum int
	input.ForEach(func(line string) {
		mp := len(line) / 2
		c1 := line[:mp]
		c2 := line[mp:]
		var checked = make(map[rune]bool)
		for _, item := range c1 {
			if f := checked[item]; f {
				continue
			}
			checked[item] = true
			if strings.ContainsRune(c2, item) {
				sum += prio(item)
			}
		}
	})
	return sum
}

func part2Sum(input *ilr.LineReader) int {
	var cnt, sum int
	var group [3]string
	input.ForEach(func(line string) {
		group[cnt%3] = line
		cnt++
		if cnt%3 != 0 {
			return
		}

		var checked = make(map[rune]bool)
		// Not that efficient, but less code then checking for, and iterating over the longest item-list
		for _, i := range group[0] + group[1] + group[2] {
			if f := checked[i]; f {
				continue
			}
			checked[i] = true
			if strings.ContainsRune(group[0], i) &&
				strings.ContainsRune(group[1], i) &&
				strings.ContainsRune(group[2], i) {
				sum += prio(i)
			}
		}
	})
	return sum
}
