package _2020

import "strings"

type seat struct {
	row, column int
}

func (s seat) getID() int {
	return s.row*8 + s.column
}

func getSeat(s string) seat {
	return seat{
		row:    findSpot(s[0:7], 'F', 'B', 127),
		column: findSpot(s[7:], 'L', 'R', 7),
	}
}

func findSpot(chars string, lowerHalf, upperHalf int32, upperBound int) int {
	lo, hi := 0, upperBound
	for _, c := range chars {
		switch c {
		case lowerHalf:
			hi = (hi + lo) / 2
		case upperHalf:
			lo = (lo + hi + 1) / 2
		}
	}
	return lo
}

func GetHighestSeatID(rawInput string) int {
	max := 0

	for _, line := range strings.Split(rawInput, "\n") {
		s := getSeat(line)
		if s.getID() > max {
			max = s.getID()
		}
	}

	return max
}
