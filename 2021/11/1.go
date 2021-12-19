package _2021

import (
	"strconv"
	"strings"
)

func CountFlashes(rawInput string, steps int) int {
	cave := parse(rawInput)
	count := 0
	for i := 0; i < steps; i++ {
		count += cave.countFlashes()
	}
	return count
}

type position struct {
	x, y int
}

func (p position) neighbours() []position {
	var positions []position
	increments := []int{-1, 0, 1}
	for y := 0; y < len(increments); y++ {
		for x := 0; x < len(increments); x++ {
			incx, incy := increments[x], increments[y]
			if incx == 0 && incy == 0 {
				continue
			}
			positions = append(positions, position{
				x: p.x + incx,
				y: p.y + incy,
			})
		}
	}
	return positions
}

type cave struct {
	octopuses map[position]int
}

func (c *cave) countFlashes() int {
	var flashingPositions []position
	for pos := range c.octopuses {
		c.octopuses[pos]++
		if c.octopuses[pos] == 10 {
			flashingPositions = append(flashingPositions, pos)
		}
	}
	return c.countFlashesRec(flashingPositions, len(flashingPositions))
}

func (c *cave) countFlashesRec(flashingPositions []position, total int) int {
	if len(flashingPositions) == 0 {
		c.clearFlashed()
		return total
	}

	var newFlashingPositions []position
	for _, pos := range flashingPositions {
		for _, n := range pos.neighbours() {
			_, ok := c.octopuses[n]
			if !ok {
				continue
			}

			c.octopuses[n]++
			if c.octopuses[n] == 10 {
				newFlashingPositions = append(newFlashingPositions, n)
			}
		}
	}

	return c.countFlashesRec(newFlashingPositions, total+len(newFlashingPositions))
}

func (c *cave) clearFlashed() {
	for pos := range c.octopuses {
		if c.octopuses[pos] > 9 {
			c.octopuses[pos] = 0
		}
	}
}

func parse(rawInput string) *cave {
	lines := strings.Split(rawInput, "\n")
	octopuses := make(map[position]int, len(lines))
	for y, line := range lines {
		for x, c := range line {
			n, _ := strconv.Atoi(string(c))
			octopuses[position{
				x: x,
				y: y,
			}] = n
		}
	}

	return &cave{
		octopuses: octopuses,
	}
}
