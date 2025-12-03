package year2025

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[3] = day3
}

func day3(ctx *shared.Context) error {
	debug := ctx.Debug

	f, _ := os.Open(shared.File(2025, 3, ctx.Test))
	scanner := bufio.NewScanner(f)
	lines := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		digits := []int{}
		for _, x := range strings.Split(line, "") {
			num, _ := strconv.Atoi(x)
			digits = append(digits, num)
		}

		if debug {
			fmt.Printf("Have the line %#v\n", digits)
		}

		lines = append(lines, digits)
	}

	part1, part2 := 0, int64(0)

	// Find the largest 2 digit number without rearranging the order of the digits
	for _, x := range lines {
		maximum, foundIdx := 0, 0
		for idx, entry := range x {
			if debug {
				fmt.Printf("%d[%d]\n", idx, entry)
			}

			if entry > maximum && idx < len(x)-1 {
				maximum = entry
				foundIdx = idx
			}
		}

		secondMaximum := 0
		for idx, entry := range x {
			if idx > foundIdx && entry > secondMaximum {
				secondMaximum = entry
			}
		}

		if debug {
			fmt.Printf("Found that %d%d is the maximum!\n", maximum, secondMaximum)
		}

		part1 += (maximum * 10) + secondMaximum
	}
	fmt.Printf("Part 1: %#v\n", part1)

	// Now it's a 12 digit number
	for _, x := range lines {
		maximum, foundIdx, total, depth := 0, 0, int64(0), 11
		for idx, entry := range x {
			if entry > maximum && idx < len(x)-depth {
				maximum = entry
				foundIdx = idx
			}
		}

	dwindle:
		depth--
		total *= 10
		total += int64(maximum)
		maximum = 0
		for idx, entry := range x {
			if idx > foundIdx && idx < len(x)-depth && entry > maximum {
				foundIdx = idx
				maximum = entry
			}
		}
		if depth >= 0 {
			goto dwindle
		}

		part2 += total
	}
	fmt.Printf("Part 2: %#v\n", part2)

	return nil
}
