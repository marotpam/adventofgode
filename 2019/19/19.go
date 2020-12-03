package _2019

import (
	"github.com/marotpam/adventofgode/2019/intcode"
)

const (
	outputPulled = 1
)

type position struct {
	x, y int
}

type drone struct {
	interpreter     *intcode.Interpreter
	maxPosition     position
	currentPosition position
	timesPulled     int
	readsCounter    int
}

func (d *drone) Read() int {
	var position int
	if d.readsCounter%2 == 0 {
		position = d.currentPosition.x
	} else {
		position = d.currentPosition.y
	}
	d.readsCounter++

	return position
}

func (d *drone) Write(n int) {
	if n == outputPulled {
		d.timesPulled++
	}
}

func newDrone(maxPosition position) *drone {
	return &drone{
		interpreter: intcode.NewInterpreter(),
		maxPosition: maxPosition,
		currentPosition: position{
			x: 0,
			y: 0,
		},
		timesPulled: 0,
	}
}

func (d *drone) work(instructions []int) {
	for d.currentPosition.y < d.maxPosition.y {
		for d.currentPosition.x < d.maxPosition.x {
			d.interpreter.Run(instructions, d, d)

			d.currentPosition.x++
		}
		d.currentPosition.x = 0
		d.currentPosition.y++
	}
}

func CountAffectedPoints(instructions []int) int {
	drone := newDrone(position{
		x: 50,
		y: 50,
	})
	drone.work(instructions)

	return drone.timesPulled
}
