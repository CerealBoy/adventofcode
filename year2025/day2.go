package year2025

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
	f, _ := os.Open(shared.File(2025, 2, ctx.Test))
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	text := scanner.Text()
	input := strings.Split(text, ",")

	total, part2 := 0, 0
	for _, line := range input {
		asStrings := strings.Split(line, "-")
		nums := [2]int{}

		nums[0], _ = strconv.Atoi(asStrings[0])
		nums[1], _ = strconv.Atoi(asStrings[1])

		// split the numbers in half and compare
		for num := nums[0]; num <= nums[1]; num++ {
			if day2Mirrored(num, 2, false) {
				total += num
			}
			if day2Mirrored(num, 2, true) {
				part2 += num
			}
		}
	}
	println("Part 1:", total)
	println("Part 2:", part2)

	return nil
}

func day2Mirrored(num, pieces int, recurse bool) bool {
	str := fmt.Sprintf("%d", num)

	// If the pieces divides the length
	// Grab the string and repeat it pieces times
	// Does it match num

	if len(str) < pieces {
		return false
	}

	if len(str) % pieces != 0 {
		if recurse {
			return day2Mirrored(num, pieces+1, recurse)
		}
		return false
	}

	find, validating := str[0:len(str)/pieces], ""
	for range pieces {
		validating += find
	}

	if string(str) == string(validating) {
		return true
	}

	if recurse {
		return day2Mirrored(num, pieces+1, recurse)
	}
	return false
}
