package _2023

func (c card) countWinningNumbers() int {
	points := 0

	m := make(map[int]struct{}, len(c.winningNumbers))
	for _, n := range c.winningNumbers {
		m[n] = struct{}{}
	}

	for _, n := range c.numbers {
		if _, ok := m[n]; ok {
			points++
		}
	}
	return points
}

func CountTotalScratchCards(rawInput string) int {
	cards := parseCards(rawInput)

	cardCounts := make(map[int]int, len(cards))
	cardCopies := make(map[int]int, len(cards))
	for i, c := range cards {
		cardCounts[i] = 1
		cardCopies[i] = c.countWinningNumbers()
	}

	total := 0
	for i := range cards {
		total += cardCounts[i]
		for j := 0; j < cardCopies[i]; j++ {
			cardCounts[i+j+1] += cardCounts[i]
		}
	}

	return total
}
