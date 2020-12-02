package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func CountValidPasswords1(rawInput string) int {
	lines := parse(rawInput)

	c := 0
	for _, l := range lines {
		if l.password.isValid(l.policy) {
			c++
		}
	}

	return c
}

type policy struct {
	minOccurrences, maxOccurrences int
	letter                         rune
}

type password string

type line struct {
	policy   policy
	password password
}

func (p password) isValid(policy policy) bool {
	occurrences := map[rune]int{}

	for _, l := range p {
		occurrences[l]++
		if l == policy.letter && occurrences[l] > policy.maxOccurrences {
			return false
		}
	}

	return occurrences[policy.letter] >= policy.minOccurrences
}

func parse(input string) []line {
	passwords := make([]line, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		var min, max int
		var letter rune
		var pass password

		_, err := fmt.Sscanf(text, `%d-%d %c: %s`, &min, &max, &letter, &pass)
		if err != nil {
			log.Fatal("cannot read", err)
		}

		passwords = append(passwords, line{
			policy: policy{
				minOccurrences: min,
				maxOccurrences: max,
				letter:         letter,
			},
			password: pass,
		})
	}

	return passwords
}
