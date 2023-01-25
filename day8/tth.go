package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func mapToInt(s []string) []int {
	i := make([]int, len(s))
	for idx, c := range s {
		d, _ := strconv.Atoi(c)
		i[idx] = d
	}
	return i
}

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	forest := make([][]int, 0)
	for sc.Scan() {
		n := sc.Text()
		trees := strings.Split(n, "")
		forest = append(forest, mapToInt(trees))
	}

	rowsc := len(forest)
	colc := len(forest[0])

	m := make(map[int]bool, 0)
	for r, trees := range forest {
		lh := 0
		for c := 0; c < colc; c++ {
			idx := r*colc + c
			if r == 0 || r == rowsc-1 || c == 0 || c == colc-1 {
				m[idx] = true
			}
			h := trees[c]
			if h > lh {
				lh = h
				m[idx] = true
			}
		}
		lh = 0
		for c := colc - 1; c >= 0; c-- {
			idx := r*colc + c
			if r == 0 || r == rowsc-1 || c == 0 || c == colc-1 {
				m[idx] = true
			}
			h := trees[c]
			if h > lh {
				lh = h
				m[idx] = true
			}
		}
	}

	for c := 0; c < colc; c++ {
		lh := 0
		for r := 0; r < rowsc; r++ {
			idx := r*colc + c
			if r == 0 || r == rowsc-1 || c == 0 || c == colc-1 {
				m[idx] = true
			}
			h := forest[r][c]
			if h > lh {
				lh = h
				m[idx] = true
			}
		}
		lh = 0
		for r := rowsc - 1; r >= 0; r-- {
			idx := r*colc + c
			if r == 0 || r == rowsc-1 || c == 0 || c == colc-1 {
				m[idx] = true
			}
			h := forest[r][c]
			if h > lh {
				lh = h
				m[idx] = true
			}
		}
	}

	log.Println(len(m))

	highScore := 0
	for c := 0; c < colc; c++ {
		for r := 0; r < rowsc; r++ {
			score := 1
			t := forest[r][c]
			open := true
			// right
			for i := c + 1; i < colc; i++ {
				if forest[r][i] >= t {
					score *= i - c
					open = false
					break
				}
			}
			if open {
				score *= colc - 1 - c
			}
			open = true
			// left
			for i := c - 1; i >= 0; i-- {
				if forest[r][i] >= t {
					score *= c - i
					open = false
					break
				}
			}
			if open {
				score *= c
			}
			open = true
			// up
			for i := r - 1; i >= 0; i-- {
				if forest[i][c] >= t {
					score *= r - i
					open = false
					break
				}
			}
			if open {
				score *= r
			}
			open = true
			// down
			for i := r + 1; i < rowsc; i++ {
				if forest[i][c] >= t {
					score *= i - r
					open = false
					break
				}
			}
			if open {
				score *= rowsc - 1 - r
			}
			open = true
			if score > highScore {
				highScore = score
			}
		}
	}
	log.Println(highScore)
}
