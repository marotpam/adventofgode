package _2017

func GetMaxAllocated(rawInput string) int {
	instructions := parse(rawInput)

	computer := newComputer()
	computer.process(instructions)

	return computer.maxAllocatedValue
}