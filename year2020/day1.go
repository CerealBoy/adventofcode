package year2020

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[1] = day1
}

func day1(ctx *shared.Context) error {
	d := []int{}

	f, err := os.Open(shared.File(2020, 1, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d = append(d, i)
	}

	for _, x := range d {
		for _, y := range d {
			if x != y && x+y == 2020 {
				fmt.Printf("#1: %d\n", x*y)
				goto two
			}
		}
	}

two:
	for _, a := range d {
		for _, b := range d {
			if a == b {
				continue
			}

			for _, c := range d {
				if a == c || b == c {
					continue
				}

				if a+b+c == 2020 {
					fmt.Printf("#2: %d\n", a*b*c)
					return nil
				}
			}
		}
	}

	return nil
}
