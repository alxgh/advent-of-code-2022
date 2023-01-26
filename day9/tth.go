package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	// {x, y}
	hpos, tpos := []int{0, 0}, []int{0, 0}
	v := make(map[[2]int]bool)
	for sc.Scan() {
		n := sc.Text()
		action := strings.Split(n, " ")
		dir := action[0]
		cnt, _ := strconv.Atoi(action[1])
		for i := 0; i < cnt; i++ {
			switch dir {
			case "R":
				hpos[1] += 1
			case "L":
				hpos[1] -= 1
			case "U":
				hpos[0] += 1
			case "D":
				hpos[0] -= 1
			}
			xdiff := hpos[0] - tpos[0]
			ydiff := hpos[1] - tpos[1]
			// check if tail is still is in touch with head
			if math.Abs(float64(xdiff)) > 1 {
				if xdiff > 0 {
					tpos[0] += 1
				} else {
					tpos[0] -= 1
				}
				if math.Abs(float64(ydiff)) > 0 {
					if ydiff > 0 {
						tpos[1] += 1
					} else {
						tpos[1] -= 1
					}
				}
			}
			if math.Abs(float64(ydiff)) > 1 {
				if ydiff > 0 {
					tpos[1] += 1
				} else {
					tpos[1] -= 1
				}
				if math.Abs(float64(xdiff)) > 0 {
					if xdiff > 0 {
						tpos[0] += 1
					} else {
						tpos[0] -= 1
					}
				}
			}
			v[[2]int{tpos[0], tpos[1]}] = true
		}
	}
	log.Println(len(v))
}
