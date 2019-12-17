package main

import (
	"fmt"

	"github.com/marotpam/adventofgode/2019/intcode"
)

const (
	newLine                 = '\n'
	elementScaffold         = '#'
	elementCrossing         = 'O'
	elementOpenSpace        = '.'
	elementRobotFacingUp    = '^'
	elementRobotFacingDown  = '^'
	elementRobotFacingRight = 'v'
	elementRobotFacingLeft  = '<'
)

type position struct {
	x, y int
}

func (p *position) crossSectionPositions() []position {
	return []position{
		{
			x: p.x + 1,
			y: p.y,
		},
		{
			x: p.x,
			y: p.y + 1,
		},
		{
			x: p.x - 1,
			y: p.y,
		},
		{
			x: p.x,
			y: p.y - 1,
		},
	}
}

type grid struct {
	layout  map[position]rune
	current position
	maxX    int
}

func (g *grid) isCrossing(p position) bool {
	if g.layout[p] != elementScaffold {
		return false
	}

	for _, otherPos := range p.crossSectionPositions() {
		if g.layout[otherPos] != elementScaffold {
			return false
		}
	}

	return true
}

func (g *grid) sumAlignmentParameters() int {
	ap := 0
	for pos := range g.layout {
		if g.isCrossing(pos) {
			ap += pos.x * pos.y
		}
	}
	return ap
}

func (g *grid) newLine() {
	if g.current.x > g.maxX {
		g.maxX = g.current.x
	}

	g.current = position{
		x: 0,
		y: g.current.y + 1,
	}
}

func (g *grid) addElement(c int) {
	g.layout[g.current] = rune(c)
	g.current.x++
}

func (g *grid) print() {
	for y := 0; y < g.current.y; y++ {
		for x := 0; x < g.maxX; x++ {
			pos := position{x: x, y: y}
			c := g.layout[pos]
			if g.isCrossing(pos) {
				c = elementCrossing
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

type vacuumRobot struct {
	interpreter *intcode.Interpreter
	grid        *grid
}

func (v *vacuumRobot) Read() int {
	panic("implement me")
}

func (v *vacuumRobot) Write(c int) {
	switch c {
	case newLine:
		v.grid.newLine()
	default:
		v.grid.addElement(c)
	}
}

func newVacuumRobot() *vacuumRobot {
	return &vacuumRobot{
		interpreter: intcode.NewInterpreter(),
		grid: &grid{
			layout: make(map[position]rune),
			current: position{
				x: 0,
				y: 0,
			},
		},
	}
}

func (v *vacuumRobot) readGrid(instructions []int) {
	v.interpreter.Run(instructions, v, v)
}

func SumAlignmentParameters(instructions []int) int {
	r := newVacuumRobot()
	r.readGrid(instructions)

	return r.grid.sumAlignmentParameters()
}
