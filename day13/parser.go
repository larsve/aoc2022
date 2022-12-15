package main

import (
	"aoc2022/pkg/ilr"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
)

func fors(v any) (float64, []any, bool) {
	switch val := v.(type) {
	case float64:
		return val, []any{val}, true
	case []any:
		return 0, val, false
	default:
		fmt.Printf("Unknown type: %T, value: %v\n", v, v)
		return 0, nil, true
	}
}

func check(l, r any) int {
	lf, ls, lfok := fors(l)
	rf, rs, rfok := fors(r)

	if lfok && rfok {
		return int(lf - rf)
	}

	for i := 0; i < len(ls) && i < len(rs); i++ {
		if n := check(ls[i], rs[i]); n != 0 {
			return n
		}
	}

	return len(ls) - len(rs)
}

func part1Sum(input *ilr.LineReader) int {
	var err error
	var idx, row, sum int
	var l, r any
	input.ForEach(func(line string) {
		row++
		switch row % 3 {
		case 1:
			err = json.Unmarshal([]byte(line), &l)
		case 2:
			err = json.Unmarshal([]byte(line), &r)
		default:
			return
		}
		if err != nil {
			fmt.Printf("Unmarshal failed with error: %v\n", err)
			return
		}

		if row%3 != 2 {
			return
		}
		idx++
		if check(l, r) <= 0 {
			sum += idx
		}
	})
	return sum
}

func part2Sum(input *ilr.LineReader) int {
	var err error
	dp1 := []any{[]any{2.0}}
	dp2 := []any{[]any{6.0}}
	lines := []any{dp1, dp2}
	input.ForEach(func(line string) {
		if line == "" {
			return
		}
		var l any
		err = json.Unmarshal([]byte(line), &l)
		if err != nil {
			fmt.Printf("Unmarshal failed with error: %v\n", err)
			return
		}
		lines = append(lines, l)
	})
	sort.Slice(lines, func(i, j int) bool { return check(lines[i], lines[j]) < 0 })
	var dpi1, dpi2 int
	for i, l := range lines {
		switch {
		case reflect.DeepEqual(l, dp1):
			dpi1 = i + 1
		case reflect.DeepEqual(l, dp2):
			dpi2 = i + 1
		}
	}
	return dpi1 * dpi2
}
