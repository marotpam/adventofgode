package _2021

import (
	"strconv"
	"strings"
)

func MultiplyPositionPart2(rawInput string) int {
	pos := position{}
	aim := 0
	for _, line := range strings.Split(rawInput, "\n") {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			continue
		}
		increment, _ := strconv.Atoi(parts[1])

		switch parts[0] {
		case "forward":
			pos.x += increment
			pos.y += increment * aim
		case "up":
			aim -= increment
		case "down":
			aim += increment
		}
	}
	return pos.x * pos.y
}
