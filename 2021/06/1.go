package _2021

func CountLanternfishAfterNDays(lanternfish []int, nDays int) int {
	if nDays == 0 {
		return len(lanternfish)
	}

	var fishesForNextCycle []int
	for _, f := range lanternfish {
		if f == 0 {
			fishesForNextCycle = append(fishesForNextCycle, 6, 8)
		} else {
			fishesForNextCycle = append(fishesForNextCycle, f-1)
		}
	}
	return CountLanternfishAfterNDays(fishesForNextCycle, nDays-1)
}
