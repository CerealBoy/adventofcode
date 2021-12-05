package year2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

const (
	size = 1000
)

type Point struct {
	X int
	Y int
}

func P(line string) *Point {
	l := strings.Split(line, ",")
	a, _ := strconv.Atoi(l[0])
	b, _ := strconv.Atoi(l[1])
	return &Point{
		X: a,
		Y: b,
	}
}

type Line struct {
	A *Point
	B *Point
}

func init() {
	days[5] = day5
}

func day5(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2021, 5, ctx.Test))
	if err != nil {
		return err
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// part 1
	grid := [size][size]int{}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			grid[x][y] = 0
		}
	}

	for _, x := range lines {
		c := strings.Split(x, " -> ")
		a := P(c[0])
		b := P(c[1])

		if Straight(a, b) {
			grid = Inc(grid, a, b)
		}
	}

	fmt.Printf("#1: %d\n", Count(grid))

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			grid[x][y] = 0
		}
	}

	for _, x := range lines {
		c := strings.Split(x, " -> ")
		a := P(c[0])
		b := P(c[1])

		grid = Inc(grid, a, b)
	}

	fmt.Printf("#2: %d\n", Count(grid))

	return nil
}

func Count(grid [size][size]int) int {
	count := 0
	for _, x := range grid {
		for _, y := range x {
			if y > 1 {
				count++
			}
		}
	}

	return count
}

func Inc(grid [size][size]int, a, b *Point) [size][size]int {
	if a.X == b.X {
		if a.Y > b.Y {
			for x := b.Y; x <= a.Y; x++ {
				grid[x][a.X]++
			}
		} else {
			for x := a.Y; x <= b.Y; x++ {
				grid[x][a.X]++
			}
		}
	} else if a.Y == b.Y {
		if a.X > b.X {
			for x := b.X; x <= a.X; x++ {
				grid[a.Y][x]++
			}
		} else {
			for x := a.X; x <= b.X; x++ {
				grid[a.Y][x]++
			}
		}
	} else {
		if a.X > b.X {
			if a.Y > b.Y {
				for x := 0; x <= (a.Y - b.Y); x++ {
					grid[b.Y+x][b.X+x]++
				}
			} else {
				for x := 0; x <= (b.Y - a.Y); x++ {
					grid[b.Y-x][b.X+x]++
				}
			}
		} else {
			if a.Y > b.Y {
				for x := 0; x <= (a.Y - b.Y); x++ {
					grid[a.Y-x][a.X+x]++
				}
			} else {
				for x := 0; x <= (b.Y - a.Y); x++ {
					grid[a.Y+x][a.X+x]++
				}
			}
		}
	}

	return grid
}

func Print(grid [size][size]int) {
	for _, x := range grid {
		for _, y := range x {
			if y == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%d", y)
			}
		}
		fmt.Println("")
	}
}

func Straight(a, b *Point) bool {
	if a.X == b.X || a.Y == b.Y {
		return true
	}

	return false
}
