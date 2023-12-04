package main

import (
	"github.com/CerealBoy/adventofcode/shared"
	"github.com/CerealBoy/adventofcode/year2015"
	"github.com/CerealBoy/adventofcode/year2016"
	"github.com/CerealBoy/adventofcode/year2017"
	"github.com/CerealBoy/adventofcode/year2018"
	"github.com/CerealBoy/adventofcode/year2019"
	"github.com/CerealBoy/adventofcode/year2020"
	"github.com/CerealBoy/adventofcode/year2021"
	"github.com/CerealBoy/adventofcode/year2022"
	"github.com/CerealBoy/adventofcode/year2023"
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
		Year2021 year2021.Cmd `cmd help:"Year 2021 for the question."`
		Year2022 year2022.Cmd `cmd help:"Year 2022 for the question."`
		Year2023 year2023.Cmd `cmd help:"Year 2023 for the question."`

		Debug bool `help:"Enable debug mode." short:"d"`
		Test  bool `help:"Whether the test should be used." short:"t"`
	}
)

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&shared.Context{Debug: cli.Debug, Test: cli.Test})
	ctx.FatalIfErrorf(err)
}
