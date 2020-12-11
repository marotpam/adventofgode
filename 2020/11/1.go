package _2020

import (
	"strings"
)

const (
	aile         = '.'
	occupiedSeat = '#'
	emptySeat    = 'L'
)

func CountSeats1(rawInput string) int {
	seatPlan := parse(rawInput)

	for round := 1; seatPlan.Tick1() > 0; round++ {
	}

	return seatPlan.CountOccupiedSeats()
}

type seat struct {
	row, column int
	seatType    rune
}

type row []seat

type seatPlan struct {
	columnCount int
	rows        []row
}

func (p *seatPlan) Tick1() int {
	c := 0
	rows := make([]row, 0, len(p.rows))

	for _, row := range p.rows {
		newSeats := make([]seat, 0, len(row))
		for _, seat := range row {
			newSeat := seat
			occupied := p.countOccupiedInAdjacent(seat)
			if seat.seatType == occupiedSeat && occupied >= 4 {
				newSeat.seatType = emptySeat
				c++
			} else if seat.seatType == emptySeat && occupied == 0 {
				newSeat.seatType = occupiedSeat
				c++
			}

			newSeats = append(newSeats, newSeat)
		}
		rows = append(rows, newSeats)
	}

	*p = seatPlan{
		rows:        rows,
		columnCount: p.columnCount,
	}

	return c
}

func (p *seatPlan) CountOccupiedSeats() int {
	c := 0
	for _, row := range p.rows {
		for _, seat := range row {
			if seat.seatType == occupiedSeat {
				c++
			}
		}
	}
	return c
}

func (p *seatPlan) countOccupiedInAdjacent(s seat) int {
	c := 0
	increments := []int{-1, 0, 1}
	for incRow := 0; incRow < len(increments); incRow++ {
		for incCol := 0; incCol < len(increments); incCol++ {
			row, col := s.row+increments[incRow], s.column+increments[incCol]
			if row >= 0 && col >= 0 &&
				row < len(p.rows) && col < p.columnCount &&
				(row != s.row || col != s.column) &&
				p.rows[row][col].seatType == occupiedSeat {
				c++
			}
		}
	}
	return c
}

func parse(input string) seatPlan {
	lines := strings.Split(input, "\n")
	columnCount := 0
	rows := make([]row, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		row := row{}
		for j := 0; j < len(lines[i]); j++ {
			row = append(row, seat{
				row:      i,
				column:   j,
				seatType: rune(lines[i][j]),
			})
		}
		columnCount = len(row)
		rows = append(rows, row)
	}
	return seatPlan{
		rows:        rows,
		columnCount: columnCount,
	}
}
