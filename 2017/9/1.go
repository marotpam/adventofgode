package _2017

func CalculateTotalScore(input string) int {
	total, c := 0, 0

	for _, r := range cleanGarbage(input) {
		if r == '{' {
			c++
		} else if c > 0 {
			total += c
			c--
		}
	}

	return total
}

func cleanGarbage(s string) string {
	clean := ""
	ignore := false

	for i := 0; i < len(s); i++ {
		r := s[i]
		switch r {
		case '{', '}':
			if !ignore {
				clean = clean + string(r)
			}
		case '<':
			ignore = true
		case '>':
			ignore = false
		case '!':
			i++
		}
	}
	return clean
}
