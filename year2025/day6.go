package year2025

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[6] = day6
}

func day6(ctx *shared.Context) error {
	debug := ctx.Debug
	_ = debug

	f, _ := os.Open(shared.File(2025, 6, ctx.Test))
	scanner := bufio.NewScanner(f)
	lines, fullLines := [][]string{}, []string{}

	re := regexp.MustCompile(`\s+`)
	for scanner.Scan() {
		line := scanner.Text()
		fullLines = append(fullLines, line)
		lineSlice := strings.Split(re.ReplaceAllString(strings.TrimSpace(line), " "), " ")
		lines = append(lines, lineSlice)
	}

	part1, part2 := 0, 0

	numProblems, numValues := len(lines[0])-1, len(lines)-1
	for x := 0; x <= numProblems; x++ {
		num, _ := strconv.Atoi(lines[0][x])
		result := num
		op := lines[numValues][x]

		for y := 1; y < numValues; y++ {
			num, _ = strconv.Atoi(lines[y][x])
			switch op {
			case "+":
				result += num

			case "*":
				result *= num

			}
		}

		if debug {
			println("Result for entry", x, "is", result)
		}
		part1 += result
	}
	println("Part 1:", part1)

	offset := 0
	// Begin at 0, grab the operator
	// Determine how many spaces until the next operator
	// From the right-most column minus 1, assemble the numbers from top to bottom
	// Use the operator to combine those numbers
	// Add to the resulting output

	for x := 0; x < len(fullLines[0]); x++ {
		op := string(fullLines[len(fullLines)-1][x])
		if op == " " {
			continue
		}

		// we have an operator
		result := 0
		for y := x + 1; y < len(fullLines[0]); y++ {
			offset = y
			if string(fullLines[len(fullLines)-1][y]) != " " {
				offset -= 2
				break
			}
		}
		if debug {
			println("At idx", x, "we have an offset of", offset)
		}

		// Pull the numbers
		for y := offset; y >= x; y-- {
			val := ""
			for col := 0; col < numValues; col++ {
				val = fmt.Sprintf("%s%s", val, string(fullLines[col][y]))
			}
			num, _ := strconv.Atoi(strings.TrimSpace(val))

			if result == 0 {
				result = num
			} else {
				switch op {
				case "+":
					result += num
				case "*":
					result *= num
				}
			}

			if debug {
				println("Result currently", result)
			}
		}

		part2 += result
	}

	println("Part 2:", part2)

	return nil
}
