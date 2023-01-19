package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	sum := int64(0)
	for sc.Scan() {
		n := sc.Text()
		h := len(n)
		c1, c2 := n[0:(h/2)], n[h/2:h]
		log.Println(c1, c2)
		var x int64
		var y int64
		for _, i := range c1 {
			pos := calcPos(i)
			x |= 1 << pos
		}
		for _, i := range c2 {
			pos := calcPos(i)
			y |= 1 << pos
		}
		z := y & x
		for i := int64(0); i < 52; i++ {
			if p := z & (1 << i); p != 0 {
				sum += 1 + i
			}
		}
	}
	fmt.Println(sum)
}

func calcPos(i int32) int {
	if i >= 65 && i <= 90 {
		return int(i - 65 + 26)
	}
	return int(i - 97)
}
