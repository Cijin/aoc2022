package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseCalories(lines []string) ([][]int, error) {
	var out [][]int
	var temp []int

	for i, line := range lines {
		if line == "" || i+1 == len(lines) {
			out = append(out, temp)
			temp = nil
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}

		temp = append(temp, num)
	}

	return out, nil
}

func getInput() ([][]int, error) {
	data, err := os.ReadFile("day1/calories.txt")
	if err != nil {
		return nil, err
	}

	// parse through data
	lines := strings.Split(string(data), "\n")

	return parseCalories(lines)
}

func sortedTotalCalories(pack [][]int) []int {
	var curr int
	var totals []int

	for _, calories := range pack {
		for _, c := range calories {
			curr += c
		}

		totals = append(totals, curr)
		curr = 0
	}

	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	return totals
}
