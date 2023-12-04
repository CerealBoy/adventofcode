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
	days[3] = day3
}

func day3(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 3, ctx.Test))
	if err != nil {
		return err
	}

	lines := []string{}
	found := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rs := scanner.Text()
		lines = append(lines, rs)

		c1 := strings.Split(rs[0:len(rs)/2], "")
		c2 := strings.Split(rs[len(rs)/2:], "")

		for _, a := range c1 {
			for _, b := range c2 {
				if a == b {
					found = append(found, a)
					goto outside
				}
			}
		}
	outside:
	}

	total := 0
	for _, x := range found {
		if strings.ToUpper(x) == x {
			a, _ := strconv.ParseInt(fmt.Sprintf("%x", x), 16, 64)
			total += int(a - 38)
		} else {
			a, _ := strconv.ParseInt(fmt.Sprintf("%x", x), 16, 64)
			total += int(a - 96)
		}
	}

	fmt.Printf("#1: %d\n", total)

	found = []string{}
	total = 0
	for x := 0; x < len(lines); x += 3 {
		for _, a := range strings.Split(lines[x], "") {
			if strings.Index(lines[x+1], a) >= 0 && strings.Index(lines[x+2], a) >= 0 {
				found = append(found, a)

				if strings.ToUpper(a) == a {
					b, _ := strconv.ParseInt(fmt.Sprintf("%x", a), 16, 64)
					total += int(b - 38)
				} else {
					b, _ := strconv.ParseInt(fmt.Sprintf("%x", a), 16, 64)
					total += int(b - 96)
				}

				goto outer
			}
		}
	outer:
	}

	fmt.Printf("#2: %d\n", total)

	return nil
}
