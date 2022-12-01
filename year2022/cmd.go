package year2022

import (
	"errors"

	"github.com/CerealBoy/adventofcode/shared"
)

var (
	days = map[int]func(*shared.Context) error{}
)

// Cmd defines the current year as a command, with any arguments.
type Cmd struct {
	Day int `help:"Day of the year to run."`
}

// Run will accept the context from the command and look to run the
// given day with it, returning an error from the execution.
func (c *Cmd) Run(ctx *shared.Context) error {
	f, ok := days[c.Day]
	if !ok {
		return errors.New(shared.InvalidDay)
	}

	return f(ctx)
}
