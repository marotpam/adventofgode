package _2020

import (
	"strconv"
	"strings"
)

const (
	directionEast  = 90
	directionWest  = 270
	directionNorth = 0
	directionSouth = 180

	actionEast    = 'E'
	actionWest    = 'W'
	actionNorth   = 'N'
	actionSouth   = 'S'
	actionLeft    = 'L'
	actionRight   = 'R'
	actionForward = 'F'
)

var incrementsByAction = map[rune]position{
	actionEast:  {x: 1, y: 0},
	actionWest:  {x: -1, y: 0},
	actionNorth: {x: 0, y: -1},
	actionSouth: {x: 0, y: 1},
}

var incrementsByDirection = map[int]position{
	directionEast:  incrementsByAction[actionEast],
	directionWest:  incrementsByAction[actionWest],
	directionNorth: incrementsByAction[actionNorth],
	directionSouth: incrementsByAction[actionSouth],
}

type instruction struct {
	action   rune
	movement int
}

type position struct {
	x, y int
}

type ship struct {
	position  position
	direction int
}

func newShip() ship {
	return ship{
		position: position{
			x: 0,
			y: 0,
		},
		direction: directionEast,
	}
}

func (s *ship) navigate1(instructions []instruction) {
	for _, i := range instructions {
		s.navigateInstruction(i)
	}

}

func (s *ship) navigateInstruction(i instruction) {
	switch i.action {
	case actionLeft:
		s.direction = s.direction - i.movement
		if s.direction < 0 {
			s.direction += 360
		}
	case actionRight:
		s.direction = (s.direction + i.movement) % 360
	case actionForward:

		s.position.move(incrementsByDirection[s.direction], i.movement)
	default:
		s.position.move(incrementsByAction[i.action], i.movement)
	}
}

func (s *ship) manhattanDistance() int {
	return abs(s.position.x) + abs(s.position.y)
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func (p *position) move(other position, movement int) {
	p.x += other.x * movement
	p.y += other.y * movement
}

func CalculateManhattanDistance1(rawInput string) int {
	instructions := parse(rawInput)

	s := newShip()
	s.navigate1(instructions)

	return s.manhattanDistance()
}

func parse(input string) []instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, 0, len(lines))

	for _, l := range lines {
		m, _ := strconv.Atoi(l[1:])
		instructions = append(instructions, instruction{
			action:   rune(l[0]),
			movement: m,
		})
	}

	return instructions
}
