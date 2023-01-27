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
	var knots [10][2]int
	v := make(map[[2]int]bool)
	for sc.Scan() {
		n := sc.Text()
		action := strings.Split(n, " ")
		dir := action[0]
		cnt, _ := strconv.Atoi(action[1])
		for i := 0; i < cnt; i++ {
			switch dir {
			case "R":
				knots[0][0] += 1
			case "L":
				knots[0][0] -= 1
			case "U":
				knots[0][1] += 1
			case "D":
				knots[0][1] -= 1
			}
			for x := 1; x < 10; x++ {
				xdiff := knots[x-1][0] - knots[x][0]
				ydiff := knots[x-1][1] - knots[x][1]
				// check if tail is still is in touch with head
				if math.Abs(float64(xdiff)) > 1 {
					if xdiff > 0 {
						knots[x][0] += 1
					} else {
						knots[x][0] -= 1
					}
					if math.Abs(float64(ydiff)) > 0 {
						if ydiff > 0 {
							knots[x][1] += 1
						} else {
							knots[x][1] -= 1
						}
					}
					continue
				}
				if math.Abs(float64(ydiff)) > 1 {
					if ydiff > 0 {
						knots[x][1] += 1
					} else {
						knots[x][1] -= 1
					}
					if math.Abs(float64(xdiff)) > 0 {
						if xdiff > 0 {
							knots[x][0] += 1
						} else {
							knots[x][0] -= 1
						}
					}
				}
			}
			//log.Println(knots)
			v[[2]int{knots[9][0], knots[9][1]}] = true
		}
	}
	log.Println(len(v))
}
