// Package shared contains any resources that are reusable across each
// year and day, simplifying the code required and keeping each day
// focused on solving the actual problem at hand.
package shared

import (
	"fmt"
)

const (
	InvalidDay = "an invalid day was referenced for the year"
)

// Context procides a general and shared context to be used between all commands.
type Context struct {
	// Debug will define if additional output should be generated.
	Debug bool

	// Test will track whether the test or actual file should be used as input for the day.
	Test bool
}

func File(y, d int, t bool) string {
	o := "in"
	if t {
		o = "test"
	}
	return fmt.Sprintf(
		"year%d/day%d-%s",
		y,
		d,
		o,
	)
}
