package main

func CountSecondPasswords(a, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		if matchesSecondPassword(i) {
			count++
		}
	}
	return count
}

func matchesSecondPassword(n int) bool {
	hasTwoEqual := false
	for ; n > 10; {
		penultimate, last := n/10%10, n%10
		if penultimate > last {
			return false
		}

		if penultimate < last {
			n = n / 10
			continue
		}

		var sameDigitsCounter int
		for sameDigitsCounter = 1; last == penultimate; sameDigitsCounter++ {
			n = n / 10
			penultimate, last = n/10%10, n%10
		}

		if sameDigitsCounter == 2 {
			hasTwoEqual = true
		}
	}
	return hasTwoEqual
}
