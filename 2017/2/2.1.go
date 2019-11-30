package main

const maxUint = ^uint(0)
const maxInt = int(maxUint >> 1)

type matrix [][]int

func SolveFirstChecksum(m matrix) int {
	result := 0

	for i := 0; i < len(m); i++ {
		min := maxInt
		max := 0
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] > max {
				max = m[i][j]
			}
			if m[i][j] < min {
				min = m[i][j]
			}
		}
		result += max - min
	}

	return result
}
