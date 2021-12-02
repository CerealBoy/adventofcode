package year2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[2] = day2
}

func day2(ctx *shared.Context) error {
	lines := []string{}

	f, err := os.Open(shared.File(2021, 2, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// part 1
	hor, depth := 0, 0
	for _, x := range lines {
		bits := strings.Split(x, " ")
		i, _ := strconv.Atoi(bits[1])

		switch bits[0] {
		case "forward":
			hor += i

		case "down":
			depth += i

		case "up":
			depth -= i
		}
	}

	fmt.Printf("#1: %d\n", hor*depth)

	// part 2
	hor, depth, aim := 0, 0, 0
	for _, x := range lines {
		bits := strings.Split(x, " ")
		i, _ := strconv.Atoi(bits[1])

		switch bits[0] {
		case "forward":
			hor += i
			depth += (aim * i)

		case "down":
			aim += i

		case "up":
			aim -= i
		}
	}

	fmt.Printf("#2: %d\n", hor*depth)
	return nil
}
