package _2020

func FindMultiplicationFor2(numbers []int) int {
	seen := make(map[int]bool, len(numbers))
	for _, n := range numbers {
		other := 2020 - n
		_, ok := seen[other]
		if ok {
			return other * n
		}
		seen[n] = true
	}

	return -1
}
