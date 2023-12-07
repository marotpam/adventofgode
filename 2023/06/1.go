package _2023

import (
	"strconv"
	"strings"
)

type race struct {
	time, distance int
}

func (r race) countNumberOfWaysForRecord() int {
	count := 0
	for i := 0; i <= r.time; i++ {
		res := (r.time - i) * i
		if res > r.distance {
			count++
		}
	}

	return count
}

func MultiplyNumberOfWaysToBeatAllRecords(rawInput string) int {
	result := 1
	races := parseRaces(rawInput)

	for _, r := range races {
		result *= r.countNumberOfWaysForRecord()
	}

	return result
}

func parseRaces(rawInput string) []race {
	races := make([]race, 0)
	lines := strings.Split(rawInput, "\n")

	times := strings.Fields(strings.Split(lines[0], ": ")[1])
	distances := strings.Fields(strings.Split(lines[1], ": ")[1])

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		races = append(races, race{
			time:     time,
			distance: distance,
		})
	}
	return races
}
