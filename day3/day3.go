package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	lowercasePriority = 1
	uppercasePriority = 27
)

/*
Goal: Find badge common in each group

Possible Solution:

  - Keep track of lines

  - Map each line

  - Check item with 3 count

  - A:65 -> Z:90

  - a:97 -> Z:122
*/

func searchBadge(matcher map[byte]int) byte {
	for k, v := range matcher {
		if v == 3 {
			return k
		}
	}

	return 0
}

func getPriority(b byte) int {
	if b <= 'Z' {
		return int(b - 'A' + uppercasePriority)
	}

	return int(b - 'a' + lowercasePriority)
}

func main() {
	f, err := os.Open("day3/rucksack.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var priorities int
	scanner := bufio.NewScanner(f)
	matcher := make(map[byte]int)
	seen := make(map[byte]bool)
	lineCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineCount += 1

		for i := 0; i < len(line); i++ {
			b := line[i]
			if v := seen[b]; !v {
				seen[b] = true
				matcher[b] += 1
			}
		}
		clear(seen)

		if lineCount%3 == 0 {
			badge := searchBadge(matcher)
			priorities += getPriority(badge)

			clear(matcher)
		}
	}

	fmt.Println(priorities)
}
