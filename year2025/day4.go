package year2025

import (
	"bufio"
	"os"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

const (
	PAPER    = "@"
	EMPTY    = "."
	REMOVING = "x"
)

func init() {
	days[4] = day4
}

func day4(ctx *shared.Context) error {
	debug := ctx.Debug
	_ = debug

	f, _ := os.Open(shared.File(2025, 4, ctx.Test))
	scanner := bufio.NewScanner(f)
	lines := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, ""))
	}

	part1, part2 := 0, 0

	width, height := len(lines), len(lines[0])
	for ya, yb := range lines {
		for xa, xb := range yb {
			// If it's not paper there's nothing to access
			if xb != PAPER {
				continue
			}

			count := 0
			if ya > 0 {
				// northwest
				if xa > 0 && lines[ya-1][xa-1] == PAPER {
					count++
				}
				// north
				if lines[ya-1][xa] == PAPER {
					count++
				}
				// northeast
				if xa < height-1 && lines[ya-1][xa+1] == PAPER {
					count++
				}
			}
			// west
			if xa > 0 && lines[ya][xa-1] == PAPER {
				count++
			}
			// east
			if xa < height-1 && lines[ya][xa+1] == PAPER {
				count++
			}
			if ya < width-1 {
				// southwest
				if xa > 0 && lines[ya+1][xa-1] == PAPER {
					count++
				}
				// south
				if lines[ya+1][xa] == PAPER {
					count++
				}
				// southeast
				if xa < height-1 && lines[ya+1][xa+1] == PAPER {
					count++
				}
			}

			if count < 4 {
				part1++
			}
		}
	}

	println("Part 1:", part1)

	width, height = len(lines), len(lines[0])
looping:
	newCount := 0
	for ya, yb := range lines {
		for xa, xb := range yb {
			// If it's not paper there's nothing to access
			if xb != PAPER {
				continue
			}

			count := 0
			if ya > 0 {
				// northwest
				if xa > 0 && (lines[ya-1][xa-1] == PAPER || lines[ya-1][xa-1] == REMOVING) {
					count++
				}
				// north
				if lines[ya-1][xa] == PAPER || lines[ya-1][xa] == REMOVING {
					count++
				}
				// northeast
				if xa < height-1 && (lines[ya-1][xa+1] == PAPER || lines[ya-1][xa] == REMOVING) {
					count++
				}
			}
			// west
			if xa > 0 && (lines[ya][xa-1] == PAPER || lines[ya][xa-1] == REMOVING) {
				count++
			}
			// east
			if xa < height-1 && (lines[ya][xa+1] == PAPER || lines[ya][xa+1] == REMOVING) {
				count++
			}
			if ya < width-1 {
				// southwest
				if xa > 0 && (lines[ya+1][xa-1] == PAPER || lines[ya+1][xa-1] == REMOVING) {
					count++
				}
				// south
				if lines[ya+1][xa] == PAPER || lines[ya+1][xa] == REMOVING {
					count++
				}
				// southeast
				if xa < height-1 && (lines[ya+1][xa+1] == PAPER || lines[ya+1][xa+1] == REMOVING) {
					count++
				}
			}

			if count < 4 {
				lines[ya][xa] = REMOVING
				newCount++
				part2++
			}
		}
	}

	for ya, yb := range lines {
		for xa, xb := range yb {
			if xb == REMOVING {
				lines[ya][xa] = EMPTY
			}
		}
	}

	if newCount > 0 {
		goto looping
	}

	println("Part 2:", part2)

	return nil
}
