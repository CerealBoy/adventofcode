package year2025

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[8] = day8
}

type point struct {
	x int
	y int
	z int
	id int
}

type pair struct {
	a point
	b point
	d float64
}

func day8(ctx *shared.Context) error {
	debug := ctx.Debug
	_ = debug

	f, _ := os.Open(shared.File(2025, 8, ctx.Test))
	scanner := bufio.NewScanner(f)
	points, id := []point{}, 0
	for scanner.Scan() {
		line := scanner.Text()
		x := strings.Split(line, ",")
		a, _ := strconv.Atoi(x[0])
		b, _ := strconv.Atoi(x[1])
		c, _ := strconv.Atoi(x[2])

		points = append(points, point{x: a, y: b, z: c, id: id})
		id++
	}

	connections := []pair{}
	for a, x := range points {
		for b, y := range points {
			if b <= a {
				continue
			}

			distance := math.Sqrt(math.Pow(float64(x.x - y.x), 2) + math.Pow(float64(x.y - y.y), 2) + math.Pow(float64(x.z - y.z), 2))
			connections = append(connections, pair{a: x, b: y, d: distance})
		}
	}

	slices.SortFunc(connections, func(a, b pair) int {
		if a.d < b.d {
			return -1
		}
		return 1
	})

	limit := 1000
	if ctx.Test {
		limit = 10
	}

	all, found, foundIdx := [][]int{}, false, 0
	for x := 0; x < limit; x++ {
		out:
		for idx, item := range all {
			for _, i := range item {
				if i == connections[x].a.id || i == connections[x].b.id {
					if found && idx != foundIdx {
						// join the arrays
						all[foundIdx] = append(all[foundIdx], all[idx]...)
						slices.Sort(all[foundIdx])
						all[foundIdx] = slices.Compact(all[foundIdx])
						all = append(all[0:idx], all[idx+1:]...)
						break out
					}

					all[idx] = append(all[idx], connections[x].a.id, connections[x].b.id)
					slices.Sort(all[idx])
					all[idx] = slices.Compact(all[idx])
					found = true
					foundIdx = idx
				}
			}
		}

		if !found {
			all = append(all, []int{connections[x].a.id, connections[x].b.id})
		}
		found = false
	}

	slices.SortFunc(all, func(a, b []int) int {
		if len(a) > len(b) {
			return -1
		}
		return 1
	})

	println("Part 1:", len(all[0]) * len(all[1]) * len(all[2]))

	all, found, foundIdx = [][]int{}, false, 0
	for x, conn := range connections {
		two:
		for idx, item := range all {
			for _, i := range item {
				if i == conn.a.id || i == conn.b.id {
					if found && idx != foundIdx {
						all[foundIdx] = append(all[foundIdx], all[idx]...)
						slices.Sort(all[foundIdx])
						all[foundIdx] = slices.Compact(all[foundIdx])
						all = append(all[0:idx], all[idx+1:]...)
						break two
					}

					all[idx] = append(all[idx], connections[x].a.id, connections[x].b.id)
					slices.Sort(all[idx])
					all[idx] = slices.Compact(all[idx])
					found = true
					foundIdx = idx
				}
			}
		}

		if !found {
			all = append(all, []int{connections[x].a.id, connections[x].b.id})
		}
		found = false

		if len(all) == 1 && len(all[0]) == len(points) {
			println("Part 2:", conn.a.x * conn.b.x)
			break
		}
	}

	return nil
}
