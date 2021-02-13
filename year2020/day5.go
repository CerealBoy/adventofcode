package year2020

import (
	"bufio"
	"fmt"
	"os"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[5] = day5
}

func day5(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2020, 5, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		_ = t
	}

	fmt.Println("#1:", "", "\n#2:", "")
	return nil
}
