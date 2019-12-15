package main

type dimensions struct {
	x, y, z int
}

type moon struct {
	coordinates dimensions
	velocities  dimensions
}

func newMoon(coordinates dimensions) *moon {
	return &moon{
		coordinates: coordinates,
		velocities: dimensions{
			x: 0,
			y: 0,
			z: 0,
		},
	}
}

func (m *moon) applyVelocity() {
	m.coordinates.x += m.velocities.x
	m.coordinates.y += m.velocities.y
	m.coordinates.z += m.velocities.z
}

func (m *moon) applyGravityWith(other *moon) {
	incX := getIncrement(m.coordinates.x, other.coordinates.x)
	m.velocities.x += incX
	other.velocities.x -= incX

	incY := getIncrement(m.coordinates.y, other.coordinates.y)
	m.velocities.y += incY
	other.velocities.y -= incY

	incZ := getIncrement(m.coordinates.z, other.coordinates.z)
	m.velocities.z += incZ
	other.velocities.z -= incZ
}

func getIncrement(a, b int) int {
	diff := a - b
	if diff > 0 {
		return -1
	}
	if diff < 0 {
		return 1
	}

	return 0
}

func (m *moon) totalEnergy() int {
	return m.potentialEnergy() * m.kineticEnergy()
}

func (m *moon) potentialEnergy() int {
	return abs(m.coordinates.x) + abs(m.coordinates.y) + abs(m.coordinates.z)
}

func (m *moon) kineticEnergy() int {
	return abs(m.velocities.x) + abs(m.velocities.y) + abs(m.velocities.z)
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func CalculateTotalEnergy(moonCoordinates []dimensions, steps int) int {
	moons := make([]*moon, 0, len(moonCoordinates))
	for _, c := range moonCoordinates {
		moons = append(moons, newMoon(c))
	}

	for s := 0; s < steps; s++ {
		for i := 0; i < len(moons); i++ {
			for j := i + 1; j < len(moons); j++ {
				moons[i].applyGravityWith(moons[j])
			}
		}

		for _, m := range moons {
			m.applyVelocity()
		}
	}

	t := 0

	for _, m := range moons {
		t += m.totalEnergy()
	}

	return t
}
