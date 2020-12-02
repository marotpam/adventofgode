package main

func CountValidPasswords2(rawInput string) int {
	passwords := parse(rawInput)

	c := 0
	for _, p := range passwords {
		if p.isValid2() {
			c++
		}
	}

	return c
}

func (p password) isValid2() bool {
	hasInFirst := rune(p.password[p.policy.minOccurrences-1]) == p.policy.letter
	hasInSecond := rune(p.password[p.policy.maxOccurrences-1]) == p.policy.letter

	if hasInFirst {
		return !hasInSecond
	}

	if hasInSecond {
		return !hasInFirst
	}

	return false
}
