package _2017

func CountNonCanceledInGarbage(input string) int {
	c := 0
	ignoring := false

	for i := 0; i < len(input); i++ {
		r := input[i]
		switch r {
		case '<':
			if ignoring {
				c++
			}
			ignoring = true
		case '>':
			ignoring = false
		case '!':
			i++
		default:
			if ignoring {
				c++
			}
		}
	}

	return c
}
