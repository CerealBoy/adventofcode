package year2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[10] = day10
}

func CapNoop(c, v int, n map[int]int) {
	if (c-20)%40 == 0 {
		n[c] = v * c
	}
}

func CapAdd(c, v int, n map[int]int) {
	if (c-20)%40 == 0 {
		n[c] = v * c
	} else if (c-20)%40 == 1 {
		n[c-1] = v * (c - 1)
	}
}

func day10(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 10, ctx.Test))
	if err != nil {
		return err
	}

	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	cycle := 0
	value := 1
	notes := make(map[int]int, 0)
	for _, l := range lines {
		b := strings.Split(l, " ")
		if b[0] == "noop" {
			cycle++
			CapNoop(cycle, value, notes)

		} else {
			x, _ := strconv.Atoi(b[1])
			cycle += 2
			CapAdd(cycle, value, notes)
			value += x
		}
	}

	total := 0
	for _, x := range notes {
		total += x
	}
	fmt.Printf("#1: %d\n", total)

	return nil
}
