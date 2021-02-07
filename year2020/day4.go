package year2020

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

var (
	fourColours = map[string]struct{}{
		"amb": struct{}{},
		"blu": struct{}{},
		"brn": struct{}{},
		"gry": struct{}{},
		"grn": struct{}{},
		"hzl": struct{}{},
		"oth": struct{}{},
	}
)

func init() {
	days[4] = day4
}

func day4(ctx *shared.Context) error {
	l := ""
	p := []string{}
	first := 0
	second := 0

	f, err := os.Open(ctx.Input)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) < 1 {
			l = strings.TrimSpace(l)
			p = append(p, l)

			l = ""
		} else {
			l = l + " " + t
		}
	}

	for _, x := range p {
		e := strings.Split(x, " ")
		c := 0
		d := 0

		for _, y := range e {
			s := strings.Split(y, ":")
			switch s[0] {
			case "byr":
				c++
				n, _ := strconv.Atoi(s[1])
				if n >= 1920 && n <= 2002 {
					d++
				}

			case "iyr":
				c++
				n, _ := strconv.Atoi(s[1])
				if n >= 2010 && n <= 2020 {
					d++
				}

			case "eyr":
				c++
				n, _ := strconv.Atoi(s[1])
				if n >= 2020 && n <= 2030 {
					d++
				}

			case "hgt":
				c++
				if strings.HasSuffix(s[1], "cm") {
					a := strings.TrimSuffix(s[1], "cm")
					n, _ := strconv.Atoi(a)
					if n >= 150 && n <= 193 {
						d++
					}
				} else {
					a := strings.TrimSuffix(s[1], "in")
					n, _ := strconv.Atoi(a)
					if n >= 59 && n <= 76 {
						d++
					}
				}

			case "hcl":
				c++
				re := regexp.MustCompile(`#[a-z0-9]{6}`)
				if re.MatchString(s[1]) {
					d++
				}

			case "ecl":
				c++
				if _, ok := fourColours[s[1]]; ok {
					d++
				}

			case "pid":
				c++
				if len(s[1]) == 9 {
					d++
				}
			}
		}

		if c == 7 {
			first++
		}
		if d == 7 {
			second++
		}
	}

	fmt.Println("#1:", first, "\n#2:", second)
	return nil
}
