package main

import (
	"strconv"
	"strings"
)

func CountRedistributionCycles(ints []int) int {
	cycles := make(map[string]struct{}, 0)
	for c := 0; ; c++ {
		key := serialize(ints)
		_, ok := cycles[key]
		if ok {
			return c
		}
		cycles[key] = struct{}{}

		maxIndex, maxValue := 0, 0
		for k, v := range ints {
			if v > maxValue {
				maxValue = v
				maxIndex = k
			}
		}
		ints[maxIndex] = 0

		for i := 0; i < maxValue; i++ {
			ints[(maxIndex+i+1)%len(ints)]++
		}
	}
	return -1
}

func serialize(ints []int) string {
	s := make([]string, 0, len(ints))

	for _, i := range ints {
		s = append(s, strconv.Itoa(i))
	}

	return strings.Join(s, ",")
}
