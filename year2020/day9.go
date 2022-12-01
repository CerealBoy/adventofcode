package year2020

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[9] = day9
}

func day9(ctx *shared.Context) error {
	numbers := []int{}

	f, err := os.Open(shared.File(2020, 9, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, n)
	}

	weak := 0
	length := 25

	for i, j := range numbers {
		if i < length {
			continue
		}

		found := false
		// j is the current number

	next:
		// go through the next length numbers
		for a := i - length; a < i; a++ {
			first := numbers[a]
			// if the current number is already larger, skip it
			if first > j {
				continue
			}

			for b := a + 1; b < i; b++ {
				second := numbers[b]
				if first+second == j {
					found = true
					break next
				}
			}
		}

		if !found {
			weak = j
			fmt.Printf("#1: %v\n", weak)
			break
		}
	}

	// now we need to look for a contiguous set of numbers that add to weak
	for i, j := range numbers {
		total := j
		min := j
		max := j

		for _, b := range numbers[i+1:] {
			total += b
			if total > weak {
				break
			}
			if min > b {
				min = b
			}
			if max < b {
				max = b
			}
			if total == weak {
				fmt.Printf("#2: %v\n", min+max)
			}
		}
	}

	return nil
}
