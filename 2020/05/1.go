package _2020

import (
	"strconv"
	"strings"
)

type seat struct {
	row, column int
}

func (s seat) getID() int {
	return s.row*8 + s.column
}

func getSeat(s string) seat {
	return seat{
		row:    decode(s[0:7], "F", "B"),
		column: decode(s[7:], "L", "R"),
	}
}

func decode(chars, encodedZero, encodedOne string) int {
	binaryNumber := strings.NewReplacer(encodedZero, "0", encodedOne, "1").Replace(chars)
	parseInt, _ := strconv.ParseInt(binaryNumber, 2, 64)
	return int(parseInt)
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
