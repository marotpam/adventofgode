package _2020

import "strings"

func CountAffirmativeAnswers1(rawInput string) int {
	groupAnswers := strings.Split(rawInput, "\n\n")
	c := 0

	for _, answers := range groupAnswers {
		individualAnswers := strings.Split(answers, "\n")
		positiveAnswersCounter := make(map[rune]int)
		for _, answer := range individualAnswers {
			for _, c := range answer {
				positiveAnswersCounter[c]++
			}
		}
		c += len(positiveAnswersCounter)
	}

	return c
}
