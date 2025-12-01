package year2019

import (
	"bufio"
	"os"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[1] = day1
}

func day1(ctx *shared.Context) error {
	f, _ := os.Open(shared.File(2019, 1, ctx.Test))
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	part1 := 0
	for _, x := range lines {
		val, _ := strconv.Atoi(x)
		part1 += (val / 3) - 2
	}
	println("Part 1:", part1)

	part2 := 0
	for _, x := range lines {
		val, _ := strconv.Atoi(x)
		part2 += day1Compound(val, 0)
	}
	println("Part 2:", part2)

	return nil
}

func day1Compound(val, total int) int {
	if val <= 8 {
		return total
	}

	newVal := (val / 3) - 2
	total += newVal
	return day1Compound(newVal, total)
}
