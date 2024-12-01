package year2024

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type day1Slice []int

func (s day1Slice) Len() int {
	return len(s)
}

func (s day1Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s day1Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *day1Slice) Pop() any {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func (s *day1Slice) Push(x any) {
	*s = append(*s, x.(int))
}

func init() {
	days[1] = day1
}

func day1(ctx *shared.Context) error {
	left, right := &day1Slice{}, &day1Slice{}
	heap.Init(left)
	heap.Init(right)
	counter, similarity := make(map[int]int, 0), make(map[int]int, 0)

	f, _ := os.Open(shared.File(2024, 1, ctx.Test))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		before, after, _ := strings.Cut(scanner.Text(), "   ")
		a, _ := strconv.Atoi(before)
		b, _ := strconv.Atoi(after)

		heap.Push(left, a)
		heap.Push(right, b)

		counter[a]++
		similarity[b]++
	}

	distance := 0
	for left.Len() > 0 {
		a, b := heap.Pop(left).(int), heap.Pop(right).(int)
		if a < b {
			distance += b - a
		} else {
			distance += a - b
		}
	}

	fmt.Println("Total distance: ", distance)

	distance = 0
	for k, v := range counter {
		if val, found := similarity[k]; found {
			distance += (v * k * val)
		}
	}

	fmt.Println("Similarity: ", distance)

	return nil
}
