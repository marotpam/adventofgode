package _2021

import (
	"strconv"
	"strings"
)

func CalculatePowerConsumption(rawInput string) int {
	gamma, epsilon := parseRates(rawInput)
	gammaRate, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaRate) * int(epsilonRate)
}

func parseRates(rawInput string) (string, string) {
	lines := strings.Split(rawInput, "\n")
	commonBits, leastCommonBits := "", ""
	if len(lines) == 0 {
		return commonBits, leastCommonBits
	}

	for i := 0; i < len(lines[0]); i++ {
		ones := 0
		for j := 0; j < len(lines); j++ {
			if lines[j][i] == '1' {
				ones++
			}
		}
		if ones >= len(lines)/2 {
			commonBits += "1"
			leastCommonBits += "0"
		} else {
			commonBits += "0"
			leastCommonBits += "1"
		}
	}
	return commonBits, leastCommonBits
}
