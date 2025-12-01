// Package year2019 holds all the solutions for 2019.
package year2019

import (
	"errors"

	"github.com/CerealBoy/adventofcode/shared"
)

var (
	days  = map[int]func(*shared.Context) error{}
	debug = false
)

type Cmd struct {
	Day int `cmd help:"Day of the year to run."`
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
