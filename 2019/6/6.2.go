package main

import "math"

const infiniteDistance = -1

func (o *object) countMinimalDistanceTo(other object) int {
	if other.name == o.name {
		return 0
	}

	if len(o.orbits) == 0 {
		return infiniteDistance
	}

	min := math.MaxInt64
	for _, orbit := range o.orbits {
		d := orbit.countMinimalDistanceTo(other)
		if d == infiniteDistance {
			continue
		}

		if d < min {
			min = d
		}
	}

	if min == math.MaxInt64 {
		return infiniteDistance
	}

	return min + 1
}

func CountOrbitalTransfers(orbitMap []string) int {
	m := getObjectMap(orbitMap)

	me := m["YOU"]
	santa := m["SAN"]

	delete(m, "YOU")
	delete(m, "SAN")

	min := math.MaxInt64
	for _, o := range m {
		myDistance := me.countMinimalDistanceTo(*o)
		if myDistance == infiniteDistance {
			continue
		}
		santasDistance := santa.countMinimalDistanceTo(*o)
		if santasDistance == infiniteDistance {
			continue
		}

		d := myDistance + santasDistance - 2
		if d < min {
			min = d
		}
	}

	return min
}
