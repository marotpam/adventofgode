package _2020

func CountValidPasswords2(rawInput string) int {
	lines := parse(rawInput)

	c := 0
	for _, l := range lines {
		if l.password.isValid2(l.policy) {
			c++
		}
	}

	return c
}

func (p password) isValid2(policy policy) bool {
	hasInFirst := p.hasRuneInPosition(policy.letter, policy.minOccurrences - 1)
	hasInSecond := p.hasRuneInPosition(policy.letter, policy.maxOccurrences - 1)

	return hasInFirst != hasInSecond
}

func (p password) hasRuneInPosition(r rune, pos int) bool {
	return rune(p[pos]) == r
}
