package _2017

import (
	"sort"
	"strings"
)

func CountSecondValidPassphrases(ps []string) int {
	result := 0

	for i := 0; i < len(ps); i++ {
		if isValidSecondPassphrase(ps[i]) {
			result++
		}
	}

	return result
}

func isValidSecondPassphrase(ps string) bool {
	words := strings.Split(ps, " ")
	wordMap := make(map[string]bool, len(words))
	for _, word := range words {
		sorted := sortCharsInWord(word)
		if wordMap[sorted] {
			return false
		}
		wordMap[sorted] = true
	}
	return true
}

func sortCharsInWord(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}