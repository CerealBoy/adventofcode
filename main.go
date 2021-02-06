package main

import (
	"github.com/CerealBoy/adventofcode/shared"
	"github.com/CerealBoy/adventofcode/year2015"
	"github.com/CerealBoy/adventofcode/year2016"
	"github.com/CerealBoy/adventofcode/year2017"
	"github.com/CerealBoy/adventofcode/year2018"
	"github.com/CerealBoy/adventofcode/year2019"
	"github.com/CerealBoy/adventofcode/year2020"
	"github.com/alecthomas/kong"
)

var (
	cli struct {
		Year2015 year2015.Cmd `cmd help:"Year 2015 for the question."`
		Year2016 year2016.Cmd `cmd help:"Year 2016 for the question."`
		Year2017 year2017.Cmd `cmd help:"Year 2017 for the question."`
		Year2018 year2018.Cmd `cmd help:"Year 2018 for the question."`
		Year2019 year2019.Cmd `cmd help:"Year 2019 for the question."`
		Year2020 year2020.Cmd `cmd help:"Year 2020 for the question."`

		Input string `required help:"The input file to run." type:"existingfile"`

		Debug bool `help:"Enable debug mode."`
	}
)

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&shared.Context{Debug: cli.Debug, Input: cli.Input})
	ctx.FatalIfErrorf(err)
}
