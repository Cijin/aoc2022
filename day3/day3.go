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
Goal: Find item types that appear in both compartments of rucksack

Possible Solution:
  - Go through each line, split items into 2 by the middle
  - Use a map to see if items in the first half appear in the second half
  - This way only need to traverse once per line
  - Might not need to convert to string
  - A:65 -> Z:90
  - a:97 -> Z:122
*/
func main() {
	var priorities int

	f, err := os.Open("day3/rucksack.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	matcher := make(map[byte]int)

	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < len(line)/2; i++ {
			r := line[i]
			matcher[r] = matcher[r] + 1
		}

		for j := len(line) / 2; j < len(line); j++ {
			r := line[j]
			if _, ok := matcher[r]; ok {
				if r <= 'Z' {
					priorities += int(r - 'A' + uppercasePriority)
					delete(matcher, r)
				}

				if r >= 'a' {
					priorities += int(r - 'a' + lowercasePriority)
					delete(matcher, r)
				}
			}
		}

		clear(matcher)
	}

	fmt.Println(priorities)
}
