package main

import (
	"fmt"
	"strconv"

	"aoc2022/pkg/ilr"
)

func part1Positions(input *ilr.LineReader) int {
	tp := make(map[string]int)
	visit := func(x, y int) {
		tp[fmt.Sprintf("%d,%d", x, y)]++
	}
	visit(0, 0)

	var hx, hy, tx, ty int
	move := func(n int, dirFunc func()) {
		for m := 0; m < n; m++ {
			dirFunc()
			xd, yd := hx-tx, hy-ty
			if xd >= -1 && xd <= 1 && yd >= -1 && yd <= 1 {
				continue
			}

			if xd < 0 {
				tx--
			} else if xd > 0 {
				tx++
			}
			if yd < 0 {
				ty--
			} else if yd > 0 {
				ty++
			}
			visit(tx, ty)
		}
	}

	input.ForEach(func(line string) {
		dir := line[0]
		n, err := strconv.Atoi(line[2:])
		if err != nil {
			fmt.Printf("ERROR: Failed to get nr of moves from '%s', error: %v", line, err)
			return
		}
		switch dir {
		case 'U':
			move(n, func() { hy-- })
		case 'L':
			move(n, func() { hx-- })
		case 'D':
			move(n, func() { hy++ })
		case 'R':
			move(n, func() { hx++ })
		default:
			fmt.Printf("ERROR: Unknown direction: %s (line: %s)", string(dir), line)
			return
		}
	})
	return len(tp)
}

func part2Positions(input *ilr.LineReader) int {
	tp := make(map[string]int)
	visit := func(x, y int) {
		tp[fmt.Sprintf("%d,%d", x, y)]++
	}
	visit(0, 0)

	var rx, ry [10]int
	move := func(n int, dirFunc func()) {
		for m := 0; m < n; m++ {
			dirFunc()
			for i := 1; i < 10; i++ {
				xd, yd := rx[i-1]-rx[i], ry[i-1]-ry[i]
				if xd >= -1 && xd <= 1 && yd >= -1 && yd <= 1 {
					continue
				}
				if xd < 0 {
					rx[i]--
				} else if xd > 0 {
					rx[i]++
				}
				if yd < 0 {
					ry[i]--
				} else if yd > 0 {
					ry[i]++
				}
			}
			visit(rx[9], ry[9])
		}
	}

	input.ForEach(func(line string) {
		dir := line[0]
		n, err := strconv.Atoi(line[2:])
		if err != nil {
			fmt.Printf("ERROR: Failed to get nr of moves from '%s', error: %v", line, err)
			return
		}

		switch dir {
		case 'U':
			move(n, func() { ry[0]-- })
		case 'L':
			move(n, func() { rx[0]-- })
		case 'D':
			move(n, func() { ry[0]++ })
		case 'R':
			move(n, func() { rx[0]++ })
		default:
			fmt.Printf("ERROR: Unknown direction: %s (line: %s)", string(dir), line)
			return
		}
	})
	return len(tp)
}
