package main

import (
	"strconv"
	"strings"
)

func SumCalibrationValuesPart1(rawInput string) int {
	sum := 0
	for _, line := range strings.Split(rawInput, "\n") {
		sum += GetCalibrationValue(line)
	}
	return sum
}

func GetCalibrationValue(line string) int {
	first, last := 0, 0
	for _, c := range line {
		n, _ := strconv.Atoi(string(c))
		if first == 0 {
			first = n
		}

		if n != 0 || last == 0 {
			last = n
		}
	}
	return first*10 + last
}
