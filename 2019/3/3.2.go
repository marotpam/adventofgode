package main

import (
	"math"
	"strconv"
)

type pointWithSteps struct {
	x, y, steps int
}

func FindSecondDistanceToIntersectionPoint(first, second []string) int {
	firstPoints := findPointsWithStepsInPath(first)
	secondPoints := findPointsWithStepsInPath(second)

	minDistance := math.MaxInt64
	for _, p := range firstPoints {
		s := findIntersectionPoint(secondPoints, p)
		if s != nil {
			d := p.steps + s.steps
			if d < minDistance {
				minDistance = d
			}
		}
	}
	return minDistance
}

func findPointsWithStepsInPath(directions []string) []pointWithSteps {
	var points []pointWithSteps
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
			points = append(points, pointWithSteps{
				x:     x,
				y:     y,
				steps: totalSteps,
			})
		}
	}
	return points
}

func findIntersectionPoint(points []pointWithSteps, searched pointWithSteps) *pointWithSteps {
	for _, p := range points {
		if p.x == searched.x && p.y == searched.y {
			return &p
		}
	}
	return nil
}
