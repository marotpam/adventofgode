package _2020

func CountSeats2(rawInput string) int {
	seatPlan := parse(rawInput)

	for round := 1; seatPlan.Tick2() > 0; round++ {
	}

	return seatPlan.CountOccupiedSeats()
}

func (p *seatPlan) Tick2() int {
	c := 0
	rows := make([]row, 0, len(p.rows))

	for _, row := range p.rows {
		newSeats := make([]seat, 0, len(row))
		for _, seat := range row {
			newSeat := seat
			occupied := p.countOccupiedInVisible(seat)
			if seat.seatType == occupiedSeat && occupied >= 5 {
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

func (p *seatPlan) countOccupiedInVisible(s seat) int {
	return p.countOccupiedInDirection(s, -1, -1) +
		p.countOccupiedInDirection(s, -1, 0) +
		p.countOccupiedInDirection(s, -1, 1) +
		p.countOccupiedInDirection(s, 0, -1) +
		p.countOccupiedInDirection(s, 0, 1) +
		p.countOccupiedInDirection(s, 1, -1) +
		p.countOccupiedInDirection(s, 1, 0) +
		p.countOccupiedInDirection(s, 1, 1)
}

func (p *seatPlan) countOccupiedInDirection(s seat, incRow, incCol int) int {
	row, col := s.row+incRow, s.column+incCol
	if row < 0 || col < 0 {
		return 0
	}

	if row >= len(p.rows) || col >= p.columnCount {
		return 0
	}

	seat := p.rows[row][col]
	switch seat.seatType {
	case emptySeat:
		return 0
	case occupiedSeat:
		return 1
	default:
		return p.countOccupiedInDirection(seat, incRow, incCol)
	}
}