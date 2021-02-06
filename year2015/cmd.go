package year2015

import (
	"fmt"

	"github.com/CerealBoy/adventofcode/shared"
)

type Cmd struct {
	Day int `cmd help:"Day of the year to run."`
}

func (c *Cmd) Run(ctx *shared.Context) error {
	fmt.Printf("%#v\n", ctx)

	println("Hello, world")

	return nil
}
