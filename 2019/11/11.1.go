package _2019

import (
	"github.com/marotpam/adventofgode/2019/intcode"
)

const (
	directionUp    = 0
	directionRight = 1
	directionDown  = 2
	directionLeft  = 3

	directionsCount = 4

	colourBlack = 0
	colourWhite = 1

	outputActionPaint = 0
	outputActionTurn  = 1

	actionTurnLeft  = 0
	actionTurnRight = 1
)

type direction int

type position struct {
	x, y int
}

type robot struct {
	interpreter   *intcode.Interpreter
	direction     direction
	position      position
	paintedPanels map[position]int
	outputAction  int
}

func (r *robot) Write(n int) {
	switch r.outputAction {
	case outputActionPaint:
		r.paintCurrentPanel(n)
		r.outputAction = outputActionTurn
	case outputActionTurn:
		if n == actionTurnLeft {
			r.turnLeft()
		} else {
			r.turnRight()
		}
		r.moveForward()
		r.outputAction = outputActionPaint
	}

}

func (r *robot) Read() int {
	c, ok := r.paintedPanels[r.position]
	if !ok {
		return colourBlack
	}

	return c
}

func newRobot() *robot {
	return &robot{
		direction: directionUp,
		position: position{
			x: 0,
			y: 0,
		},
		interpreter:   intcode.NewInterpreter(),
		outputAction:  outputActionPaint,
		paintedPanels: make(map[position]int, 0),
	}
}

func (r *robot) moveForward() {
	incX, incY := 0, 0
	switch r.direction {
	case directionUp:
		incY = 1
	case directionRight:
		incX = 1
	case directionDown:
		incY = -1
	case directionLeft:
		incX = -1
	}
	r.position.x += incX
	r.position.y += incY
}

func (r *robot) turnLeft() {
	r.direction = (r.direction - 1 + directionsCount) % directionsCount
}

func (r *robot) turnRight() {
	r.direction = (r.direction + 1) % directionsCount
}

func (r *robot) countPaintedPanels() int {
	return len(r.paintedPanels)
}

func (r *robot) paint(instructions []int) {
	r.interpreter.Run(instructions, r, r)
}

func (r *robot) paintCurrentPanel(colour int) {
	r.paintedPanels[r.position] = colour
}

func CountPanelsPaintedAtLeastOnce(instructions []int) int {
	r := newRobot()
	r.paint(instructions)

	return r.countPaintedPanels()
}
