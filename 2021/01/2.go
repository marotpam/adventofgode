package _2021

const windowSize = 3

func CountDepthIncreasesWithWindows(rawInput string) int {
	depths := parseDepths(rawInput)
	if len(depths) < 2*windowSize {
		return 0
	}

	increases := 0
	for i := 0; i < len(depths)-windowSize; i++ {
		if windowSum(depths, i+1) > windowSum(depths, i) {
			increases++
		}
	}
	return increases
}

func windowSum(depths []int, i int) int {
	s := 0
	for j := 0; j < 3; j++ {
		s += depths[i+j]
	}
	return s
}
