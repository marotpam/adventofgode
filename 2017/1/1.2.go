package _2017

func SolveSecondCaptcha(c string) int {
	result := 0
	len := len(c)
	increment := len/2

	for i := 0; i < len; i++ {
		a := intAtPosition(c, i)
		b := intAtPosition(c, (i+increment)%len)
		if a == b {
			result += a
		}
 	}

	return result
}