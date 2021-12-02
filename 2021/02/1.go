package _2021

import (
	"strconv"
	"strings"
)

type position struct {
	x, y int
}

func MultiplyPositionPart1(rawInput string) int {
	pos := position{}
	for _, line := range strings.Split(rawInput, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		increment, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			pos.x += increment
		case "up":
			pos.y -= increment
		case "down":
			pos.y += increment
		}
	}
	return pos.x * pos.y
}
