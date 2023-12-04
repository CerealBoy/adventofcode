package year2022

import (
	"bufio"
	"os"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[20] = day20
}

func day20(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 20, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//
	}

	return nil
}
