package _2020

import "strings"

func CountAffirmativeAnswers2(rawInput string) int {
	groupAnswers := strings.Split(rawInput, "\n\n")
	c := 0

	for _, answers := range groupAnswers {
		individualAnswers := strings.Split(answers, "\n")
		positiveAnswers := make(map[rune]int)
		for _, answer := range individualAnswers {
			for _, c := range answer {
				positiveAnswers[c]++
			}
		}
		for k := range positiveAnswers {
			if positiveAnswers[k] == len(individualAnswers) {
				c++
			}
		}
	}

	return c
}
