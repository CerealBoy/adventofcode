package year2023

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[4] = day4
}

type day4data struct {
	ID      int
	Winners []string
	Have    []string
	Count   int
}

func newDay4(l string) *day4data {
	a := strings.Split(strings.TrimPrefix(l, "Card "), ":")
	id, _ := strconv.Atoi(strings.TrimSpace(a[0]))

	b := strings.Split(a[1], " | ")

	winners := []string{}
	for _, x := range strings.Split(b[0], " ") {
		if len(x) > 0 && x != "" {
			winners = append(winners, x)
		}
	}

	have := []string{}
	for _, x := range strings.Split(b[1], " ") {
		if len(x) > 0 && x != "" {
			have = append(have, x)
		}
	}

	return &day4data{ID: id, Winners: winners, Have: have, Count: 0}
}

func (d *day4data) score() int {
	s := 0

	for _, x := range d.Winners {
		for _, y := range d.Have {
			if x == y {
				d.Count++
				if s == 0 {
					s = 1
				} else {
					s *= 2
				}
			}
		}
	}

	return s
}

func day4(ctx *shared.Context) error {
	total, last := 0, 0
	bonus := make(map[int]int, 0)

	f, _ := os.Open(shared.File(2023, 4, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		d := newDay4(line)
		total += d.score()
		if _, ok := bonus[d.ID]; !ok {
			bonus[d.ID] = 1
			if ctx.Debug {
				fmt.Println("Doing base for", d.ID)
			}
		} else {
			bonus[d.ID]++
			if ctx.Debug {
				fmt.Println("Increment", d.ID, "to", bonus[d.ID])
			}
		}

		if ctx.Debug {
			fmt.Println("For", d.ID, "there's a bonus of", bonus[d.ID])
		}

		// loop for each card we have of this ID
		for y := d.ID; y < d.ID+bonus[d.ID]; y++ {
			if ctx.Debug {
				fmt.Println("y is", y, "and bonus", bonus[d.ID], "increment the next", d.Count)
			}
			for x := d.ID + 1; x <= d.ID+d.Count; x++ {
				if _, ok := bonus[x]; !ok {
					bonus[x] = 0
				}
				bonus[x]++
			}
		}

		last = d.ID
	}
	fmt.Println("#1:", total)

	total = 0
	for x := 1; x <= last; x++ {
		total += bonus[x]
	}
	fmt.Println("#2:", total)

	return nil
}
