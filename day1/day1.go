package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	max := 0
	maxs := make([]int, 3)
	current := 0
	for sc.Scan() {
		n := sc.Text()
		if len(n) == 0 {
			if current > max {
				max = current
			}
			sort.Ints(maxs)
			if current > maxs[0] {
				maxs[0] = current
			}
			current = 0
			continue
		}
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		current += i
	}
	if current > max {
		max = current
	}
	sort.Ints(maxs)
	if current > maxs[0] {
		maxs[0] = current
	}
	fmt.Println("MAX:", max)
	fmt.Println("Three max sum", maxs[0]+maxs[1]+maxs[2])
}
