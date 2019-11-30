package main

import "strings"

func CountFirstValidPassphrases(ps []string) int {
	result := 0

	for i := 0; i < len(ps); i++ {
		if isValidFirstPassphrase(ps[i]) {
			result++
		}
	}

	return result
}

func isValidFirstPassphrase(ps string) bool {
	words := strings.Split(ps, " ")
	wordMap := make(map[string]bool, len(words))
	for _, word := range words {
		if wordMap[word] {
			return false
		}
		wordMap[word] = true
	}
	return true
}