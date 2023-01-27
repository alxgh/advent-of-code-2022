package main

import (
	"bufio"
	"log"
	"os"
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
	cycle := 0
	xreg := 1
	var fn func()
	istrCycles := 0
	sum := 0
	for sc.Scan() {
		n := sc.Text()
		// addx -> 2 cycles
		// noop -> 1 cycle

		if n == "noop" {
			istrCycles = 1
			fn = nil
		} else if n[0:4] == "addx" {
			num, _ := strconv.Atoi(n[5:])
			istrCycles = 2
			fn = func() {
				xreg += num
			}
		}
		for istrCycles > 0 {
			cycle += 1
			istrCycles -= 1

			if (cycle+20)%40 == 0 {
				log.Println(cycle, xreg)
				sum += cycle * xreg
			}
		}
		if fn != nil {
			fn()
		}
	}
	log.Println(sum)
}
