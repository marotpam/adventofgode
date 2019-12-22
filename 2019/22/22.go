package main

import (
	"strconv"
	"strings"
)

type deck struct {
	cards []int
}

func newDeck(cardsCount int) *deck {
	cards := make([]int, 0, cardsCount)

	for i := 0; i < cardsCount; i++ {
		cards = append(cards, i)
	}

	return &deck{cards: cards}
}

func (d *deck) dealIntoNewStack() {
	newCards := make([]int, len(d.cards), len(d.cards))

	for i := range newCards {
		newCards[i] = d.cards[len(d.cards)-1-i]
	}

	d.cards = newCards
}

func (d *deck) cut(n int) {
	cut := n
	if n < 0 {
		cut = len(d.cards) + n
	}

	d.cards = append(d.cards[cut:len(d.cards)], d.cards[:cut]...)
}

func (d *deck) dealWithIncrement(n int) {
	newCards := make([]int, len(d.cards), len(d.cards))
	pos := 0

	for i := 0; i < len(d.cards); i++ {
		newCards[pos] = d.cards[i]
		pos = (pos + n) % len(d.cards)
	}

	d.cards = newCards
}

func (d *deck) findPositionOfCard(c int) int {
	for i, card := range d.cards {
		if card == c {
			return i
		}
	}
	return -1
}

const (
	techniqueDealIntoNewStack        = "deal into new stack"
	techniqueDealWithIncrementPrefix = "deal with increment"
)

func FindCardPositionAfterShuffling(instructions []string, card int) int {
	d := newDeck(10007)

	for _, instruction := range instructions {
		if instruction == techniqueDealIntoNewStack {
			d.dealIntoNewStack()
			continue
		}
		parts := strings.Fields(instruction)
		n, _ := strconv.Atoi(parts[len(parts)-1])

		if strings.HasPrefix(instruction, techniqueDealWithIncrementPrefix) {
			d.dealWithIncrement(n)
			continue
		}

		d.cut(n)
	}

	return d.findPositionOfCard(card)
}
