package _2020

func newWaypoint() *ship {
	return &ship{
		position:  position{
			x: 10,
			y: -1,
		},
		direction: 0,
	}
}

func (s *ship) navigate2(instructions []instruction, waypoint *ship) {
	for _, i := range instructions {
		switch i.action {
		case actionLeft:
			waypoint.position.rotateLeft(i.movement)
		case actionRight:
			waypoint.position.rotateRight(i.movement)
		case actionForward:
			s.position.move(waypoint.position, i.movement)
		default:
			waypoint.navigateInstruction(i)
		}
	}
}

func (p *position) rotateLeft(movement int) {
	switch movement {
	case 90:
		p.x, p.y = p.y, -p.x
	case 180:
		p.x, p.y = -p.x, -p.y
	case 270:
		p.x, p.y = -p.y, p.x
	}
}

func (p *position) rotateRight(movement int) {
	switch movement {
	case 90:
		p.x, p.y = -p.y, p.x
	case 180:
		p.x, p.y = -p.x, -p.y
	case 270:
		p.x, p.y = p.y, -p.x
	}
}

func CalculateManhattanDistance2(rawInput string) int {
	instructions := parse(rawInput)

	s := newShip()
	wp := newWaypoint()
	s.navigate2(instructions, wp)

	return s.manhattanDistance()
}
