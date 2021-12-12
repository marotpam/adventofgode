package _2021

import "strings"

var closingChars = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var openingChars = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var syntaxErrorScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func CalculateSyntaxErrorScore(rawInput string) int {
	score := 0
	for _, line := range strings.Split(rawInput, "\n") {
		score += syntaxErrorScoreForLine(line)
	}
	return score
}

func syntaxErrorScoreForLine(line string) int {
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
			return syntaxErrorScores[c]
		}
		chars = chars[1:]
	}
	return 0
}
