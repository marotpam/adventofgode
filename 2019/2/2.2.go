package _2019

const WantedOutput = 19690720

func CalculateSecondOpcode(ints []int) []int {
	for i := 0; i <= 99; i++ {
		for j := i; j <= 99; j++ {
			newInts := append([]int{}, ints...)
			newInts[1] = i
			newInts[2] = j

			if CalculateFirstOpcode(newInts)[0] == WantedOutput {
				return newInts
			}
		}
	}
	return []int{}
}
