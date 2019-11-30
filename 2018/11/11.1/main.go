package main

import (
	"fmt"
	"math"
)

type xy struct {
	x, y int
}

type cell struct {
	xy xy
	serialNumber int
}

func (c *cell) powerLevel() int {
	rackID := c.xy.x + 10
	powerLevel := rackID*c.xy.y
	powerLevel += c.serialNumber
	powerLevel *= rackID

	if powerLevel < 100 {
		powerLevel = 0
	} else {
		powerLevel = powerLevel/100%10
	}

	return powerLevel - 5
}

type grid struct {
	cells map[xy]*cell
}

func newGrid(n int) *grid {
	cells := map[xy]*cell{}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			xy := xy{x, y}
			cells[xy] = &cell{xy, n}
		}
	}

	return &grid{cells}
}

func (g grid) bestPosition() xy {
	bestPosition := xy{}
	powerInBestPosition := math.MinInt64

	for y := 1; y <= 300 - 3; y++ {
		for x := 1; x <= 300 - 3; x++ {
			powerAtPosition := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					powerAtPosition += g.cells[xy{x+j, y+i}].powerLevel()
				}
			}
			if powerAtPosition > powerInBestPosition {
				powerInBestPosition = powerAtPosition
				bestPosition = xy{x, y}
			}
		}
	}

	return bestPosition
}

func main() {
	g := newGrid(1308)

	fmt.Println(g.bestPosition()) // 21,41
}
