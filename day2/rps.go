package main

import (
	"aoc2022/pkg/ilr"
)

func part1Score(input *ilr.LineReader) int {
	var sum int
	input.ForEach(func(line string) {
		switch line {
		case "A X": // Rock = Rock (1)
			sum += 4
		case "A Y": // Rock < Paper (2)
			sum += 8
		case "A Z": // Rock > Scissors (3)
			sum += 3
		case "B X": // Paper > Rock (1)
			sum += 1
		case "B Y": // Paper = Paper (2)
			sum += 5
		case "B Z": // Paper < Scissors (3)
			sum += 9
		case "C X": // Scissors < Rock (1)
			sum += 7
		case "C Y": // Scissors > Paper (2)
			sum += 2
		case "C Z": // Scissors = Scissors (3)
			sum += 6
		}
	})
	return sum
}

func part2Score(input *ilr.LineReader) int {
	var sum int
	input.ForEach(func(line string) {
		switch line {
		// Lose
		case "A X": // Rock > Scissors (3)
			sum += 3
		case "B X": // Paper > Rock (1)
			sum += 1
		case "C X": // Scissors < Paper (2)
			sum += 2
		// Draw
		case "A Y": // Rock = Rock (1)
			sum += 4
		case "B Y": // Paper = Paper (2)
			sum += 5
		case "C Y": // Scissors = Scissors (3)
			sum += 6
		// Win
		case "A Z": // Rock < Paper (2)
			sum += 8
		case "B Z": // Paper < Scissors (3)
			sum += 9
		case "C Z": // Scissors > Rock (1)
			sum += 7
		}
	})
	return sum
}
