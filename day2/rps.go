package main

import (
	"bufio"
	"io"
)

func scorePart1(input io.Reader) int {
	r := bufio.NewScanner(input)
	r.Split(bufio.ScanLines)
	var score int
	for r.Scan() {
		switch r.Text() {
		case "A X": // Rock = Rock (1)
			score += 4
		case "A Y": // Rock < Paper (2)
			score += 8
		case "A Z": // Rock > Scissors (3)
			score += 3
		case "B X": // Paper > Rock (1)
			score += 1
		case "B Y": // Paper = Paper (2)
			score += 5
		case "B Z": // Paper < Scissors (3)
			score += 9
		case "C X": // Scissors < Rock (1)
			score += 7
		case "C Y": // Scissors > Paper (2)
			score += 2
		case "C Z": // Scissors = Scissors (3)
			score += 6
		}
	}
	return score
}

func scorePart2(input io.Reader) int {
	r := bufio.NewScanner(input)
	r.Split(bufio.ScanLines)
	var score int
	for r.Scan() {
		switch r.Text() {
		// Lose
		case "A X": // Rock > Scissors (3)
			score += 3
		case "B X": // Paper > Rock (1)
			score += 1
		case "C X": // Scissors < Paper (2)
			score += 2
		// Draw
		case "A Y": // Rock = Rock (1)
			score += 4
		case "B Y": // Paper = Paper (2)
			score += 5
		case "C Y": // Scissors = Scissors (3)
			score += 6
		// Win
		case "A Z": // Rock < Paper (2)
			score += 8
		case "B Z": // Paper < Scissors (3)
			score += 9
		case "C Z": // Scissors > Rock (1)
			score += 7
		}
	}
	return score
}
