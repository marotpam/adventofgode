package _2020

import "math"

func GetEncryptionWeakness(rawInput string, preambleLength int) int {
	ints := parse(rawInput)

	for i, n := range ints[preambleLength:] {
		if isMissingInSums(n, ints[i:i+preambleLength]) {
			return sumSmallestAndLargest(getContiguousSetAmountingTo(n, ints))
		}
	}

	return -1
}

func getContiguousSetAmountingTo(n int, ints []int) []int {
	lo, hi := 0, 1
	sum := ints[lo]

	for {
		sum += ints[hi]
		if sum == n {
			return ints[lo : hi+1]
		}

		if sum < n {
			hi++
			continue
		}

		if ints[hi] > n {
			lo = hi
		}

		lo++
		sum = ints[lo]
		hi = lo + 1
	}
}

func sumSmallestAndLargest(ints []int) int {
	min, max := math.MaxInt64, 0

	for _, i := range ints {
		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	return min + max
}
