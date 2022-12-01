package year2022

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/CerealBoy/adventofcode/shared"
)

func init() {
	days[1] = day1
}

func day1(ctx *shared.Context) error {
	elf := 0
	elves := make(map[int]int, 0)
	elves[elf] = 0

	f, err := os.Open(shared.File(2022, 1, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			x, _ := strconv.Atoi(text)
			elves[elf] += x
		} else {
			elf++
			elves[elf] = 0
		}
	}

	list := []int{}
	for _, x := range elves {
		list = append(list, x)
	}
	sort.Ints(list)

	fmt.Printf("#1: %v\n", list[len(list)-1])

	fmt.Printf("#2: %v\n", list[len(list)-1]+list[len(list)-2]+list[len(list)-3])

	return nil
}
