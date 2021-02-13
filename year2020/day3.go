package year2020

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[3] = day3
}

func day3(ctx *shared.Context) error {
	hill := []string{}

	f, err := os.Open(shared.File(2020, 3, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	for scanner.Scan() {
		hill = append(hill, scanner.Text())
	}

	fmt.Printf("#1: %d\n", ski(hill, 3, 1))
	fmt.Printf("#2: %d\n", ski(hill, 1, 1)*ski(hill, 3, 1)*ski(hill, 5, 1)*ski(hill, 7, 1)*ski(hill, 1, 2))
	return nil
}

func ski(h []string, offset, speed int) (trees int) {
	trees = 0
	off := 0

	for y, x := range h {
		if (y+1)%speed != 0 {
			continue
		}

		off = (offset + off) % len(x)
		if x[off] == byte('#') {
			trees++
		}
	}

	return trees
}
