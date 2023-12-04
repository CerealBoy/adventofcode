package year2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

func IsVisible(x, y, v int, trees map[int]map[int]int) bool {
	// top
	for a := 0; a < x; a++ {
		if trees[a][y] >= v {
			goto down
		}
	}
	return true

down:
	for a := x + 1; a < len(trees); a++ {
		if trees[a][y] >= v {
			goto left
		}
	}
	return true

left:
	for a := 0; a < y; a++ {
		if trees[x][a] >= v {
			goto right
		}
	}
	return true

right:
	for a := y + 1; a < len(trees); a++ {
		if trees[x][a] >= v {
			return false
		}
	}

	return true
}

func ScenicScore(x, y, v int, trees map[int]map[int]int) int {
	count := 0
	for a := x - 1; a >= 0; a-- {
		count++
		if trees[a][y] >= v {
			break
		}
	}
	score := count

	count = 0
	for a := x + 1; a < len(trees); a++ {
		count++
		if trees[a][y] >= v {
			break
		}
	}
	score *= count

	count = 0
	for a := y - 1; a >= 0; a-- {
		count++
		if trees[x][a] >= v {
			break
		}
	}
	score *= count

	count = 0
	for a := y + 1; a < len(trees); a++ {
		count++
		if trees[x][a] >= v {
			break
		}
	}
	score *= count

	return score
}

func init() {
	days[8] = day8
}

func day8(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 8, ctx.Test))
	if err != nil {
		return err
	}

	lines := []string{}
	trees := make(map[int]map[int]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for x, line := range lines {
		for y, letter := range strings.Split(line, "") {
			if _, ok := trees[x]; !ok {
				trees[x] = make(map[int]int, 0)
			}

			tree, _ := strconv.Atoi(letter)
			trees[x][y] = tree
		}
	}

	visible := len(trees)*4 - 4
	for x := 1; x < len(trees)-1; x++ {
		for y := 1; y < len(trees)-1; y++ {
			if IsVisible(x, y, trees[x][y], trees) {
				visible++
			}
		}
	}
	fmt.Printf("#1: %d\n", visible)

	highest := 0
	for x := 1; x < len(trees)-1; x++ {
		for y := 1; y < len(trees)-1; y++ {
			score := ScenicScore(x, y, trees[x][y], trees)
			if score > highest {
				highest = score
			}
		}
	}
	fmt.Printf("#2: %d\n", highest)

	return nil
}
