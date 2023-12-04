package year2022

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CerealBoy/adventofcode/shared"
)

type File struct {
	Name     string
	Size     int
	Type     string
	Children map[string]*File
	Parent   *File
}

func (f *File) AddChild(n string, s int) {
	f.Children[n] = &File{Name: n, Size: s, Parent: f, Type: "f"}
}

func (f *File) AddDir(n string) {
	f.Children[n] = &File{Name: n, Children: make(map[string]*File, 0), Parent: f, Type: "d"}
}

func (f *File) Enter(n string) *File {
	return f.Children[n]
}

func (f *File) GoUp() *File {
	return f.Parent
}

func (f *File) GetSize() int {
	// recursively determine the size
	if f.Size > 0 {
		return f.Size
	}

	for _, x := range f.Children {
		f.Size += x.GetSize()
	}

	return f.Size
}

func (f *File) CountAllDUnder() int {
	count := 0

	for _, x := range f.Children {
		if x.Type == "d" {
			count += x.CountAllDUnder()
			if x.Size < 100000 {
				count += x.Size
			}
		}
	}

	return count
}

func (f *File) FindClosestLarger(s int) int {
	max := f.Size

	// look for all directories that are larger than 's', tracking that which is closest to 's'
	for _, x := range f.Children {
		if x.Type == "d" {
			if x.Size < s {
				continue
			}

			out := x.FindClosestLarger(s)
			if out < max {
				max = out
			} else if x.Size < max {
				max = x.Size
			}
		}
	}

	return max
}

const (
	maximumToConsider   = 100_000
	minimumForUpdate    = 30_000_000
	totalSpaceAvailable = 70_000_000
)

func init() {
	days[7] = day7
}

func day7(ctx *shared.Context) error {
	f, err := os.Open(shared.File(2022, 7, ctx.Test))
	if err != nil {
		return err
	}

	tree := map[string]*File{"/": &File{Name: "/", Children: make(map[string]*File, 0), Parent: nil}}
	in := []string{}
	node := tree["/"]

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}

	for _, x := range in {
		tokens := strings.Split(x, " ")

		switch tokens[0] {
		case "$":
			if tokens[1] == "cd" {
				if tokens[2] == "/" {
					continue
				}

				if tokens[2] == ".." {
					node = node.GoUp()
				} else {
					node = node.Enter(tokens[2])
				}

			} else if tokens[1] == "ls" { //skip?
			}

		case "dir":
			node.AddDir(tokens[1])

		default:
			size, _ := strconv.Atoi(tokens[0])
			node.AddChild(tokens[1], size)

		}
	}

	usedSpace := tree["/"].GetSize()
	fmt.Printf("#1: %d\n", tree["/"].CountAllDUnder())
	fmt.Printf("#2: %d\n", tree["/"].FindClosestLarger(minimumForUpdate-(totalSpaceAvailable-usedSpace)))

	return nil
}
