package main

import (
	"aoc2022/pkg/ilr"
	"fmt"
	"strings"
)

type pos struct {
	x int
	y int
}

func loadGrid(input *ilr.LineReader) (map[pos]bool, int) {
	var yMax int
	grid := make(map[pos]bool)
	input.ForEach(func(line string) {
		var points []pos
		for _, s := range strings.Split(line, " -> ") {
			var np pos
			fmt.Sscanf(s, "%d,%d", &np.x, &np.y)
			points = append(points, np)
		}
		current := points[0]
		grid[current] = true
		if yMax < current.y {
			yMax = current.y
		}
		for _, next := range points[1:] {
			var move func(p *pos)
			switch {
			case current.x > next.x:
				move = func(p *pos) { p.x-- }
			case current.x < next.x:
				move = func(p *pos) { p.x++ }
			case current.y > next.y:
				move = func(p *pos) { p.y-- }
			case current.y < next.y:
				move = func(p *pos) { p.y++ }
			}
			for current != next {
				move(&current)
				grid[current] = true
			}
			if yMax < current.y {
				yMax = current.y
			}
		}
	})
	return grid, yMax
}

func part1Sum(input *ilr.LineReader) int {
	grid, yMax := loadGrid(input)
	rocks := len(grid)
outside:
	for {
		g := pos{500, 0}
	drop:
		for {
			switch {
			case !grid[pos{g.x, g.y + 1}]:
				g.y++
			case !grid[pos{g.x - 1, g.y + 1}]:
				g.y++
				g.x--
			case !grid[pos{g.x + 1, g.y + 1}]:
				g.y++
				g.x++
			default:
				grid[g] = true
				break drop
			}
			if g.y > yMax {
				break outside
			}
		}
	}
	return len(grid) - rocks
}

func part2Sum(input *ilr.LineReader) int {
	grid, yMax := loadGrid(input)
	rocks := len(grid)
	start := &pos{500, 0}
	bottom := yMax + 1
outer:
	for {
		g := *start
	drop:
		for {
			switch {
			case !grid[pos{g.x, g.y + 1}] && g.y < bottom:
				g.y++
			case !grid[pos{g.x - 1, g.y + 1}] && g.y < bottom:
				g.y++
				g.x--
			case !grid[pos{g.x + 1, g.y + 1}] && g.y < bottom:
				g.y++
				g.x++
			default:
				grid[g] = true
				if g == *start {
					break outer
				}
				break drop
			}
		}
	}
	return len(grid) - rocks
}
