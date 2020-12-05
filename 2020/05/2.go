package _2020

import "strings"

func FindMySeatID(rawInput string) int {
	lines := strings.Split(rawInput, "\n")
	assignedIDs := make(map[int]string, len(lines))

	for _, line := range lines {
		assignedIDs[getSeat(line).getID()] = line
	}

	var i int
	for i = 0; assignedIDs[i] == ""; i++ {
	}

	for ; assignedIDs[i] != ""; i++ {
	}

	return i
}
