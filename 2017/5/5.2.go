package main

func CountJumpsSecond(ints []int) int {
	previous, jumps := 0, 0
	for i := 0; i < len(ints); jumps++ {
		inc := 1
		if ints[i] >= 3 {
			inc = -1
		}
		previous = i
		i += ints[i]

		ints[previous] += inc

	}
	return jumps
}
