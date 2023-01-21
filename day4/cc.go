package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ati(s string) int {
	i, _ := strconv.Atoi(s)
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
	sum := 0
	for sc.Scan() {
		n := sc.Text()
		p := strings.Split(n, ",")
		p1 := strings.Split(p[0], "-")
		p2 := strings.Split(p[1], "-")
		if ati(p1[0]) <= ati(p2[1]) && ati(p1[1]) >= ati(p2[0]) {
			sum += 1
		}
		//if (ati(p1[0]) >= ati(p2[0]) && ati(p1[1]) <= ati(p2[1])) || (ati(p2[0]) >= ati(p1[0]) && ati(p2[1]) <= ati(p1[1])) {
		//	sum += 1
		//}
	}
	fmt.Println(sum)
}
