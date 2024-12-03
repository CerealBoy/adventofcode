package year2024

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[3] = day3
}

func day3(ctx *shared.Context) error {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	alt := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	count, altCount, on := 0, 0, true

	f, _ := os.Open(shared.File(2024, 3, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		// get all the matches out
		matches := re.FindAllStringSubmatch(line, -1)
		for _, x := range matches {
			a, _ := strconv.Atoi(x[1])
			b, _ := strconv.Atoi(x[2])

			count += a * b
		}

		altMatch := alt.FindAllStringSubmatch(line, -1)
		for _, x := range altMatch {
			if x[0] == "do()" {
				on = true
			} else if x[0] == "don't()" {
				on = false
			} else {
				if on {
					a, _ := strconv.Atoi(x[1])
					b, _ := strconv.Atoi(x[2])

					altCount += a * b
				}
			}
		}
	}

	fmt.Println("#1:", count)
	fmt.Println("#2:", altCount)

	return nil
}
