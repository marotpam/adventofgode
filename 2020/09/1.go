package _2020

import (
	"strconv"
	"strings"
)

func FindNumber(rawInput string, preambleLength int) int {
	ints := parse(rawInput)

	for i, n := range ints[preambleLength:] {
		if isMissingInSums(n, ints[i:i+preambleLength]) {
			return n
		}
	}

	return -1
}

func isMissingInSums(n int, ints []int) bool {
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i]+ints[j] == n {
				return false
			}
		}
	}
	return true
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
