package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseCalories(lines []string) ([]int, error) {
	var out []int
	var temp int

	for i, line := range lines {
		if line == "" || i+1 == len(lines) {
			out = append(out, temp)
			temp = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		temp += num
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] > out[j]
	})

	return out, nil
}

func getSortedTotalCalories() ([]int, error) {
	data, err := os.ReadFile("day1/calories.txt")
	if err != nil {
		return nil, err
	}

	// parse through data
	lines := strings.Split(string(data), "\n")

	return parseCalories(lines)
}
