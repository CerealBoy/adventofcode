package year2021

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
	count := 0
	prev := 0

	f, err := os.Open(shared.File(2021, 1, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		if prev == 0 {
			prev = i
		} else {
			if i > prev {
				count++
			}
			prev = i
		}
	}

	fmt.Printf("#1: %d\n", count)

	prev, count = 0, 0
	a, b := 0, 0

	f, err = os.Open(shared.File(2021, 1, ctx.Test))
	if err != nil {
		return err
	}

	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}

		if a == 0 {
			a = i
			continue
		}

		if b == 0 {
			b = i
			continue
		}

		if prev == 0 {
			prev = a + b + i
			a = b
			b = i
			continue
		}

		if a+b+i > prev {
			count++
		}

		prev = a + b + i
		a = b
		b = i
	}

	fmt.Printf("#2: %d\n", count)
	return nil
}
