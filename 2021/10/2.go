package _2021

import (
	"sort"
	"strings"
)

var autoCompleteScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func CalculateMiddleAutocompleteScore(rawInput string) int {
	lines := strings.Split(rawInput, "\n")
	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		score := autoCompleteScoreForLine(line)
		if score != 0 {
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func autoCompleteScoreForLine(line string) int {
	chars := make([]rune, 0, len(line))
	for _, c := range line {
		_, isOpeningChar := openingChars[c]
		if isOpeningChar {
			chars = append([]rune{c}, chars...)
			continue
		}

		_, isClosingChar := closingChars[c]
		if !isClosingChar {
			continue
		}

		if len(chars) == 0 || closingChars[c] != chars[0] {
			return 0
		}
		chars = chars[1:]
	}
	return autocompleteScoreFor(chars)
}

func autocompleteScoreFor(chars []rune) int {
	score := 0
	for _, c := range chars {
		score = score*5 + autoCompleteScores[openingChars[c]]
	}
	return score
}
