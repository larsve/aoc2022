package main

import (
	"strconv"
	"strings"

	"aoc2022/pkg/ilr"
)

func line2move(line string) (int, int, int) {
	v := strings.Split(line, " ")
	n, err := strconv.Atoi(v[1])
	if err != nil {
		panic(err)
	}
	f, err := strconv.Atoi(v[3])
	if err != nil {
		panic(err)
	}
	t, err := strconv.Atoi(v[5])
	if err != nil {
		panic(err)
	}
	return n, f - 1, t - 1
}

func crane(input *ilr.LineReader, move func([]string, int, int, int)) string {
	var stacks []string
	input.ForEach(func(line string) {
		if strings.Contains(line, "[") {
			if len(stacks) == 0 {
				stacks = make([]string, len(line)/4+1)
			}
			for i := 0; i < len(line); i += 4 {
				if line[i+1] == ' ' {
					continue
				}
				stacks[i/4] += string(line[i+1])
			}
			return
		}

		if !strings.HasPrefix(line, "move") {
			return
		}

		n, f, t := line2move(line)
		move(stacks, n, f, t)
	})

	var res string
	for _, s := range stacks {
		if len(s) == 0 {
			continue
		}
		res += string(s[0])
	}
	return res
}

func part1Crates(input *ilr.LineReader) string {
	return crane(input, func(stacks []string, n, f, t int) {
		crates := stacks[f][:n]
		stacks[f] = stacks[f][n:]
		for _, c := range crates {
			stacks[t] = string(c) + stacks[t]
		}
	})
}

func part2Crates(input *ilr.LineReader) string {
	return crane(input, func(stacks []string, n, f, t int) {
		crates := stacks[f][:n]
		stacks[f] = stacks[f][n:]
		stacks[t] = crates + stacks[t]
	})
}
