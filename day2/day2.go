package day2

import (
	"log"
	"os"
	"strings"
)

/*
Score for single round:

	Rock 1, Paper 2, Scissors 3
	Plus
	Loose 0, Draw 3, Win 6

Symbols:

	A|X Rock, B|Y Paper, C|Z Scissors

Updated Strategy Symbol:

	A Rock, B Paper, C Scissors
	X Loose, Y Draw, Z Win
*/

var initialStrategy = map[string]int{
	"AX": 4,
	"AY": 8,
	"AZ": 3,

	"BX": 1,
	"BY": 5,
	"BZ": 9,

	"CX": 7,
	"CY": 2,
	"CZ": 6,
}

var updatedStrategy = map[string]int{
	"AX": 3,
	"AY": 4,
	"AZ": 8,

	"BX": 1,
	"BY": 5,
	"BZ": 9,

	"CX": 2,
	"CY": 6,
	"CZ": 7,
}

func readStrategyFile() []string {
	data, err := os.ReadFile("day2/strategy.txt")
	if err != nil {
		log.Fatal("Error reading strategy file:", err)
	}

	plays := strings.Split(string(data), "\n")
	return plays
}

func calculateScore(strategy map[string]int) int {
	var total int
	plays := readStrategyFile()

	for _, p := range plays {
		idx := strings.ReplaceAll(p, " ", "")
		total += strategy[idx]
	}

	return total
}
