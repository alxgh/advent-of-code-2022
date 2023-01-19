package main

import (
	"bufio"
	"fmt"
	"math"
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
	sack := []int64{0, 0, 0}
	iter := 0
	for sc.Scan() {
		n := sc.Text()
		for _, i := range n {
			pos := calcPos(i)
			sack[iter] |= 1 << pos
		}
		iter += 1
		if iter == 3 {
			z := sack[0] & sack[1] & sack[2]
			if z > 0 {
				sum += int64(math.Log2(float64(z))) + 1
			}
			sack = []int64{0, 0, 0}
			iter = 0
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
