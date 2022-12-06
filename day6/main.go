package main

import (
	"fmt"
	"io"
	"os"
)

func run(part string, ofs func(io.Reader) int) {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Failed to open file, error: %v\n", err)
		return
	}
	defer f.Close()

	fmt.Printf("%s offset: %d\n", part, ofs(f))
}

func main() {
	run("Part1", part1Ofs)
	run("Part2", part2Ofs)
}
