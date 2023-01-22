package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack[T any] struct {
	s []T
}

func (s *stack[T]) push(r T) {
	s.s = append(s.s, r)
}

func (s *stack[T]) pop() *T {
	idx := len(s.s) - 1
	if idx < 0 {
		return nil
	}
	x := s.s[idx]
	s.s = s.s[0:idx]
	return &x
}

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	stacks := make([]stack[uint8], 0)
	for sc.Scan() {
		n := sc.Text()
		if len(n) == 0 {
			continue
		}
		if n[0:4] == "move" {
			sp := strings.Split(n[5:], " from ")
			cnt, _ := strconv.Atoi(sp[0])
			stackIdxs := strings.Split(sp[1], " to ")
			fromStackIdx, _ := strconv.Atoi(stackIdxs[0])
			toStackIdx, _ := strconv.Atoi(stackIdxs[1])
			temps := stack[uint8]{}
			for i := 0; i < cnt; i++ {
				x := stacks[fromStackIdx-1].pop()
				temps.push(*x)
			}
			for {
				x := temps.pop()
				if x == nil {
					break
				}
				stacks[toStackIdx-1].push(*x)
			}
		} else {
			for i := 0; i < len(n); i += 4 {
				stackIdx := i / 4
				if len(stacks) < stackIdx+1 {
					stacks = append(stacks, stack[uint8]{s: make([]uint8, 0)})
				}
				if n[i] == '[' {
					stacks[stackIdx].s = append([]uint8{n[i+1]}, stacks[stackIdx].s...)
				}
			}
		}
	}

	for _, s := range stacks {
		fmt.Printf("%c", *s.pop())
	}
}
