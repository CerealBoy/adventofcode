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
	days[5] = day5
}

func day5(ctx *shared.Context) error {
	debug = ctx.Debug
	rules, pages, switched := [][]int{}, [][]int{}, false

	f, _ := os.Open(shared.File(2024, 5, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 1 {
			switched = true
			continue
		}

		if !switched {
			p := strings.Split(line, "|")
			l, _ := strconv.Atoi(p[0])
			r, _ := strconv.Atoi(p[1])
			rules = append(rules, []int{l, r})
		} else {
			page := []int{}
			for _, x := range strings.Split(line, ",") {
				i, _ := strconv.Atoi(x)
				page = append(page, i)
			}
			pages = append(pages, page)
		}
	}

	if debug {
		fmt.Println("Have", rules, "and then", pages)
	}

	first, second := 0, 0
	for _, p := range pages {
		isValid := true

		// determine if it's valid based on the rules
		for _, r := range rules {
			if !day5isValid(p, r) {
				if debug {
					fmt.Println("Invalid", p, "with rule", r)
				}

				isValid = false
				break
			}
		}

		if isValid {
			// find the middle and sum
			if debug {
				fmt.Println("Page", p, "is valid")
			}

			first += p[(len(p)-1)/2]
		}
	}
	fmt.Println("#1:", first)

	for _, p := range pages {
		needed := false

	redo:
		for _, r := range rules {
			if !day5isValid(p, r) {
				if debug {
					fmt.Println("Changing", p, "based on broken rule", r)
				}

				// shift the page slice around and redo
				p = day5changePage(p, r)
				needed = true
				goto redo
			}
		}

		if needed {
			second += p[(len(p)-1)/2]
		}
	}
	fmt.Println("#2:", second)

	return nil
}

func day5changePage(p, r []int) []int {
	a, b := 0, 0
	for idx, x := range p {
		switch x {
		case r[1]:
			a = idx
		case r[0]:
			b = idx
		}
	}

	if debug {
		fmt.Println("A", p[:a], "B", p[a+1:b], "C", r, "D", p[b+1:])
	}

	q := append([]int{}, p[:a]...)
	q = append(q, p[a+1:b]...)
	q = append(q, r...)
	q = append(q, p[b+1:]...)

	return q
}

func day5isValid(p, r []int) bool {
	seen, found := false, false

	for _, x := range p {
		if x == r[1] {
			if found {
				return true
			}

			seen = true
			continue
		}

		if x == r[0] {
			if seen {
				return false
			}

			found = true
			continue
		}
	}

	if !found || !seen {
		return true
	}
	return false
}
