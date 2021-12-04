package _2021

import (
	"strconv"
	"strings"
)

func PlayBingoToWin(rawInput string) int {
	game := parse(rawInput)

	return game.playToWin()
}

type number struct {
	value int
	seen  bool
}

type board struct {
	numbers          [][]*number
	missingInRows    []int
	missingInColumns []int
}

func newBoard(numbers [][]*number) *board {
	var missingInColumns []int
	var missingInRows []int
	for i := 0; i < len(numbers); i++ {
		missingInRows = append(missingInRows, len(numbers))
		missingInColumns = append(missingInColumns, len(numbers))
	}
	return &board{
		numbers:          numbers,
		missingInRows:    missingInRows,
		missingInColumns: missingInColumns,
	}
}

func (b *board) mark(n int) bool {
	finished := false
	for i := 0; i < len(b.numbers); i++ {
		for j := 0; j < len(b.numbers[i]); j++ {
			number := b.numbers[i][j]
			if number.value != n || number.seen {
				continue
			}
			number.seen = true

			b.missingInRows[i]--
			b.missingInColumns[j]--
			if b.missingInRows[i] == 0 || b.missingInColumns[j] == 0 {
				finished = true
			}
		}
	}
	return finished
}

func (b *board) score(lastDrawn int) int {
	score := 0
	for i := 0; i < len(b.numbers); i++ {
		for j := 0; j < len(b.numbers[i]); j++ {
			number := b.numbers[i][j]
			if !number.seen {
				score += number.value
			}
		}
	}
	return score * lastDrawn
}

type game struct {
	drawnNumbers []int
	boards       []*board
}

func (g *game) playToWin() int {
	for _, n := range g.drawnNumbers {
		for _, b := range g.boards {
			won := b.mark(n)
			if won {
				return b.score(n)
			}
		}
	}
	return 0
}

func parse(rawInput string) *game {
	lines := strings.Split(rawInput, "\n")
	drawnNumbers := getInts(lines[0])

	var boards []*board
	var numbers [][]*number
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			if len(numbers) != 0 {
				boards = append(boards, newBoard(numbers))
				numbers = [][]*number{}
			}
			continue
		}
		numbers = append(numbers, getNumbers(lines[i]))
	}
	boards = append(boards, newBoard(numbers))

	return &game{
		drawnNumbers: drawnNumbers,
		boards:       boards,
	}
}

func getInts(line string) []int {
	parts := strings.Split(line, ",")
	ints := make([]int, 0, len(parts))
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		ints = append(ints, n)
	}
	return ints
}

func getNumbers(line string) []*number {
	parts := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' })
	numbers := make([]*number, 0, len(parts))
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		numbers = append(numbers, &number{value: n})
	}
	return numbers
}
