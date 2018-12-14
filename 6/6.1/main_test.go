package main

import (
	"testing"
)

func TestWhenGridOnlyHasOneRowWithTwoElements(t *testing.T) {
	ps := []coordinate{
		{0, 0},
		{3, 0},
	}

	g := newGrid(ps)

	expectedClosestPoints := []struct{
		x, y, closestPoint int
	} {
		{0, 0, 0},
		{1, 0, 0},
		{2, 0, 1},
		{3, 0, 1},
	}

	for _, e := range expectedClosestPoints {
		if g.closestPoints()[e.x][e.y] != e.closestPoint {
			t.Fatalf("Failed asserting closest point to %d, %d was %d, got %d", e.x, e.y, e.closestPoint, g.closestPoints()[e.x][e.y])
		}
	}
}