package main

func FindMultiplicationFor3(numbers []int) int {
	seen := make(map[int]int, len(numbers)*2)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			seen[numbers[i]+numbers[j]] = numbers[i] * numbers[j]
		}
	}

	for _, n := range numbers {
		otherSum := 2020 - n
		otherProduct, ok := seen[otherSum]
		if ok {
			return n * otherProduct
		}
	}

	return -1
}
