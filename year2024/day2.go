package year2024

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
	if ctx.Debug {
		debug = ctx.Debug
	}

	reports := [][]int{}

	f, _ := os.Open(shared.File(2024, 2, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		report := []int{}
		v := strings.Split(scanner.Text(), " ")

		for _, x := range v {
			p, _ := strconv.Atoi(x)
			report = append(report, p)
		}

		reports = append(reports, report)
	}

	count := 0
	for _, x := range reports {
		if day2valid(x) {
			count++
		}
	}
	fmt.Println("Part #1:", count)

	count = 0
	for _, x := range reports {
		if day2moreValid(x, false) {
			count++
		}
	}
	fmt.Println("Part #2:", count)

	return nil
}

func day2valid(x []int) bool {
	asc := false
	if x[0] < x[1] && x[1] < x[2] {
		asc = true
	} else if x[0] > x[1] && x[1] > x[2] {
		asc = false
	} else {
		return false
	}

	for a := 0; a < len(x)-1; a++ {
		if asc {
			if x[a] >= x[a+1] {
				return false
			}
			if x[a+1]-x[a] > 3 || x[a+1]-x[a] < 1 {
				return false
			}

		} else {
			if x[a+1] >= x[a] {
				return false
			}
			if x[a]-x[a+1] > 3 || x[a]-x[a+1] < 1 {
				return false
			}

		}
	}

	return true
}

func day2moreValid(x []int, fail bool) bool {
	asc := false
	if x[0] < x[1] && x[1] < x[2] {
		asc = true
	}

	for a := 0; a < len(x)-1; a++ {
		if !day2cmp(asc, x[a], x[a+1]) {
			if fail {
				return false
			}

			// reconstruct and try again
			n := append([]int{}, x[:a]...)
			n = append(n, x[a+1:]...)
			if day2moreValid(n, true) {
				if debug {
					fmt.Println("Got it from", x, "with", n)
				}
				return true
			}

			m := append([]int{}, x[:a+1]...)
			m = append(m, x[a+2:]...)
			if day2moreValid(m, true) {
				if debug {
					fmt.Println("Got it from", x, "with alternate", m)
				}
				return true
			}

			l := []int{}
			if a > 0 {
				l = append([]int{}, x[:a-1]...)
				l = append(l, x[a:]...)
				if day2moreValid(l, true) {
					if debug {
						fmt.Println("Got it from", x, "with final", l)
					}
					return true
				}
			}

			if debug {
				fmt.Println("No good on", x, "with", n, "or", m, "even", l)
			}
			return false
		}
	}

	return true
}

func day2cmp(a bool, x, y int) bool {
	if a {
		if x >= y {
			return false
		}
		if y-x > 3 || y-x < 1 {
			return false
		}

	} else {
		if y >= x {
			return false
		}
		if x-y > 3 || x-y < 1 {
			return false
		}

	}

	return true
}
