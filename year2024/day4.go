package year2024

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[4] = day4
}

func day4(ctx *shared.Context) error {
	debug = ctx.Debug
	grid := []string{}

	f, _ := os.Open(shared.File(2024, 4, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	// because I'm too brained at the end of the work day, we brute force this

	// looking for XMAS in any direction
	count := 0
	for y, row := range grid {
		for x, col := range strings.Split(row, "") {
			if debug {
				fmt.Println("At", x, ":", y, "letter is", col)
			}

			if col == "X" {
				if x > 2 {
					// left
					if string(grid[y][x-1]) == "M" && string(grid[y][x-2]) == "A" && string(grid[y][x-3]) == "S" {
						count++
					}

					if y > 2 {
						// up left
						if string(grid[y-1][x-1]) == "M" && string(grid[y-2][x-2]) == "A" && string(grid[y-3][x-3]) == "S" {
							count++
						}
					}

					if y < len(grid)-3 {
						// down left
						if string(grid[y+1][x-1]) == "M" && string(grid[y+2][x-2]) == "A" && string(grid[y+3][x-3]) == "S" {
							count++
						}
					}
				}

				if y > 2 {
					// up
					if string(grid[y-1][x]) == "M" && string(grid[y-2][x]) == "A" && string(grid[y-3][x]) == "S" {
						count++
					}

					if x < len(row)-3 {
						// up right
						if string(grid[y-1][x+1]) == "M" && string(grid[y-2][x+2]) == "A" && string(grid[y-3][x+3]) == "S" {
							count++
						}
					}
				}

				if x < len(row)-3 {
					// right
					if string(grid[y][x+1]) == "M" && string(grid[y][x+2]) == "A" && string(grid[y][x+3]) == "S" {
						count++
					}

					if y < len(grid)-3 {
						// right down
						if string(grid[y+1][x+1]) == "M" && string(grid[y+2][x+2]) == "A" && string(grid[y+3][x+3]) == "S" {
							count++
						}
					}
				}

				if y < len(grid)-3 {
					// down
					if string(grid[y+1][x]) == "M" && string(grid[y+2][x]) == "A" && string(grid[y+3][x]) == "S" {
						count++
					}
				}
			}
		}
	}
	fmt.Println("#1:", count)

	// now it's X-MAS
	count = 0
	for y, row := range grid {
		for x, col := range strings.Split(row, "") {
			if y == 0 || y == len(grid)-1 || x == 0 || x == len(row)-1 {
				// quick fail through
				continue
			}

			// we anchor around A here
			if col == "A" {
				if string(grid[y-1][x-1]) == "M" && string(grid[y+1][x+1]) == "S" {
					if string(grid[y-1][x+1]) == "M" && string(grid[y+1][x-1]) == "S" {
						count++
					}
					if string(grid[y-1][x+1]) == "S" && string(grid[y+1][x-1]) == "M" {
						count++
					}
				}

				if string(grid[y-1][x-1]) == "S" && string(grid[y+1][x+1]) == "M" {
					if string(grid[y-1][x+1]) == "M" && string(grid[y+1][x-1]) == "S" {
						count++
					}
					if string(grid[y-1][x+1]) == "S" && string(grid[y+1][x-1]) == "M" {
						count++
					}
				}
			}
		}
	}
	fmt.Println("#2:", count)

	return nil
}
