package _2020

import (
	"strconv"
	"strings"
)

func GetNth(rawInput string, wanted int) int {
	ints := parse(rawInput)
	occurrences := map[int][]int{}
	for i, n := range ints {
		occurrences[n] = []int{i + 1}
	}
	ints = append(ints, 0)
	_, seen0 := occurrences[0]
	if !seen0 {
		occurrences[0] = []int{len(ints)}
	} else {
		occurrences[0] = []int{len(ints), occurrences[0][0]}
	}

	for i := len(ints); i < wanted; i++ {
		timesSeen, seen := occurrences[ints[i-1]]
		if !seen {
			timesSeen = []int{i + 1}
		}

		n := i - timesSeen[0]
		if len(timesSeen) == 2 {
			n = timesSeen[0] - timesSeen[1]
		}

		ints = append(ints, n)
		_, seenN := occurrences[n]
		if !seenN {
			occurrences[n] = []int{i + 1}
		} else {
			occurrences[n] = []int{i + 1, occurrences[n][0]}
		}
	}

	return ints[len(ints)-1]
}

func parse(input string) []int {
	numbers := strings.Split(input, ",")
	ints := make([]int, 0, len(numbers))

	for _, rawNumber := range numbers {
		n, _ := strconv.Atoi(rawNumber)
		ints = append(ints, n)
	}

	return ints
}
