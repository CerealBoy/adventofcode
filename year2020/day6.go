package year2020

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type person struct {
	answers map[string]struct{}
}

func init() {
	days[6] = day6
}

func day6(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2020, 6, ctx.Test))
	if err != nil {
		return err
	}

	group := make([]person, 0)
	groups := make([][]person, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		if len(t) > 0 {
			p := person{answers: make(map[string]struct{}, 26)}
			for _, x := range strings.Split(t, "") {
				p.answers[x] = struct{}{}
			}

			group = append(group, p)
		} else {
			groups = append(groups, group)
			group = make([]person, 0)
		}
	}

	one := 0
	for _, x := range groups {
		letters := make(map[string]struct{}, 0)
		for _, y := range x {
			for letter, _ := range y.answers {
				letters[letter] = struct{}{}
			}
		}
		one += len(letters)
	}

	two := 0
	for _, x := range groups {
		letters := make(map[string]int, 0)
		for _, y := range x {
			for letter, _ := range y.answers {
				letters[letter]++
			}
		}

		for _, letter := range letters {
			if letter == len(x) {
				two++
			}
		}
	}

	fmt.Println("#1:", one, "\n#2:", two)
	return nil
}
