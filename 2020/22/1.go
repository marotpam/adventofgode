package _2020

import (
	"strconv"
	"strings"
)

type deck []int

func (d deck) isEmpty() bool {
	return len(d) == 0
}

func (d deck) score() int {
	c := 0
	for i := 1; i <= len(d); i++ {
		c += i*d[len(d)-i]
	}
	return c
}

func (d deck) pop() (int, deck) {
	return d[0], d[1:]
}

func (d deck) add(cards ...int) deck {
	return append(d, cards...)
}

func Play(rawInput string) int {
	deck1, deck2 := parse(rawInput)
	for {
		if deck1.isEmpty() {
			return deck2.score()
		}

		if deck2.isEmpty() {
			return deck1.score()
		}

		var card1, card2 int
		card1, deck1 = deck1.pop()
		card2, deck2 = deck2.pop()

		if card1 > card2 {
			deck1 = deck1.add(card1, card2)
		} else {
			deck2 = deck2.add(card2, card1)
		}
	}
}

func parse(input string) (deck, deck) {
	rawDecks := strings.Split(input, "\n\n")
	return parseDeck(rawDecks[0]), parseDeck(rawDecks[1])
}

func parseDeck(s string) deck {
	lines := strings.Split(s, "\n")
	d := make(deck, 0, len(lines))
	for _, line := range lines[1:] {
		n, _ := strconv.Atoi(line)
		d = append(d, n)
	}
	return d
}
