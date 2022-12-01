package main

import (
	"bufio"
	"io"
	"strconv"
)

func forEachElf(inventory io.Reader, f func([]int)) error {
	r := bufio.NewScanner(inventory)
	r.Split(bufio.ScanLines)
	var items []int
	for r.Scan() {
		line := r.Text()
		if line == "" {
			f(items)
			items = items[:0]
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		items = append(items, n)
	}

	if len(items) > 0 {
		f(items)
	}

	return nil
}
