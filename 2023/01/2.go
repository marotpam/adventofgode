package main

import (
	"strings"
)

var numbersAsStrings = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func SumCalibrationValuesPart2(rawInput string) int {
	sum := 0
	for _, line := range strings.Split(rawInput, "\n") {
		sum += GetCalibrationValue(StringsToNumbers(line))
	}
	return sum
}

func StringsToNumbers(line string) string {
	res := ""

OUTER:
	for i, c := range line {
		if c >= '0' && c <= '9' {
			res += string(c)
		}

		remaining := line[i:]
		for str, num := range numbersAsStrings {
			if strings.HasPrefix(remaining, str) {
				res += string(num)
				continue OUTER
			}
		}
	}

	return res
}
