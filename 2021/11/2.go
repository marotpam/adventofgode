package _2021

func CountStepsUntilSynchronization(rawInput string) int {
	cave := parse(rawInput)
	for s := 1; ; s++ {
		flashedCount := cave.countFlashes()
		if flashedCount == len(cave.octopuses) {
			return s
		}
	}
}
