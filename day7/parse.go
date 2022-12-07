package main

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2022/pkg/ilr"
)

type entry struct {
	parent *entry
	childs map[string]*entry
	size   int
}

func parseFS(input *ilr.LineReader) *entry {
	root := &entry{}
	cwd := root

	input.ForEach(func(line string) {
		switch {
		case line == "$ cd /":
			cwd = root
		case line == "$ cd ..":
			if cwd.parent == nil {
				fmt.Printf("ERROR: '%s': already at root", line)
			}

			cwd = cwd.parent
		case strings.HasPrefix(line, "$ cd "):
			e, ok := cwd.childs[line[5:]]
			if !ok {
				fmt.Printf("ERROR: dir '%s' was not found", line[5:])
				return
			}
			cwd = e
		case strings.HasPrefix(line, "$ ls"):
		case strings.HasPrefix(line, "dir "):
			if cwd.childs == nil {
				cwd.childs = make(map[string]*entry)
			}
			cwd.childs[line[4:]] = &entry{parent: cwd}
		default:
			// This should be a file
			parts := strings.Split(line, " ")
			s, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("ERROR: '%s' don't appear to be a file!\n", line)
				return
			}
			cwd.size += s
		}
	})
	var sum func(*entry) int
	sum = func(e *entry) int {
		for _, c := range e.childs {
			e.size += sum(c)
		}
		return e.size
	}
	sum(root)
	return root
}
func part1Size(input *ilr.LineReader) int {
	root := parseFS(input)
	var partSum int
	var findCandidate func(*entry)
	findCandidate = func(e *entry) {
		for _, c := range e.childs {
			findCandidate(c)
		}
		if e.size <= 100000 {
			partSum += e.size
		}
	}
	findCandidate(root)
	return partSum
}

func part2Size(input *ilr.LineReader) int {
	root := parseFS(input)
	free := 70000000 - root.size
	need := 30000000
	cds := root.size
	var findCandidate func(*entry)
	findCandidate = func(e *entry) {
		if free+e.size < need {
			return
		}
		if cds > e.size {
			cds = e.size
		}
		for _, c := range e.childs {
			findCandidate(c)
		}
	}
	findCandidate(root)
	return cds
}
