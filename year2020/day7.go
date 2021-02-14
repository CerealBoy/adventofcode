package year2020

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type bag struct {
	name string
	list map[string]int
}

var (
	re = regexp.MustCompile(`(\d+) ([a-zA-Z ]+) bags?[.,]`)
)

func init() {
	days[7] = day7
}

func day7(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2020, 7, ctx.Test))
	if err != nil {
		return err
	}

	bags := make(map[string]bag, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		d := strings.Split(t, " bags contain ")
		m := re.FindAllStringSubmatch(d[1], -1)
		b := bag{name: d[0], list: make(map[string]int, 0)}

		for _, x := range m {
			n, _ := strconv.Atoi(x[1])
			b.list[x[2]] = n
		}

		bags[d[0]] = b
	}

	contains := 0
	for name, _ := range bags {
		if con(bags, name) {
			contains++
		}
	}

	count := count(bags, "shiny gold")

	fmt.Println("#1:", contains, "\n#2:", count)
	return nil
}

func con(b map[string]bag, n string) bool {
	if b[n].list["shiny gold"] > 0 {
		return true
	}

	for cur, _ := range b[n].list {
		if con(b, cur) {
			return true
		}
	}

	return false
}

func count(b map[string]bag, n string) (t int) {
	for name, c := range b[n].list {
		t = t + c + (c * count(b, name))
	}

	return
}
