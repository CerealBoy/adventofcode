package year2022

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[6] = day6
}

func day6(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 6, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	chars := strings.Split(scanner.Text(), "")

	x := 0
	for x = 0; x < len(chars)-4; x++ {
		if chars[x] != chars[x+1] &&
			chars[x] != chars[x+2] &&
			chars[x] != chars[x+3] &&
			chars[x+1] != chars[x+2] &&
			chars[x+1] != chars[x+3] &&
			chars[x+2] != chars[x+3] {
			break
		}
	}
	fmt.Printf("#1: %d\n", x+4)

	pos := 0
	for pos = 0; pos < len(chars)-14; pos++ {
		found := make(map[string]int, 14)
		c := chars[pos : pos+14]

		for _, x := range c {
			if _, ok := found[x]; !ok {
				found[x] = 1
			} else {
				found[x]++
			}
		}

		broken := false
		for _, x := range found {
			if x > 1 {
				broken = true
			}
		}

		if !broken {
			break
		}
	}

	fmt.Printf("#2: %d\n", pos+14)

	return nil
}
