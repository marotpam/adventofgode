package main

import "strconv"

func SolveFirstCaptcha(c string) int {
	result := 0
	len := len(c)

	for i := 0; i < len; i++ {
		a := intAtPosition(c, i)
		b := intAtPosition(c, (i+1)%len)
		if a == b {
			result += a
		}
 	}

	return result
}

func intAtPosition(s string, i int) int {
	a, _ := strconv.Atoi(string(s[i]))
	return a
}