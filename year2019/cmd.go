package year2019

import (
	"github.com/CerealBoy/adventofcode/shared"
)

type Cmd struct {
	Day int `cmd help:"Day of the year to run."`
}

func (c *Cmd) Run(ctx *shared.Context) error {
	println("Hello, world")

	return nil
}
