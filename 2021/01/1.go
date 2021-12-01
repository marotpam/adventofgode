package _2021

import (
	"strconv"
	"strings"
)

func CountDepthIncreases(rawInput string) int {
	depths := parseDepths(rawInput)
	if len(depths) <= 1 {
		return 0
	}

	increases := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			increases++
		}
	}
	return increases
}

func parseDepths(rawInput string) []int {
	var depths []int
	for _, line := range strings.Split(rawInput, "\n") {
		d, _ := strconv.Atoi(line)
		depths = append(depths, d)
	}
	return depths
}
