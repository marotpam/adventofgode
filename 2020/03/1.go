package main

import (
	"strings"
)

const tree = '#'

type slope struct {
	incRow, incCol int
}

func CountTreesInSlope(rawInput string, slope slope) int {
	row, col := 0, 0
	rows := strings.Split(rawInput, "\n")
	cols := len(rows[0])

	c := 0
	for row < len(rows) {
		if rune(rows[row][col%cols]) == tree {
			c++
		}

		row += slope.incRow
		col += slope.incCol
	}

	return c
}
