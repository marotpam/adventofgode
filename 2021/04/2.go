package _2021

func PlayBingoToLose(rawInput string) int {
	game := parse(rawInput)

	return game.playToLose()
}

func (g *game) playToLose() int {
	for _, n := range g.drawnNumbers {
		i := 0
		for ; i < len(g.boards); i++ {
			b := g.boards[i]
			won := b.mark(n)
			if !won {
				continue
			}

			if len(g.boards) == 1 {
				return b.score(n)
			}
			g.boards = append(g.boards[:i], g.boards[i+1:]...)
			i--
		}
	}
	return 0
}
