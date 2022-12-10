package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/pkg/ilr"
)

func l2i(line string) (ic int, ax int) {
	var err error
	p := strings.Split(line, " ")
	switch p[0] {
	case "addx":
		ic = 2
		ax, err = strconv.Atoi(p[1])
		if err != nil {
			fmt.Printf("%s isn't an integer, line: %s, error: %v", p[1], line, err)
			return 0, 0
		}
	case "noop":
		ic = 1
		ax = 0
	default:
		fmt.Printf("Unknown input %s (line: %s)", p[0], line)
		return 0, 0
	}
	return
}

func part1Signal(input *ilr.LineReader) int {
	var c, ic, s, x, ax int
	x = 1
	input.ForEach(func(line string) {
		ic, ax = l2i(line)
		if 39-((c+20)%40) < ic {
			s += ((20+c+ic)/40*40 - 20) * x
		}
		c += ic
		x += ax
	})
	return s
}

func part2Lines(input *ilr.LineReader) string {
	var c, ic, x, ax int
	lines := make([]byte, 0, 246)
	x = 1
	input.ForEach(func(line string) {
		ic, ax = l2i(line)
		for i := 0; i < ic; i++ {
			pd := c%40 - x
			c++
			if pd <= 1 && pd >= -1 {
				lines = append(lines, '#')
			} else {
				lines = append(lines, '.')
			}
			if c%40 == 0 {
				lines = append(lines, '\n')
			}
		}
		x += ax
	})
	return string(lines)
}
