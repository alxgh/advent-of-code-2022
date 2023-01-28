package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	op          func(old int) int
	test        func(w int) int
	inspections int
}

func (m *monkey) pop() (int, error) {
	if len(m.items) == 0 {
		return 0, errors.New("no item")
	}
	item := m.items[0]
	m.items = m.items[1:]
	return item, nil
}

func (m *monkey) push(item int) {
	m.items = append(m.items, item)
}

func lcm(n []int) int {
	c := make([]int, len(n))
	copy(c[:], n[:])
	for {
		prev := c[0]
		f := true
		min := math.MaxInt64
		minIdx := 0
		for idx, i := range c {
			if i != prev {
				f = false
			}
			if i < min {
				min = i
				minIdx = idx
			}
		}
		if f {
			return c[0]
		}
		c[minIdx] = c[minIdx] + n[minIdx]
	}
}

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	monkeys := make(map[int]*monkey)
	d := make([]int, 0)
	for {
		if !sc.Scan() {
			break
		}
		m := monkey{}
		n := sc.Text()
		number, _ := strconv.Atoi(n[7:8])
		sc.Scan()
		n = strings.TrimSpace(sc.Text())
		itemsStr := strings.Split(n[16:], ", ")
		m.items = make([]int, 0, len(itemsStr))
		for _, i := range itemsStr {
			itemInt, _ := strconv.Atoi(i)
			m.items = append(m.items, itemInt)
		}
		sc.Scan()
		n = sc.Text()
		opStr := strings.Split(strings.TrimPrefix(strings.TrimSpace(n), "Operation: new = "), " ")

		m.op = func(old int) int {
			a, b := old, old
			if opStr[0] != "old" {
				a, _ = strconv.Atoi(opStr[0])
			}
			if opStr[2] != "old" {
				b, _ = strconv.Atoi(opStr[2])
			}
			if opStr[1] == "+" {
				return a + b
			}
			return a * b
		}
		sc.Scan()
		n = strings.TrimSpace(sc.Text())
		testNum, _ := strconv.Atoi(strings.TrimSpace(n)[19:])
		sc.Scan()
		n = strings.TrimSpace(sc.Text())
		trueAction, _ := strconv.Atoi(n[25:])

		sc.Scan()
		n = strings.TrimSpace(sc.Text())
		falseAction, _ := strconv.Atoi(n[26:])
		m.test = func(w int) int {
			if w%testNum == 0 {
				return trueAction
			}
			return falseAction
		}
		monkeys[number] = &m
		d = append(d, testNum)
		sc.Scan()
	}

	l := len(monkeys)
	div := lcm(d)
	for i := 0; i < 10000; i++ {
		for z := 0; z < l; z++ {
			m := monkeys[z]
			for {
				item, err := m.pop()
				if err != nil {
					break
				}
				m.inspections += 1
				newWorry := m.op(item) / 1
				newMonkey := m.test(newWorry)
				monkeys[newMonkey].push(newWorry % div)
			}
		}
	}

	lvl := make([]int, 0, l)
	for _, m := range monkeys {
		lvl = append(lvl, m.inspections)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lvl)))
	fmt.Println(int(lvl[0]) * int(lvl[1]))
}
