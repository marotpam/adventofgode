package main

func MultiplyTreesInSlopes(rawInput string, slopes []slope) int {
	c := 1
	for _, s := range slopes {
		res := CountTreesInSlope(rawInput, s)
		c *= res
	}
	return c
}
