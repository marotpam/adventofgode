package _2020

import (
	"fmt"
	"strings"
)

const (
	activeCell   = '#'
	inactiveCell = '.'
)

type position struct {
	x, y, z int
}

func (p position) neighbours() []position {
	positions := make([]position, 0, 0)

	incs := []int{0, 1, -1}
	for incX := 0; incX < len(incs); incX++ {
		for incY := 0; incY < len(incs); incY++ {
			for incZ := 0; incZ < len(incs); incZ++ {
				if incX+incY+incZ > 0 {
					positions = append(positions, position{
						x: p.x + incs[incX],
						y: p.y + incs[incY],
						z: p.z + incs[incZ],
					})
				}
			}
		}
	}

	return positions
}

type world struct {
	cells            map[position]rune
	minX, minY, minZ int
	maxX, maxY, maxZ int
}

func (w *world) countActive() int {
	c := 0
	for _, cell := range w.cells {
		if cell == activeCell {
			c++
		}
	}
	return c
}

func (w *world) tick() {
	newCells := make(map[position]rune)
	for z := w.minZ; z <= w.maxZ; z++ {
		for y := w.minY; y <= w.maxY; y++ {
			for x := w.minX; x <= w.maxX; x++ {
				pos := position{
					x: x,
					y: y,
					z: z,
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

	w.cells = newCells
	w.minX--
	w.minY--
	w.minZ--
	w.maxX++
	w.maxY++
	w.maxZ++
}

func (w *world) countActiveNeighbours(p position) int {
	c := 0
	for _, neighbourPos := range p.neighbours() {
		if w.hasActiveCellIn(neighbourPos) {
			c++
		}
	}
	return c
}

func (w *world) hasActiveCellIn(pos position) bool {
	cell, ok := w.cells[pos]
	return ok && cell == activeCell
}

func (w *world) print() {
	for z := w.minZ; z <= w.maxZ; z++ {
		fmt.Printf("z=%d\n", z)
		for y := w.minY; y < w.maxY; y++ {
			for x := w.minX; x < w.maxX; x++ {
				cell, ok := w.cells[position{
					x: x,
					y: y,
					z: z,
				}]
				if ok {
					fmt.Print(string(cell))
				} else {
					fmt.Print(string(inactiveCell))
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func CountActiveCubes(rawInput string) int {
	world := parse(rawInput)
	for i := 0; i < 6; i++ {
		world.tick()
	}
	return world.countActive()
}

func parse(input string) world {
	lines := strings.Split(input, "\n")
	cells := make(map[position]rune)

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			cells[position{
				x: x,
				y: y,
				z: 0,
			}] = rune(lines[y][x])
		}
	}

	return world{
		cells: cells,
		minX:  -1, minY: -1, minZ: -1,
		maxX: len(lines[0]),
		maxY: len(lines),
		maxZ: 1,
	}
}
