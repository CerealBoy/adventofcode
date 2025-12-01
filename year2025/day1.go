package year2025

import (
	"bufio"
	"os"
	"regexp"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[1] = day1
}

func day1(ctx *shared.Context) error {
	f, _ := os.Open(shared.File(2025, 1, ctx.Test))
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	part1Count, part2Count, pointer := 0, 0, 50
	re := regexp.MustCompile(`([LR])(\d+)`)
	for _, x := range lines {
		match := re.FindStringSubmatch(x)
		num, _ := strconv.Atoi(match[2])

		switch match[1] {
		case "L":
			pointer -= num
		case "R":
			pointer += num
		}

		// lol
	positionPointer:
		if pointer < 0 {
			pointer += 100
			part2Count++
			goto positionPointer
		} else if pointer > 99 {
			pointer -= 100
			part2Count++
			goto positionPointer
		}

		if pointer == 0 {
			part1Count++
		}
	}
	println("Part 1:", part1Count)
	println("Part 2:", part2Count)

	return nil
}
