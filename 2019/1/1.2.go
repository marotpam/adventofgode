package main

func CalculateSecondAmountOfFuel(masses []int) int {
	result := 0

	for _, m := range masses {
		f := m/3 - 2
		for ; f > 0; {
			result += f
			f = f/3 - 2
		}
	}

	return result
}
