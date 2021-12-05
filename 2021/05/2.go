package _2021

func CountOverlappingHydrothermalVentLines2(rawInput string) int {
	points := map[coordinate]int{}
	overlappingCount := 0

	lines := parseLines(rawInput)
	for _, l := range lines {
		coords := l.start.getCoordinatesUntil2(l.end)
		for _, c := range coords {
			if _, ok := points[c]; !ok {
				points[c] = 0
			}
			points[c]++

			if points[c] == 2 {
				overlappingCount++
			}
		}
	}

	return overlappingCount
}

func (c coordinate) getCoordinatesUntil2(other coordinate) []coordinate {
	var coords []coordinate

	incrX, incrY, canReach := c.getIncrements2(other)
	if !canReach {
		return coords
	}

	pos := c
	for {
		coords = append(coords, pos)
		if pos == other {
			return coords
		}

		pos.x += incrX
		pos.y += incrY
	}
}

func (c coordinate) getIncrements2(other coordinate) (int, int, bool) {
	if c == other {
		return 0, 0, true
	}

	incrX, incrY := other.x-c.x, other.y-c.y
	if incrX != 0 && incrY == 0 {
		return normalise(incrX), 0, true
	}

	if incrX == 0 && incrY != 0 {
		return 0, normalise(incrY), true
	}

	if abs(incrX) == abs(incrY) {
		return normalise(incrX), normalise(incrY), true
	}

	return 0, 0, false
}
