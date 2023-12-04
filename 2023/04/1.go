package _2023

import (
	"strconv"
	"strings"
)

type card struct {
	winningNumbers []int
	numbers        []int
}

func (c card) countPoints() int {
	points := 0

	m := make(map[int]struct{}, len(c.winningNumbers))
	for _, n := range c.winningNumbers {
		m[n] = struct{}{}
	}

	for _, n := range c.numbers {
		if _, ok := m[n]; ok {
			if points == 0 {
				points = 1
			} else {
				points = points * 2
			}
		}
	}
	return points
}

func GetPoints(rawInput string) int {
	sum := 0
	for _, card := range parseCards(rawInput) {
		sum += card.countPoints()
	}
	return sum
}

func parseCards(rawInput string) []card {
	lines := strings.Split(rawInput, "\n")
	cards := make([]card, 0, len(lines))

	for _, l := range lines {
		lineParts := strings.Split(l, ": ")
		if len(lineParts) != 2 {
			continue
		}
		numbersParts := strings.Split(lineParts[1], " | ")
		if len(numbersParts) != 2 {
			continue
		}

		cards = append(cards, card{
			winningNumbers: parseNumbers(numbersParts[0]),
			numbers:        parseNumbers(numbersParts[1]),
		})
	}

	return cards
}

func parseNumbers(rawNumbers string) []int {
	numbers := make([]int, 0)
	for _, p := range strings.Split(rawNumbers, " ") {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			continue
		}
		numbers = append(numbers, n)
	}
	return numbers
}
