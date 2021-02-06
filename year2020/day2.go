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

func init() {
	days[2] = day2
}

func day2(ctx *shared.Context) error {
	re := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): (.+)$`)
	first := 0
	second := 0

	f, err := os.Open(ctx.Input)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())

		min, _ := strconv.Atoi(m[1])
		max, _ := strconv.Atoi(m[2])
		count := strings.Count(m[4], m[3])

		if count >= min && count <= max {
			first++
		}

		x := string(m[4][min-1]) == m[3]
		y := string(m[4][max-1]) == m[3]
		if (x || y) && !(x && y) {
			second++
		}
	}

	fmt.Printf("#1: %d\n#2: %d\n", first, second)
	return nil
}
