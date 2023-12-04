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
	days[4] = day4
}

func day4(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 4, ctx.Test))
	if err != nil {
		return err
	}

	pairs := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		pairs = append(pairs, line)
	}

	total := 0
	atAll := 0
	for _, line := range pairs {
		elves := strings.Split(line, ",")
		left := strings.Split(elves[0], "-")
		right := strings.Split(elves[1], "-")

		// determine full containment
		leftLow, _ := strconv.Atoi(left[0])
		leftHigh, _ := strconv.Atoi(left[1])
		rightLow, _ := strconv.Atoi(right[0])
		rightHigh, _ := strconv.Atoi(right[1])

		if (leftLow <= rightLow && leftHigh >= rightHigh) ||
			(rightLow <= leftLow && rightHigh >= leftHigh) {
			total += 1
		}

		if (leftHigh >= rightLow && leftLow <= rightHigh) ||
			(rightHigh <= leftLow && rightLow >= leftHigh) {
			atAll += 1
		}
	}

	fmt.Printf("#1: %d\n", total)
	fmt.Printf("#2: %d\n", atAll)

	return nil
}
