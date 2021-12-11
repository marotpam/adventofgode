package _2021

import "strings"

func CountUniqueDigits(rawInput string) int {
	uniqueSegmentsInDigits := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}

	count := 0
	entries := parse(rawInput)
	for _, e := range entries {
		for _, outputDigit := range e.output {
			_, ok := uniqueSegmentsInDigits[len(outputDigit)]
			if ok {
				count++
			}
		}
	}
	return count
}

type entry struct {
	signalPatterns, output []string
}

func parse(input string) []entry {
	lines := strings.Split(input, "\n")
	entries := make([]entry, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		entries = append(entries, entry{
			signalPatterns: strings.Split(parts[0], " "),
			output:         strings.Split(parts[1], " "),
		})
	}
	return entries
}
