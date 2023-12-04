package year2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type Stack struct {
	Items []string
}

func (s *Stack) Push(i string) {
	s.Items = append(s.Items, i)
}

func (s *Stack) Pop() string {
	if len(s.Items) < 1 {
		fmt.Println("Unable to pop!")
	}

	el := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]

	return el
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.Items)
}

/**

**/

func init() {
	days[5] = day5
}

func day5(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 5, ctx.Test))
	if err != nil {
		return err
	}

	stack := make(map[int]*Stack, 0)
	secondStack := make(map[int]*Stack, 0)
	steps := []string{}
	stacks := []string{}
	isSteps := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isSteps = true
			continue
		}

		if isSteps {
			steps = append(steps, scanner.Text())
		} else {
			stacks = append(stacks, line)
		}
	}

	stackCount := (len(stacks[len(stacks)-1]) + 1) / 4
	for x := 0; x < stackCount; x++ {
		stack[x+1] = new(Stack)
		secondStack[x+1] = new(Stack)

		// push the non-empty entries from stacks into stack
		offset := 1 + (4 * x)
		for y := len(stacks) - 2; y >= 0; y-- {
			if string(stacks[y][offset]) != " " {
				stack[x+1].Push(string(stacks[y][offset]))
				secondStack[x+1].Push(string(stacks[y][offset]))
			} else {
				goto next
			}
		}
	next:
	}

	for _, cmd := range steps {
		// parse
		bits := strings.Split(cmd, " ")

		count, _ := strconv.Atoi(bits[1])
		from, _ := strconv.Atoi(bits[3])
		to, _ := strconv.Atoi(bits[5])

		for x := 0; x < count; x++ {
			popped := stack[from].Pop()
			stack[to].Push(popped)
		}
	}

	out := ""
	for x := 1; x <= stackCount; x++ {
		out += stack[x].Pop()
	}
	fmt.Printf("#1: %s\n", out)

	for _, cmd := range steps {
		// parse
		bits := strings.Split(cmd, " ")

		count, _ := strconv.Atoi(bits[1])
		from, _ := strconv.Atoi(bits[3])
		to, _ := strconv.Atoi(bits[5])

		tmp := &Stack{}
		for x := 0; x < count; x++ {
			popped := secondStack[from].Pop()
			tmp.Push(popped)
		}

		for x := 0; x < count; x++ {
			secondStack[to].Push(tmp.Pop())
		}
	}

	out = ""
	for x := 1; x <= stackCount; x++ {
		out += secondStack[x].Pop()
	}
	fmt.Printf("#2: %s\n", out)

	return nil
}
