package year2020

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[5] = day5
}

func day5(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2020, 5, ctx.Test))
	if err != nil {
		return err
	}

	highest := 0
	seats := make(map[int]map[int]bool, 128)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		r := t[0:7]
		s := t[7:10]

		row := bin2dec(r, "B", 7)
		seat := bin2dec(s, "R", 3)

		x := (row * 8) + seat
		if x > highest {
			highest = x
		}

		if seats[row] == nil {
			seats[row] = make(map[int]bool, 128)
		}
		seats[row][seat] = true
	}

	missing := 0
	for y, x := range seats {
		if len(x) == 7 {
			for a := 0; a < 8; a++ {
				if _, ok := x[a]; !ok {
					missing = (y * 8) + a
				}
			}
		}
	}

	fmt.Println("#1:", highest, "\n#2:", missing)
	return nil
}

func bin2dec(s, e string, a int) (t int) {
	for _, x := range strings.Split(s, "") {
		if x == e {
			t += (int(math.Pow(2, float64(a))) / 2)
		}
		a--
	}

	return
}
