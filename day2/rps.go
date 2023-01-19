package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
A | X -> Rock | score : 1
B| Y -> Paper | score: 2
C | Z -> Scissors | score: 3

Lose: 0
Draw: 3
Win: 6
*/

// First key is yours, second key is your opponent.
var res = map[string]map[string]int{
	"X": {
		"A": 4,
		"B": 1,
		"C": 7,
	},
	"Y": {
		"A": 8,
		"B": 5,
		"C": 2,
	},
	"Z": {
		"A": 3,
		"B": 9,
		"C": 6,
	},
}

// X -> Lose
// Y -> Draw
// Z -> Win

var det = map[string]map[string]string{
	"A": {
		"X": "Z",
		"Y": "X",
		"Z": "Y",
	},
	"B": {
		"X": "X",
		"Y": "Y",
		"Z": "Z",
	},
	"C": {
		"X": "Y",
		"Y": "Z",
		"Z": "X",
	},
}

func main() {
	f, err := os.Open("inp.txt")
	defer func() { _ = f.Close() }()

	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)
	score := 0
	for sc.Scan() {
		n := sc.Text()
		strat := strings.Split(n, " ")
		opp, act := strat[0], strat[1]
		score += res[det[opp][act]][opp]
	}
	fmt.Println(score)
}
