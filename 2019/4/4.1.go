package main

func CountFirstPasswords(a, b int) int {
	count := 0
	for i := a; i <= b; i++ {
		if matchesFirstPassword(i) {
			count++
		}
	}
	return count
}

func matchesFirstPassword(n int) bool {
	hasTwoEqual := false
	for ; n > 10; {
		last, penultimate := n%10, n/10%10
		if last < penultimate {
			return false
		}

		if last == penultimate {
			hasTwoEqual = true
		}

		n = n / 10
	}
	return hasTwoEqual
}