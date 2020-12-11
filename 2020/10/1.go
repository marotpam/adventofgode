package _2020

import (
	"sort"
	"strconv"
	"strings"
)

func GetResult(rawInput string) int {
	ints := parse(rawInput)
	minVoltageIncrements := make(map[int]int, 3)

	sort.Ints(ints)

	for i := 0; i < len(ints)-1; i++ {
		minVoltageIncrements[ints[i+1]-ints[i]]++
	}

	return (minVoltageIncrements[1] + 1) * (minVoltageIncrements[3] + 1)
}

func parse(input string) []int {
	lines := strings.Split(input, "\n")
	ints := make([]int, 0, len(lines))

	for _, l := range lines {
		n, _ := strconv.Atoi(l)
		ints = append(ints, n)
	}

	return ints
}
