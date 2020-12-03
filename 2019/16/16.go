package _2019

import (
	"strconv"
	"strings"
)

func CalculateFlawedFrequencyTransmission(signal string, basePattern []int, phases int) string {
	fft := signal

	for p := 0; p < phases; p++ {
		fft = applyPhase(fft, basePattern)
	}

	return fft
}

func applyPhase(signal string, basePattern []int) string {
	signalInts := make([]int, 0, len(signal))
	for _, c := range strings.Split(signal, "") {
		number, _ := strconv.Atoi(c)
		signalInts = append(signalInts, number)
	}

	parts := ""
	for i := range signal {
		pattern := getPattern(basePattern, signal, i+1)
		digit := 0
		for j := 0; j < len(pattern); j++ {
			digit += pattern[j] * signalInts[j]
		}
		parts += strconv.Itoa(abs(digit) % 10)
	}

	return parts
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func getPattern(basePattern []int, signal string, position int) []int {
	result := make([]int, 0, len(signal))
	for i := 0; len(result) <= len(signal); i++ {
		for j := 0; j < position; j++ {
			result = append(result, basePattern[i%len(basePattern)])
		}
	}

	return result[1 : len(signal)+1]
}
