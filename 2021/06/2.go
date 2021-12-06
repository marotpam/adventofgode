package _2021

func CountLanternfishAfterNDays2(lanternfish []int, nDays int) int {
	m := map[int]int{}
	for _, d := range lanternfish {
		m[d] += 1
	}
	return countLanternfishAfterNDays(m, nDays)
}

func countLanternfishAfterNDays(lanternfish map[int]int, nDays int) int {
	if nDays == 0 {
		total := 0
		for _, c := range lanternfish {
			total += c
		}
		return total
	}

	fishForNextCycle := map[int]int{}
	for daysLeft, c := range lanternfish {
		if daysLeft == 0 {
			fishForNextCycle[8] = c
			fishForNextCycle[6] += c
		} else {
			fishForNextCycle[daysLeft-1] += c
		}
	}
	return countLanternfishAfterNDays(fishForNextCycle, nDays-1)
}
