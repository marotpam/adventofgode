package main

import (
	"math"
	"strconv"
)

type pointWithSteps struct {
	point
	steps int
}

func FindSecondDistanceToIntersectionPoint(first, second []string) int {
	firstPoints := findPointsWithStepsInPath(first)
	secondPoints := findPointsWithStepsInPath(second)

	minDistance := math.MaxInt64
	for _, p := range firstPoints {
		s, ok := secondPoints[p.key()]
		if ok {
			d := p.steps + s.steps
			if d < minDistance {
				minDistance = d
			}
		}
	}
	return minDistance
}

func findPointsWithStepsInPath(directions []string) map[string]pointWithSteps {
	points := make(map[string]pointWithSteps, 0)
	x, y := 0, 0
	totalSteps := 0
	for _, d := range directions {
		steps, _ := strconv.Atoi(d[1:])
		incX, incY := 0, 0
		switch d[0] {
		case directionUp:
			incY = 1
		case directionDown:
			incY = -1
		case directionLeft:
			incX = -1
		case directionRight:
			incX = 1
		}
		for i := 0; i < steps; i++ {
			totalSteps++
			x += incX
			y += incY
			p := pointWithSteps{
				point: point{
					x: x,
					y: y,
				},
				steps: totalSteps,
			}
			points[p.key()] = p
		}
	}
	return points
}
