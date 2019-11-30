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
	cells map[xy]int
}

func newGrid(n int) *grid {
	cells := map[xy]int{}

	for y := 1; y <= 300; y++ {
		for x := 1; x <= 300; x++ {
			xy := xy{x, y}
			cell := &cell{xy, n}
			cells[xy] = cell.powerLevel()
		}
	}

	return &grid{cells}
}

func (g grid) bestForSquareOfSize(s int) (xy, int) {
	bestPosition := xy{}
	powerInBestPosition := math.MinInt64

	for y := 1; y <= 300 - s; y++ {
		for x := 1; x <= 300 - s; x++ {
			powerAtPosition := 0
			for i := 0; i < s; i++ {
				for j := 0; j < s; j++ {
					powerAtPosition += g.cells[xy{x+j, y+i}]
				}
			}
			if powerAtPosition > powerInBestPosition {
				powerInBestPosition = powerAtPosition
				bestPosition = xy{x, y}
			}
		}
	}

	return bestPosition, powerInBestPosition
}

type result struct {
	squareSize, powerInBestSquare int
	bestPosition xy
}

func main() {
	g := newGrid(1308)

	ch := make(chan result, 300)

	for i := 1; i <= 300; i++ {
		go func(i int) {
			pos, power := g.bestForSquareOfSize(i)
			ch <- result{squareSize:i, bestPosition:pos, powerInBestSquare:power}
		}(i)
	}

	bestSquareSize := 0
	positionInBestSquare := xy{}
	powerInBestSquare := 0
	for i := 1; i <= 300; i++ {
		r := <- ch

		if r.powerInBestSquare > powerInBestSquare {
			bestSquareSize = r.squareSize
			positionInBestSquare = r.bestPosition
			powerInBestSquare = r.powerInBestSquare
		}
	}

	fmt.Println(positionInBestSquare, bestSquareSize) // {227,199} 19
}
