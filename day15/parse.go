package main

import (
	"fmt"

	"aoc2022/pkg/ilr"
)

func abs(v int) int {
	if v >= 0 {
		return v
	}
	return -v
}

func dist(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func part1Count(input *ilr.LineReader, atLine int) int {
	var beacon = make(map[int]bool)
	var xMap = make(map[int]bool)
	var x1, y1, x2, y2 int
	input.ForEach(func(line string) {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)
		if y2 == atLine {
			beacon[x2] = true
		}
		d := dist(x1, y1, x2, y2)
		if y1+d < atLine || y1-d > atLine {
			return
		}
		w := d - abs(y1-atLine) + 1
		if w == 1 {
			if beacon[x1] {
				return
			}
			xMap[x1] = true
			return
		}
		xs := x1 - w + 1
		xe := x1 + w
		for x := xs; x < xe; x++ {
			if beacon[x] {
				continue
			}
			xMap[x] = true
		}
	})
	return len(xMap)
}

type (
	pos struct {
		x int
		y int
	}
	sensor struct {
		pos
		d int
	}
)

func part2Frequency(input *ilr.LineReader, maxPos int) int {
	var sensorsPerimiter []pos
	var sensors []sensor
	var x1, y1, x2, y2 int
	input.ForEach(func(line string) {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &x1, &y1, &x2, &y2)
		d := dist(x1, y1, x2, y2)
		sensors = append(sensors, sensor{pos: pos{x1, y1}, d: d})
		p := pos{x: x1, y: y1 - d - 1}
		for i := 0; i <= d; i++ {
			if p.x >= 0 && p.y >= 0 && p.x <= maxPos && p.y <= maxPos {
				sensorsPerimiter = append(sensorsPerimiter, p)
			}
			p.x++
			p.y++
		}
		for i := 0; i <= d; i++ {
			if p.x >= 0 && p.y >= 0 && p.x <= maxPos && p.y <= maxPos {
				sensorsPerimiter = append(sensorsPerimiter, p)
			}
			p.x--
			p.y++
		}
		for i := 0; i <= d; i++ {
			if p.x >= 0 && p.y >= 0 && p.x <= maxPos && p.y <= maxPos {
				sensorsPerimiter = append(sensorsPerimiter, p)
			}
			p.x--
			p.y--
		}
		for i := 0; i <= d; i++ {
			if p.x >= 0 && p.y >= 0 && p.x <= maxPos && p.y <= maxPos {
				sensorsPerimiter = append(sensorsPerimiter, p)
			}
			p.x++
			p.y--
		}
		if p.x >= 0 && p.y >= 0 && p.x <= maxPos && p.y <= maxPos {
			sensorsPerimiter = append(sensorsPerimiter, p)
		}
	})

	var covered bool
	for _, p := range sensorsPerimiter {
		covered = false
		for _, s := range sensors {
			if dist(p.x, p.y, s.x, s.y) <= s.d {
				covered = true
				break
			}
		}
		if !covered {
			return p.x*4000000 + p.y
		}
	}
	return -1
}
