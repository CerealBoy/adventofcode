package year2021

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[3] = day3
}

func day3(ctx *shared.Context) error {
	lines := []string{}

	f, err := os.Open(shared.File(2021, 3, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	gamma := ""
	b := 0
	for y := 0; y < len(lines[0]); y++ {
		for _, x := range lines {
			if x[y] == '0' {
				b--
			} else {
				b++
			}
		}

		if b >= 0 {
			gamma += "1"
		} else {
			gamma += "0"
		}
		b = 0
	}

	epsilon := ""
	for _, x := range strings.Split(gamma, "") {
		if x == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}

	fmt.Printf("#1: %d\n", bin2dec(gamma)*bin2dec(epsilon))

	// part 2
	new_lines := lines
	oxygen := []string{}
	foundOxy := ""
	b = 0
	for y := 0; y < len(new_lines[0]); y++ {
		for _, x := range new_lines {
			if x[y] == '0' {
				b--
			} else {
				b++
			}
		}

		if b >= 0 {
			for _, z := range new_lines {
				if z[y] == '1' {
					oxygen = append(oxygen, z)
				}
			}
		} else {
			for _, z := range new_lines {
				if z[y] == '0' {
					oxygen = append(oxygen, z)
				}
			}
		}

		new_lines = oxygen
		oxygen = []string{}
		b = 0

		if len(new_lines) < 2 {
			foundOxy = new_lines[0]
			break
		}
	}

	new_lines = lines
	co2 := []string{}
	foundCo2 := ""
	b = 0
	for y := 0; y < len(new_lines[0]); y++ {
		for _, x := range new_lines {
			if x[y] == '0' {
				b--
			} else {
				b++
			}
		}

		if b >= 0 {
			for _, z := range new_lines {
				if z[y] == '0' {
					co2 = append(co2, z)
				}
			}
		} else {
			for _, z := range new_lines {
				if z[y] == '1' {
					co2 = append(co2, z)
				}
			}
		}

		new_lines = co2
		co2 = []string{}
		b = 0

		if len(new_lines) < 2 {
			foundCo2 = new_lines[0]
			break
		}
	}

	fmt.Printf("#2: %d\n", bin2dec(foundOxy)*bin2dec(foundCo2))
	return nil
}

func bin2dec(n string) int {
	sum := 0
	for x := 0; x < len(n); x++ {
		if n[len(n)-x-1] == '1' {
			sum += int(math.Pow(2, float64(x)))
		}
	}

	return sum
}
