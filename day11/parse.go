package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc2022/pkg/ilr"
)

type (
	monkeys []*monkey
	monkey  struct {
		items []int
		op    func(old int) int
		td    int
		tm    int
		fm    int
		ic    int
		lcm   int
	}
)

func (m monkeys) play(part int, rounds int) {
	for r := 0; r < rounds; r++ {
		for _, mn := range m {
			mn.play(part, m)
		}
	}
}

func (m monkeys) calcLcm() {
	var td []int
	for _, mn := range m {
		td = append(td, mn.td)
	}
	mlcm := lcm(td[0], td[1], td[2:]...)
	for _, mn := range m {
		mn.lcm = mlcm
	}
}

func (m *monkey) play(part int, mb monkeys) {
	var tmn int
	for _, i := range m.items {
		m.ic++
		i = m.op(i)
		if part == 1 {
			i = i / 3
		} else {
			i = i % m.lcm
		}
		tmn = m.fm
		if i%m.td == 0 {
			tmn = m.tm
		}
		mb[tmn].items = append(mb[tmn].items, i)
	}
	m.items = m.items[:0]
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func line2items(line string) (items []int) {
	nrs := strings.Split(line[18:], ",")
	for _, nr := range nrs {
		i, err := strconv.Atoi(strings.TrimSpace(nr))
		if err != nil {
			fmt.Printf("%s is not a number, line: %s, error: %v", nr, line, err)
			return
		}
		items = append(items, i)
	}
	return
}

func line2op(line string) func(int) int {
	f := line[23]
	o := line[25:]
	if o == "old" {
		if f == '+' {
			return func(old int) int { return old + old }
		}
		return func(old int) int { return old * old }
	}

	n, err := strconv.Atoi(o)
	if err != nil {
		fmt.Printf("ERROR: %s is not a number, line: %s, error: %v", o, line, err)
		return nil
	}

	if f == '+' {
		return func(old int) int { return old + n }
	}
	return func(old int) int { return old * n }
}

func line2int(line string) int {
	n, err := strconv.Atoi(strings.TrimSpace(line[len(line)-2:]))
	if err != nil {
		fmt.Printf("ERROR: line: %s, error: %v", line, err)
		return -1
	}
	return n
}

func parse(input *ilr.LineReader) monkeys {
	var ln int
	var m *monkey
	var mb monkeys
	input.ForEach(func(line string) {
		switch ln % 7 {
		case 0:
			m = &monkey{}
			mb = append(mb, m)
		case 1:
			m.items = line2items(line)
		case 2:
			m.op = line2op(line)
		case 3:
			m.td = line2int(line)
		case 4:
			m.tm = line2int(line)
		case 5:
			m.fm = line2int(line)
		}
		ln++
	})
	return mb
}

func part1Score(input *ilr.LineReader) int {
	monkeys := parse(input)
	monkeys.play(1, 20)
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].ic > monkeys[j].ic })
	return monkeys[0].ic * monkeys[1].ic
}

func part2Score(input *ilr.LineReader) int {
	monkeys := parse(input)
	monkeys.calcLcm()
	// start := time.Now()
	monkeys.play(2, 10000)
	// d := time.Since(start)
	// fmt.Printf("Play time: %v\n", d)
	sort.Slice(monkeys, func(i, j int) bool { return monkeys[i].ic > monkeys[j].ic })
	return monkeys[0].ic * monkeys[1].ic
}
