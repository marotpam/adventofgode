package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func CountValidPasswords1(rawInput string) int {
	passwords := parse(rawInput)

	c := 0
	for _, p := range passwords {
		if p.isValid() {
			c++
		}
	}

	return c
}

type policy struct {
	minOccurrences, maxOccurrences int
	letter rune
}

type password struct {
	policy policy
	password string
}

func (p password) isValid() bool {
	occurrences := map[rune]int{}

	for _, l := range p.password {
		occurrences[l]++
		if l == p.policy.letter && occurrences[l] > p.policy.maxOccurrences{
			return false
		}
	}

	return occurrences[p.policy.letter] >= p.policy.minOccurrences
}

func parse(input string) []password {
	passwords := make([]password, 0)

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}

		var min, max int
		var letter rune
		var pass string

		_, err := fmt.Sscanf(text, `%d-%d %c: %s`, &min, &max, &letter, &pass)
		if err != nil {
			log.Fatal("cannot read", err)
		}

		passwords = append(passwords, password{
			policy:   policy{
				minOccurrences: min,
				maxOccurrences: max,
				letter:         letter,
			},
			password: pass,
		})
	}

	return passwords
}
