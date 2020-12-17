package _2020

import (
	"strings"
)

type position4d struct {
	x, y, z, w int
}

func (p position4d) neighbours() []position4d {
	positions := make([]position4d, 0, 0)

	incs := []int{0, 1, -1}
	for incX := 0; incX < len(incs); incX++ {
		for incY := 0; incY < len(incs); incY++ {
			for incZ := 0; incZ < len(incs); incZ++ {
				for incW := 0; incW < len(incs); incW++ {
					if incX+incY+incZ+incW > 0 {
						positions = append(positions, position4d{
							x: p.x + incs[incX],
							y: p.y + incs[incY],
							z: p.z + incs[incZ],
							w: p.w + incs[incW],
						})
					}
				}
			}
		}
	}

	return positions
}

type world4d struct {
	cells                  map[position4d]rune
	minX, minY, minZ, minW int
	maxX, maxY, maxZ, maxW int
}

func (w *world4d) countActive() int {
	c := 0
	for _, cell := range w.cells {
		if cell == activeCell {
			c++
		}
	}
	return c
}

func (w *world4d) tick4d() {
	w.expandDimensions()

	newCells := make(map[position4d]rune)
	for ww := w.minW; ww < w.maxW; ww++ {
		for z := w.minZ; z < w.maxZ; z++ {
			for y := w.minY; y < w.maxY; y++ {
				for x := w.minX; x < w.maxX; x++ {
					pos := position4d{
						x: x,
						y: y,
						z: z,
						w: ww,
					}
					activeNeighbours := w.countActiveNeighbours(pos)
					cell := inactiveCell
					isActiveCell := w.hasActiveCellIn(pos)
					if (isActiveCell && (activeNeighbours == 2 || activeNeighbours == 3)) ||
						(!isActiveCell && activeNeighbours == 3) {
						cell = activeCell
					}
					newCells[pos] = cell
				}
			}
		}
	}

	w.cells = newCells
}

func (w *world4d) expandDimensions() {
	w.minX--
	w.minY--
	w.minZ--
	w.minW--

	w.maxX++
	w.maxY++
	w.maxZ++
	w.maxW++
}

func (w *world4d) countActiveNeighbours(p position4d) int {
	c := 0
	for _, neighbourPos := range p.neighbours() {
		if w.hasActiveCellIn(neighbourPos) {
			c++
		}
	}
	return c
}

func (w *world4d) hasActiveCellIn(pos position4d) bool {
	cell, ok := w.cells[pos]
	return ok && cell == activeCell
}

func CountActiveCubes4d(rawInput string) int {
	world := parse4d(rawInput)
	for i := 0; i < 6; i++ {
		world.tick4d()
	}
	return world.countActive()
}

func parse4d(input string) world4d {
	lines := strings.Split(input, "\n")
	cells := make(map[position4d]rune)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			cells[position4d{
				x: x,
				y: y,
				z: 0,
				w: 0,
			}] = rune(lines[y][x])
		}
	}

	return world4d{
		cells: cells,
		minX:  0, minY: 0, minZ: 0, minW: 0,
		maxX: len(lines[0]),
		maxY: len(lines),
		maxZ: 1,
		maxW: 1,
	}
}
