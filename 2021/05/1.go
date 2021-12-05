package _2021

import (
	"strconv"
	"strings"
)

func CountOverlappingHydrothermalVentLines1(rawInput string) int {
	points := map[coordinate]int{}
	overlappingCount := 0

	lines := parseLines(rawInput)
	for _, l := range lines {
		coords := l.start.getCoordinatesUntil(l.end)
		for _, c := range coords {
			if _, ok := points[c]; !ok {
				points[c] = 0
			}
			points[c]++

			if points[c] == 2 {
				overlappingCount++
			}
		}
	}

	return overlappingCount
}

type coordinate struct {
	x, y int
}

func (c coordinate) getCoordinatesUntil(other coordinate) []coordinate {
	var coords []coordinate

	incrX, incrY, canReach := c.getIncrements(other)
	if !canReach {
		return coords
	}

	pos := c
	for {
		coords = append(coords, pos)
		if pos == other {
			return coords
		}

		pos.x += incrX
		pos.y += incrY
	}
}

func (c coordinate) getIncrements(other coordinate) (int, int, bool) {
	incrX, incrY := other.x-c.x, other.y-c.y
	if incrX != 0 && incrY == 0 {
		return normalise(incrX), 0, true
	}
	if incrX == 0 && incrY != 0 {
		return 0, normalise(incrY), true
	}
	return 0, 0, c == other
}

func normalise(n int) int {
	return n / abs(n)
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

type line struct {
	start, end coordinate
}

func parseLines(input string) []line {
	rawLines := strings.Split(input, "\n")
	lines := make([]line, 0, len(rawLines))
	for _, l := range rawLines {
		parts := strings.Split(l, " -> ")
		lines = append(lines, line{
			start: parseCoordinate(parts[0]),
			end:   parseCoordinate(parts[1]),
		})
	}
	return lines
}

func parseCoordinate(s string) coordinate {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	return coordinate{
		x: x,
		y: y,
	}
}
