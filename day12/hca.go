package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type path struct {
	c [][2]int
}

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	dest := [2]int{}
	m := make([][]uint8, 0)
	r := 0

	paths := make([]*path, 0)
	v := map[[2]int]bool{}
	for sc.Scan() {
		n := sc.Text()
		m = append(m, make([]uint8, 0, len(n)))
		for c, b := range strings.Split(n, "") {
			if b[0] == 'S' || b[0] == 'a' {
				m[r] = append(m[r], 'a')
				paths = append(paths, &path{c: [][2]int{{r, c}}})
				v[[2]int{r, c}] = true
			} else if b[0] == 'E' {
				dest[0] = r
				dest[1] = c
				m[r] = append(m[r], 'z')
			} else {
				m[r] = append(m[r], b[0])
			}
		}

		r += 1
	}
	cellsCount := len(m[0])
	rowsCount := len(m)

	var dpath path
	for {
		//sort.Slice(paths, func(i, j int) bool {
		//	return len(paths[i].c) < len(paths[j].c)
		//})
		shortestPath := paths[0]
		lastStep := shortestPath.c[len(shortestPath.c)-1]
		paths = paths[1:]
		//4 directions
		cells := [4][2]int{
			{lastStep[0], lastStep[1] + 1},
			{lastStep[0], lastStep[1] - 1},
			{lastStep[0] + 1, lastStep[1]},
			{lastStep[0] - 1, lastStep[1]},
		}

		for _, cell := range cells {
			if cell[0] < 0 || cell[0] >= rowsCount || cell[1] < 0 || cell[1] >= cellsCount {
				continue
			}
			if _, ok := v[cell]; ok {
				continue
			}

			if m[cell[0]][cell[1]]-1 > m[lastStep[0]][lastStep[1]] {
				continue
			}
			if cell == dest {
				shortestPath.c = append(shortestPath.c, cell)
				dpath = *shortestPath
				goto end
			}
			newc := make([][2]int, len(shortestPath.c))
			copy(newc, shortestPath.c)
			newc = append(newc, cell)
			v[cell] = true
			paths = append(paths, &path{
				c: newc,
			})
		}
		if len(paths) == 0 {
			break
		}
	}
end:
	log.Println(dpath.c)
	log.Println(len(dpath.c) - 1)
}
