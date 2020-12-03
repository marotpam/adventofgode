package _2019

import (
	"reflect"
	"testing"
)

func TestTotalEnergyOfAMoon(t *testing.T) {
	m := &moon{
		coordinates: dimensions{
			x: 2,
			y: 1,
			z: 3,
		},
		velocities: dimensions{
			x: -3,
			y: -2,
			z: 1,
		},
	}
	expected := 36
	if got := m.totalEnergy(); got != expected {
		t.Errorf("totalEnergy() = %v, want %v", got, expected)
	}
}

func TestApplyingVelocityToAMoon(t *testing.T) {
	m := &moon{
		coordinates: dimensions{
			x: 1,
			y: 2,
			z: 3,
		},
		velocities: dimensions{
			x: -2,
			y: 0,
			z: 3,
		},
	}

	m.applyVelocity()

	expectedCoordinates := dimensions{
		x: -1,
		y: 2,
		z: 6,
	}
	if !reflect.DeepEqual(expectedCoordinates, m.coordinates) {
		t.Errorf("coordinates after applying velocities = %+v, want %+v", m.coordinates, expectedCoordinates)
	}
}

func TestApplyingGravityBetweenTwoMoons(t *testing.T) {
	ganymede := newMoon(dimensions{
		x: 3,
		y: 5,
		z: 10,
	})

	europa := newMoon(dimensions{
		x: 5,
		y: 3,
		z: 10,
	})

	ganymede.applyGravityWith(europa)

	expectedVelocitiesForGanymede := dimensions{
		x: 1,
		y: -1,
		z: 0,
	}
	if !reflect.DeepEqual(expectedVelocitiesForGanymede, ganymede.velocities) {
		t.Errorf("ganymede velocities after applying velocities = %+v, want %+v", ganymede.velocities, expectedVelocitiesForGanymede)
	}

	expectedVelocitiesForEuropa := dimensions{
		x: -1,
		y: 1,
		z: 0,
	}
	if !reflect.DeepEqual(expectedVelocitiesForEuropa, europa.velocities) {
		t.Errorf("europa velocities after applying velocities = %+v, want %+v", europa.velocities, expectedVelocitiesForEuropa)
	}
}

func TestCalculateTotalEnergyForAllMoons(t *testing.T) {
	t.Run("given input", func(t *testing.T) {
		moonCoordinates := []dimensions{
			{
				x: 5,
				y: -1,
				z: 5,
			},
			{
				x: 0,
				y: -14,
				z: 2,
			},
			{
				x: 16,
				y: 4,
				z: 0,
			},
			{
				x: 18,
				y: 1,
				z: 16,
			},
		}

		actual := CalculateTotalEnergy(moonCoordinates, 1000)
		expected := 7928

		if actual != expected {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	})
	t.Run("first example", func(t *testing.T) {
		moonCoordinates := []dimensions{
			{
				x: -1,
				y: 0,
				z: 2,
			},
			{
				x: 2,
				y: -10,
				z: -7,
			},
			{
				x: 4,
				y: -8,
				z: 8,
			},
			{
				x: 3,
				y: 5,
				z: -1,
			},
		}

		actual := CalculateTotalEnergy(moonCoordinates, 10)
		expected := 179

		if actual != expected {
			t.Errorf("expected %d but got %d", expected, actual)
		}
	})
}
