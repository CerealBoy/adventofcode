package year2025

import (
	"bufio"
	"os"
	"slices"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[7] = day7
}

func day7(ctx *shared.Context) error {
	debug := ctx.Debug

	f, _ := os.Open(shared.File(2025, 7, ctx.Test))
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// find the S
	beams, fullBeams := []int{}, map[int]int{}
	for idx, x := range strings.Split(lines[0], "") {
		fullBeams[idx] = 0
		if x == "S" {
			beams = append(beams, idx)
			fullBeams[idx] = 1
		}
	}

	part1 := 0
	for depth := 1; depth < len(lines); depth++ {
		currentLine := strings.Split(lines[depth], "")

		newBeams := []int{}
		for _, x := range beams {
			if currentLine[x] == "^" {
				part1++
				if debug {
					println("Splitting at", x, "on line", depth)
				}
				newBeams = append(newBeams, x-1, x+1)
			} else {
				newBeams = append(newBeams, x)
			}
		}
		beams = slices.Compact(newBeams)

		altBeams := map[int]int{}
		for idx, x := range fullBeams {
			if currentLine[idx] == "^" {
				if debug {
					println("Splitting at", idx, "on line", depth)
				}
				altBeams[idx-1] += x
				altBeams[idx+1] += x
			} else {
				altBeams[idx] += x
			}
		}
		fullBeams = altBeams
	}
	println("Part 1:", part1)

	part2 := 0
	for _, x := range fullBeams {
		part2 += x
	}
	println("Part 2:", part2)

	return nil
}
