package main

import (
	"strings"
)

const (
	layoutBug        = '#'
	layoutEmptySpace = '.'
	dimensions       = 5
)

type position struct {
	x, y int
}

type world struct {
	layout map[position]bool
}

func newWorld(rawLayout string) *world {
	rows := strings.Split(strings.Trim(rawLayout, "\n"), "\n")

	l := make(map[position]bool, dimensions*dimensions)
	for y := 0; y < len(rows); y++ {
		for x := 0; x < len(rows[y]); x++ {
			l[position{x: x, y: y}] = rows[y][x] == layoutBug
		}
	}

	return &world{
		layout: l,
	}
}

func (w *world) countAdjacentBugs(p position) int {
	result := 0

	adjacentPositions := []position{
		{x: p.x + 1, y: p.y},
		{x: p.x - 1, y: p.y},
		{x: p.x, y: p.y + 1},
		{x: p.x, y: p.y - 1},
	}

	for _, adjacentPosition := range adjacentPositions {
		if w.hasBugIn(adjacentPosition) {
			result++
		}
	}

	return result
}

func (w *world) hasBugIn(p position) bool {
	return w.layout[p]
}

func (w *world) calculateBiodiversityRating() int {
	result := 0
	cellRating := 1
	for y := 0; y < dimensions; y++ {
		for x := 0; x < dimensions; x++ {
			if w.layout[position{x: x, y: y}] {
				result += cellRating
			}
			cellRating = cellRating << 1
		}
	}
	return result
}

func (w *world) print() string {
	rawLayout := ""
	for y := 0; y < dimensions; y++ {
		for x := 0; x < dimensions; x++ {
			if w.layout[position{x: x, y: y}] {
				rawLayout += string(layoutBug)
			} else {
				rawLayout += string(layoutEmptySpace)
			}
		}
		rawLayout += "\n"
	}
	return rawLayout
}

func (w *world) tick() {
	newLayout := make(map[position]bool, dimensions*dimensions)
	for position := range w.layout {
		adjacentBugs := w.countAdjacentBugs(position)
		cell := w.layout[position]
		if w.hasBugIn(position) {
			if adjacentBugs != 1 {
				cell = false
			}
		} else if adjacentBugs == 1 || adjacentBugs == 2 {
			cell = true
		}
		newLayout[position] = cell
	}
	w.layout = newLayout
}

func GetFirstRepeatedRating(rawLayout string) int {
	w := newWorld(rawLayout)
	ratings := map[int]bool{
		w.calculateBiodiversityRating(): true,
	}

	for {
		w.tick()

		r := w.calculateBiodiversityRating()
		if _, ok := ratings[r]; ok {
			return r
		}

		ratings[r] = true
	}
}
