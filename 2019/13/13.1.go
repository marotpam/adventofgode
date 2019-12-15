package main

import (
	"fmt"

	interpreter "github.com/marotpam/adventofgode/2019/intcode"
)

const (
	emptyTile  = 0
	wallTile   = 1
	blockTile  = 2
	paddleTile = 3
	ballTile   = 4

	outputActionX      = 0
	outputActionY      = 1
	outputActionTileID = 2
)

type position struct {
	x, y int
}

func (p *position) key() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

type arcade struct {
	interpreter  *interpreter.Interpreter
	pixels       map[string]int
	outputAction int
	position     position
}

func (a *arcade) Write(n int) {
	switch a.outputAction {
	case outputActionX:
		a.position.x = n
	case outputActionY:
		a.position.y = n
	case outputActionTileID:
		a.pixels[a.position.key()] = n
	}
	a.outputAction = (a.outputAction + 1) % 3
}

func (a *arcade) Read() int {
	panic("implement me")
}

func (a *arcade) render(instructions []int) {
	a.interpreter.Run(instructions, a, a)
}

func (a *arcade) countTilesOfType(tile int) int {
	c := 0
	for _, i := range a.pixels {
		if i == tile {
			c++
		}
	}

	return c
}

func newArcade() *arcade {
	return &arcade{
		interpreter:  interpreter.NewInterpreter(),
		pixels:       make(map[string]int, 0),
		outputAction: outputActionX,
	}
}

func CountBlockTiles(instructions []int) int {
	a := newArcade()
	a.render(instructions)

	return a.countTilesOfType(blockTile)
}
