package _2020

import (
	"sort"
)

func CountDistinctArrangements(rawInput string) int {
	ints := parse(rawInput)
	sort.Ints(ints)

	accum := map[int]int{0: 1}

	for _, i := range ints {
		accum[i] = accum[i-1] + accum[i-2] + accum[i-3]
	}

	return accum[ints[len(ints)-1]]
}
