package main

import (
	"aoc2022/pkg/ilr"
)

func parseGrid(input *ilr.LineReader) ([]int, int, int) {
	var grid []int
	var xMax int
	input.ForEach(func(line string) {
		if grid == nil {
			xMax = len(line)
			grid = make([]int, 0, xMax*xMax)
		}
		for _, b := range line {
			grid = append(grid, int(b))
		}
	})

	return grid, xMax, len(grid) / xMax
}

func part1Visible(input *ilr.LineReader) int {
	grid, xMax, yMax := parseGrid(input)

	visible := xMax*2 + yMax*2 - 4

	check := func(x, y, h int) int {
		th := grid[y*xMax+x]
		if th >= 48 {
			th -= 48
		}

		if h >= th {
			return h
		}

		if grid[y*xMax+x] >= 48 {
			visible++
			grid[y*xMax+x] -= 48
		}
		h = th

		return h
	}

	for x := 1; x < xMax-1; x++ {
		th := grid[x] - 48
		for y := 1; y < yMax-1; y++ {
			th = check(x, y, th)
		}
		th = grid[(yMax-1)*xMax+x] - 48
		for y := yMax - 2; y > 0; y-- {
			th = check(x, y, th)
		}
	}

	for y := 1; y < yMax-1; y++ {
		th := grid[y*xMax] - 48
		for x := 1; x < xMax-1; x++ {
			th = check(x, y, th)
		}
		th = grid[y*xMax+xMax-1] - 48
		for x := xMax - 2; x > 0; x-- {
			th = check(x, y, th)
		}
	}

	return visible
}

func part2Score(input *ilr.LineReader) int {
	grid, xMax, yMax := parseGrid(input)

	var maxScore int
	score := func(x, y int) {
		// Current Tree Height
		cth := grid[y*xMax+x]
		var s int

		// Look Up
		for yp := y - 1; yp >= 0; yp-- {
			if grid[yp*xMax+x] >= cth || yp == 0 {
				s = y - yp
				break
			}
		}
		// Look Left
		for xp := x - 1; xp >= 0; xp-- {
			if grid[y*xMax+xp] >= cth || xp == 0 {
				s *= x - xp
				break
			}
		}
		// Look Down
		for yp := y + 1; yp < yMax; yp++ {
			if grid[yp*xMax+x] >= cth || yp == yMax-1 {
				s *= yp - y
				break
			}
		}
		// Look Right
		for xp := x + 1; xp < xMax; xp++ {
			if grid[y*xMax+xp] >= cth || xp == xMax-1 {
				s *= xp - x
				break
			}
		}
		if s > maxScore {
			maxScore = s
		}
	}

	for y := 1; y < yMax-1; y++ {
		for x := 1; x < xMax-1; x++ {
			score(x, y)
		}
	}
	return maxScore
}
