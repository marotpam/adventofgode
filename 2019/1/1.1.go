package _2019

func CalculateFirstAmountOfFuel(masses []int) int {
	result := 0

	for _, m := range masses {
		result += m/3 - 2
	}

	return result
}
