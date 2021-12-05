package year2021

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type Board struct {
	Rows   [][]int
	Marked [][]bool
}

func (b *Board) Take(r []string) {
	row := []int{}
	for _, x := range r {
		i, _ := strconv.Atoi(x)
		row = append(row, i)
	}
	b.Rows = append(b.Rows, row)
	b.Marked = append(b.Marked, make([]bool, 5))
}

func (b *Board) Mark(n int) {
	for a, x := range b.Rows {
		for c, z := range x {
			if z == n {
				b.Marked[a][c] = true
				return
			}
		}
	}
}

func (b *Board) Bingo() bool {
	// rows
	for _, x := range b.Marked {
		count := 0
		for _, y := range x {
			if !y {
				break
			} else {
				count++
			}

			if count == 5 {
				return true
			}
		}
	}

	// columns
	for x := 0; x < 5; x++ {
		count := 0
		for y := 0; y < 5; y++ {
			if !b.Marked[y][x] {
				break
			} else {
				count++
			}

			if count == 5 {
				return true
			}
		}
	}

	return false
}

func (b *Board) Score(prev int) int {
	total := 0
	for x, a := range b.Marked {
		for y, c := range a {
			if !c {
				total += b.Rows[x][y]
			}
		}
	}

	return total * prev
}

func init() {
	days[4] = day4
}

func day4(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2021, 4, ctx.Test))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	numbers := scanner.Text()
	scanner.Scan()

	boards := []*Board{}
	board := &Board{}
	for scanner.Scan() {
		l := strings.ReplaceAll(strings.Trim(scanner.Text(), " "), "  ", " ")
		if len(l) < 1 {
			boards = append(boards, board)
			board = &Board{}
			continue
		}

		board.Take(strings.Split(l, " "))
	}
	boards = append(boards, board)

	for _, x := range strings.Split(numbers, ",") {
		n, _ := strconv.Atoi(x)
		for _, board := range boards {
			board.Mark(n)
			if board.Bingo() {
				fmt.Printf("#1: %d\n", board.Score(n))
				goto two
			}
		}
	}

two:
	newB := []*Board{}
	for _, x := range strings.Split(numbers, ",") {
		n, _ := strconv.Atoi(x)

		board := &Board{}
		for _, board = range boards {
			board.Mark(n)
			if !board.Bingo() {
				newB = append(newB, board)
			}
		}

		if len(newB) < 1 {
			fmt.Printf("#2: %d\n", board.Score(n))
			return nil
		}

		boards = newB
		newB = []*Board{}
	}

	return nil
}
