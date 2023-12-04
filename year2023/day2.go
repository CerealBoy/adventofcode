package year2023

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

const (
	maximumRed   = 12
	maximumGreen = 13
	maximumBlue  = 14
)

func init() {
	days[2] = day2
}

type day2game struct {
	ID       int
	Rounds   []day2round
	Possible bool
}

type day2round struct {
	Red   int
	Green int
	Blue  int
}

func (g *day2game) Power() int {
	minGreen, minBlue, minRed := 0, 0, 0

	for _, x := range g.Rounds {
		if x.Red > minRed {
			minRed = x.Red
		}
		if x.Green > minGreen {
			minGreen = x.Green
		}
		if x.Blue > minBlue {
			minBlue = x.Blue
		}
	}

	return minRed * minBlue * minGreen
}

func day2(ctx *shared.Context) error {
	games := make(map[int]*day2game, 0)
	re := regexp.MustCompile(`^Game ([0-9]+): (.*)$`)
	reg := regexp.MustCompile(`([0-9]+) ([a-z]+),?`)

	f, _ := os.Open(shared.File(2023, 2, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		m := re.FindStringSubmatch(line)
		idx, _ := strconv.Atoi(m[1])
		rounds := strings.Split(m[2], ";")

		games[idx] = &day2game{ID: idx, Rounds: make([]day2round, 0), Possible: true}

		for _, y := range rounds {
			x := reg.FindAllStringSubmatch(y, -1)

			red, blue, green := 0, 0, 0

			for _, item := range x {
				val, _ := strconv.Atoi(item[1])
				switch item[2] {
				case "red":
					if val > maximumRed {
						games[idx].Possible = false
					}
					red = val
				case "green":
					if val > maximumGreen {
						games[idx].Possible = false
					}
					green = val
				case "blue":
					if val > maximumBlue {
						games[idx].Possible = false
					}
					blue = val
				default:
					panic("no")
				}
			}

			games[idx].Rounds = append(games[idx].Rounds, day2round{Red: red, Green: green, Blue: blue})
		}
	}

	sum := 0
	for _, x := range games {
		if x.Possible {
			sum += x.ID
		}
	}
	fmt.Println("#1:", sum)

	sum = 0
	for _, x := range games {
		sum += x.Power()
	}
	fmt.Println("#2:", sum)

	return nil
}
