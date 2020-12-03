package _2019

import (
	"math"

	"github.com/marotpam/adventofgode/2019/intcode"
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

	joystickNeutral   = 0
	joystickTiltRight = 1
	joystickTiltLeft  = -1
)

type position struct {
	x, y int
}

func (p *position) isToTheRightOf(other position) bool {
	return p.x > other.x
}

func (p *position) isToTheLeftOf(other position) bool {
	return p.x < other.x
}

type arcade struct {
	interpreter      *intcode.Interpreter
	pixels           map[position]int
	outputAction     int
	position         position
	score            int
	ballTilePosition position
	paddlePosition   position
}

func (a *arcade) Write(n int) {
	switch a.outputAction {
	case outputActionX:
		a.position.x = n
	case outputActionY:
		a.position.y = n
	case outputActionTileID:
		if n == ballTile {
			a.ballTilePosition = a.position
		} else if n == paddleTile {
			a.paddlePosition = a.position
		}

		scorePosition := position{x: -1, y: 0}
		if a.position == scorePosition {
			a.score = n
		} else {
			a.pixels[a.position] = n
		}
	}

	a.outputAction = (a.outputAction + 1) % 3
}

func (a *arcade) Read() int {
	if a.paddlePosition.isToTheRightOf(a.ballTilePosition) {
		return joystickTiltLeft
	}

	if a.paddlePosition.isToTheLeftOf(a.ballTilePosition) {
		return joystickTiltRight
	}

	return joystickNeutral
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
		interpreter:  intcode.NewInterpreter(),
		pixels:       make(map[position]int, 0),
		outputAction: outputActionX,
		score:        0,
	}
}

func CountBlockTiles(instructions []int) int {
	a := newArcade()
	a.render(instructions)

	return a.countTilesOfType(blockTile)
}

func BeatGame(instructions []int) int {
	a := newArcade()
	instructions[0] = 2

	for blockCount := math.MaxInt64; blockCount > 0; blockCount = a.countTilesOfType(blockTile) {
		a.render(instructions)
	}

	return a.score
}
