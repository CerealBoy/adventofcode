package year2020

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

type ins struct {
	n   string
	up  bool
	val int
}

var (
	re8 = regexp.MustCompile(`(jmp|acc|nop) ([+-])(\d+)$`)
)

func init() {
	days[8] = day8
}

func day8(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2020, 8, ctx.Test))
	if err != nil {
		return err
	}

	second := 0
	instructions := make(map[int]ins, 0)
	i := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		x := re8.FindStringSubmatch(t)

		n, _ := strconv.Atoi(x[3])
		instructions[i] = ins{n: x[1], up: x[2] == "+", val: n}
		i++
	}

	first, _ := run(instructions)

	for idx := 0; idx < len(instructions); idx++ {
		str := instructions[idx]

		fmt.Println(idx)

		switch str.n {
		case "jmp":
			str.n = "nop"
			instructions[idx] = str

			acc, err := run(instructions)
			if err == nil {
				second = acc
				goto completed
			}

			str.n = "jmp"
			instructions[idx] = str

		case "nop":
			str.n = "jmp"
			instructions[idx] = str
			acc, err := run(instructions)
			if err == nil {
				second = acc
				goto completed
			}

			str.n = "nop"
			instructions[idx] = str

		case "acc":
			continue

		}
	}

completed:

	fmt.Println("#1:", first, "\n#2:", second)
	return nil
}

func run(inst map[int]ins) (acc int, err error) {
	i := 0
	used := make(map[int]struct{}, 0)

	for {
		if _, ok := inst[i]; !ok {
			break
		}

		switch inst[i].n {
		case "jmp":
			used[i] = struct{}{}
			if inst[i].up {
				i += inst[i].val
				if _, ok := used[i]; ok {
					err = errors.New("yupp")
					return
				}
			} else {
				i -= inst[i].val
				if _, ok := used[i]; ok {
					err = errors.New("yupp")
					return
				}
			}

		case "acc":
			if inst[i].up {
				acc += inst[i].val
			} else {
				acc -= inst[i].val
			}
			used[i] = struct{}{}
			i++

		case "nop":
			used[i] = struct{}{}
			i++

		}
	}

	return
}
