package year2022

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[2] = day2
}

func day2(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 2, ctx.Test))
	if err != nil {
		return err
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0
	for _, line := range lines {
		switch line[0] {
		case 'A':
			if line[2] == 'X' {
				total += 4 // rock, draw
			} else if line[2] == 'Y' {
				total += 8 // paper, win
			} else {
				total += 3 // scissors, lose
			}

		case 'B':
			if line[2] == 'X' {
				total += 1 // rock, lose
			} else if line[2] == 'Y' {
				total += 5 // paper, draw
			} else {
				total += 9 // scissors, win
			}

		case 'C':
			if line[2] == 'X' {
				total += 7 // rock, win
			} else if line[2] == 'Y' {
				total += 2 // paper, lose
			} else {
				total += 6 // scissors, draw
			}

		}
	}

	fmt.Printf("#1: %v\n", total)

	total = 0
	for _, line := range lines {
		switch line[0] {
		case 'A':
			if line[2] == 'X' { // lose
				total += 3
			} else if line[2] == 'Y' { // draw
				total += 4
			} else { // win
				total += 8
			}

		case 'B':
			if line[2] == 'X' { // lose
				total += 1
			} else if line[2] == 'Y' { // draw
				total += 5
			} else { // win
				total += 9
			}

		case 'C':
			if line[2] == 'X' { // lose
				total += 2
			} else if line[2] == 'Y' { // draw
				total += 6
			} else { // win
				total += 7
			}

		}
	}

	fmt.Printf("#2: %v\n", total)

	return nil
}
