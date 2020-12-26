package _2020

func CalculateEncryptionKey(cardPK, doorPK int) int {
	cardLoops := calculateLoopSize(7, cardPK)

	return transform(doorPK, cardLoops)
}

func calculateLoopSize(subjectNumber, wantedPK int) int {
	res, i := 1, 1
	for {
		res *= subjectNumber
		res = res%20201227
		if res == wantedPK {
			return i
		}
		i++
	}
}

func transform(subjectNumber, loopSize int) int {
	res := 1
	for i := 0; i < loopSize; i++ {
		res *= subjectNumber
		res = res%20201227
	}
	return res
}
