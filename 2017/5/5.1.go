package main

func CountJumps(ints []int) int {
	previous, jumps := 0, 0
	for i := 0; i < len(ints); jumps++ {
		previous = i
		i += ints[i]
		ints[previous]++
	}
	return jumps
}
