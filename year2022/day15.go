package year2022

import (
	"bufio"
	"os"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[15] = day15
}

func day15(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 15, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//
	}

	return nil
}
