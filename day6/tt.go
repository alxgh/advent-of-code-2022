package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanRunes)
	marker := [4]uint8{0, 0, 0, 0}
	var x int64
	cnt := 0
	for sc.Scan() {
		cnt += 1
		n := sc.Bytes()
		c := uint8(n[0])
		marker[0] = c
		x = 0
		found := true
		for _, m := range marker {
			if m == 0 {
				found = false
				break
			}
			mm := int64(1 << (m - 64))
			if x&mm > 0 {
				found = false
				break
			}
			x |= mm
		}
		if found {
			fmt.Println(cnt)
			break
		}
		copy(marker[1:4], marker[0:3])
		marker[0] = 0
	}
}
