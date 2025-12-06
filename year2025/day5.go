package year2025

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type Range struct {
	Lower, Upper int
}

func day5(ctx *shared.Context) error {
	debug := ctx.Debug
	_ = debug

	f, _ := os.Open(shared.File(2025, 5, ctx.Test))
	scanner := bufio.NewScanner(f)
	fresh, list, ing := [][]int{}, []int{}, false
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			ing = true
			continue
		}

		if !ing {
			bits := strings.Split(line, "-")
			low, _ := strconv.Atoi(bits[0])
			high, _ := strconv.Atoi(bits[1])

			fresh = append(fresh, []int{low, high})
		} else {
			num, _ := strconv.Atoi(line)
			list = append(list, num)
		}
	}

	part1, part2 := 0, 0

	for _, i := range list {
		if debug {
			println("Current ingredient is:", i)
		}

		for _, r := range fresh {
			if i >= r[0] && i <= r[1] {
				part1++
				break
			}
		}
	}

	println("Part 1:", part1)

	changed := true
	for changed {
		next := [][]int{}
		changed = false
		for i, r := range fresh {
			for j, o := range fresh {
				if i == j {
					continue
				}

				if r[0] > o[1] || r[1] < o[0] {
					// no potential overlap
					continue
				}

				if r[0] >= o[0] && r[1] <= o[1] {
					// Contained entirely.
					changed = true
					break
				}

				if r[0] <= o[0] && r[1] >= o[0] && r[1] <= o[1] {
					fresh[j][0] = r[0]
					changed = true
					break
				}

				if r[1] >= o[1] && r[0] >= o[0] && r[0] <= o[1] {
					fresh[j][1] = r[1]
					changed = true
					break
				}
			}

			if changed {
				for k, v := range fresh {
					if k == i {
						continue
					}

					next = append(next, v)
				}
				fresh = next
				break
			}
		}
	}

	for _, r := range fresh {
		part2 += 1 + r[1] - r[0]
	}

	println("Part 2:", part2)

	return nil
}
