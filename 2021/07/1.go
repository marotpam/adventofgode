package _2021

import "math"

func CalculateMinimumFuel1(positions []int) int {
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
		sum := 0
		for position, occurrences := range crabCountsPerPosition {
			sum += abs(position-pos) * occurrences
		}
		if sum < minAmountOfFuel {
			minAmountOfFuel = sum
		}
	}
	return minAmountOfFuel
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
