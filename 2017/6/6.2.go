package _2017

func CountRedistributionCyclesSecond(ints []int) int {
	cycles := make(map[string]int, 0)
	for c := 0; ; c++ {
		key := serialize(ints)
		cycle, ok := cycles[key]
		if ok {
			return c - cycle
		}
		cycles[key] = c

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
