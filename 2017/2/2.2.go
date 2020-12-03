package _2017

import "sort"

func SolveSecondChecksum(m matrix) int {
	result := 0

	for i := 0; i < len(m); i++ {
		result += aux(m[i])
	}

	return result
}

func aux(row []int) int {
	sort.Ints(row)
	for i := len(row)-1; i > 0; i-- {
		for j := i-1; j >= 0; j-- {
			if row[i] % row[j] == 0 {
				return row[i] / row[j]
			}
		}
	}

	return 0
}
