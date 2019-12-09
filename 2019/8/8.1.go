package main

import "math"

type layer struct {
	digits map[rune]int
}

func newLayer() *layer {
	return &layer{digits: make(map[rune]int, 10)}
}

func CalculateFirstMultiplication(s string, width, height int) int {
	minCountOfZeros := math.MaxInt64
	layerSize := width * height

	l := newLayer()
	result := 0
	for i, r := range s {
		if i%layerSize == 0 {
			zerosInLayer := l.digits['0']
			if zerosInLayer > 0 && zerosInLayer < minCountOfZeros {
				result = l.digits['1']*l.digits['2']
				minCountOfZeros = zerosInLayer
			}
			l = newLayer()
		}
		l.digits[r] = l.digits[r] + 1
	}
	zerosInLayer := l.digits['0']
	if zerosInLayer > 0 && zerosInLayer < minCountOfZeros {
		result = l.digits['1']*l.digits['2']
		minCountOfZeros = zerosInLayer
	}
	return result
}
