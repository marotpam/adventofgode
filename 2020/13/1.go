package _2020

import (
	"math"
	"strconv"
	"strings"
)

func SolvePart1(earliestTimestamp int, rawBusIDs string) int {
	minTimeDifference := math.MaxInt64
	minBusID := 0

	for _, rawID := range strings.Split(rawBusIDs, ",") {
		if rawID == "x" {
			continue
		}

		id, _ := strconv.Atoi(rawID)
		d := findTimeDiff(earliestTimestamp, id)
		if d < minTimeDifference {
			minTimeDifference = d
			minBusID = id
		}
	}

	return minTimeDifference * minBusID
}

func findTimeDiff(timestamp, id int) int {
	div := timestamp/id

	return id*(div+1)%timestamp
}
