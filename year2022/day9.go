package year2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type Point struct {
	X int
	Y int
}

func AwayFrom(a, b *Point) bool {
	x := a.X - b.X
	if x < -1 || x > 1 {
		return true
	}
	y := a.Y - b.Y
	if y < -1 || y > 1 {
		return true
	}
	return false
}

func Move(from, to *Point) {
	if from.X == to.X {
		// moving along Y
		if from.Y > to.Y {
			from.Y--
		} else {
			from.Y++
		}

	} else if from.Y == to.Y {
		// moving along X
		if from.X > to.X {
			from.X--
		} else {
			from.X++
		}

	} else {
		// moving diagonally
		if from.X-to.X > 0 {
			from.X--
		} else {
			from.X++
		}

		if from.Y-to.Y > 0 {
			from.Y--
		} else {
			from.Y++
		}

	}
}

func init() {
	days[9] = day9
}

func day9(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 9, ctx.Test))
	if err != nil {
		return err
	}

	coords := make(map[string]struct{}, 0)
	steps := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		steps = append(steps, scanner.Text())
	}

	coords["0,0"] = struct{}{}
	points := make(map[int]*Point, 10)
	for p := 0; p < 10; p++ {
		points[p] = &Point{X: 0, Y: 0}
	}

	ninth := make(map[string]struct{}, 0)
	ninth["0,0"] = struct{}{}
	for _, x := range steps {
		step := strings.Split(x, " ")
		length, _ := strconv.Atoi(step[1])

		for i := 0; i < length; i++ {
			switch step[0] {
			case "R":
				points[0].X++
			case "U":
				points[0].Y++
			case "L":
				points[0].X--
			case "D":
				points[0].Y--
			}

			for a := 1; a < 10; a++ {
				if AwayFrom(points[a-1], points[a]) {
					Move(points[a], points[a-1])
				}
			}

			coords[fmt.Sprintf("%d,%d", points[1].X, points[1].Y)] = struct{}{}
			ninth[fmt.Sprintf("%d,%d", points[9].X, points[9].Y)] = struct{}{}
		}
	}

	fmt.Printf("#1: %d\n", len(coords))
	fmt.Printf("#2: %d\n", len(ninth))

	return nil
}
