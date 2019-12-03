package main

import (
	"math"
	"strconv"
)

const (
	directionUp    = 'U'
	directionDown  = 'D'
	directionLeft  = 'L'
	directionRight = 'R'
)

type point struct {
	x, y int
}

func (p *point) hammingDistance() int {
	return abs(p.x) + abs(p.y)
}

func abs(i int) int {
	if i > 0 {
		return i
	}
	return i * -1
}

func FindFirstDistanceToIntersectionPoint(first, second []string) int {
	firstPoints := findPointsInPath(first)
	secondPoints := findPointsInPath(second)

	intersections := findIntersections(firstPoints, secondPoints)
	minDistance := math.MaxInt64
	for _, i := range intersections {
		hd := i.hammingDistance()
		if hd < minDistance {
			minDistance = hd
		}
	}
	return minDistance
}

func findPointsInPath(directions []string) []point {
	var points []point
	x, y := 0, 0
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
			x += incX
			y += incY
			points = append(points, point{
				x: x,
				y: y,
			})
		}
	}
	return points
}

func findIntersections(first, second []point) []point {
	intersections := []point{}
	for _, p := range first {
		if containsPoint(second, p) {
			intersections = append(intersections, p)
		}
	}
	return intersections
}

func containsPoint(points []point, searched point) bool {
	for _, p := range points {
		if p.x == searched.x && p.y == searched.y {
			return true
		}
	}
	return false
}
