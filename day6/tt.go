package main

import (
	"bufio"
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
	sc.Split(bufio.ScanRunes)
	pmarker := [4]uint8{0, 0, 0, 0}
	mmarker := [14]uint8{}
	mc := 4
	pmf := false
	var x int64
	cnt := 0

	for sc.Scan() {
		cnt += 1
		n := sc.Bytes()
		c := uint8(n[0])
		pmarker[0] = c
		mmarker[0] = c
		x = 0
		found := true
		for i := 0; i < mc; i++ {
			var m uint8
			if pmf {
				m = mmarker[i]
			} else {
				m = pmarker[i]
			}
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
			if pmf {
				log.Println(cnt)
				break
			} else {
				pmf = true
				mc = 14
			}
		}
		copy(pmarker[1:4], pmarker[0:3])
		pmarker[0] = 0
		copy(mmarker[1:14], mmarker[0:13])
		mmarker[0] = 0
	}
}
