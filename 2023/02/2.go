package _2023

func (g game) getPower() int {
	scores := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range g {
		for _, cubePick := range set {
			if scores[cubePick.colour] < cubePick.count {
				scores[cubePick.colour] = cubePick.count
			}
		}
	}

	res := 1
	for _, score := range scores {
		res *= score
	}
	return res
}

func SumGamePowers(rawInput string) int {
	sum := 0
	for _, game := range parseGames(rawInput) {
		sum += game.getPower()
	}
	return sum
}
