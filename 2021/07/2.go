package _2021

import "math"

func CalculateMinimumFuel2(positions []int) int {
	crabCountsPerPosition := map[int]int{}
	minPosition, maxPosition := math.MaxInt64, 0
	for _, p := range positions {
		crabCountsPerPosition[p]++
		if p > maxPosition {
			maxPosition = p
		}
		if p < minPosition {
			minPosition = p
		}
	}

	minAmountOfFuel := math.MaxInt64
	for pos := minPosition; pos <= maxPosition; pos++ {
		totalFuelToReachPosition := 0
		for crabPosition, occurrences := range crabCountsPerPosition {
			distanceToPosition := abs(crabPosition - pos)
			// sum of natural numbers until n is (n*(n+1))/2
			fuelUntilDistance := (distanceToPosition * (distanceToPosition + 1)) / 2
			totalFuelToReachPosition += fuelUntilDistance * occurrences
		}
		if totalFuelToReachPosition < minAmountOfFuel {
			minAmountOfFuel = totalFuelToReachPosition
		}
	}
	return minAmountOfFuel
}
