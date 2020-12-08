package _2020

func GetAccumulatorAfterSuccessfulFinish(rawInput string) int {
	instructions := parse(rawInput)

	c := console{}
	ok := c.processUntilSuccess(instructions, opJump, opNoOp)
	if !ok {
		c.processUntilSuccess(instructions, opNoOp, opJump)
	}
	return c.accumulator
}

func (c *console) processUntilSuccess(instructions []instruction, fromOp, toOp string) bool {
	positionsWithJumps := findPositionsWith(fromOp, instructions)
	for _, p := range positionsWithJumps {
		modified := make([]instruction, len(instructions))
		copy(modified, instructions)
		modified[p].operation = toOp

		if c.process(modified) {
			return true
		}
	}

	return false
}

func (c *console) reset() {
	c.pointer = 0
	c.accumulator = 0
}

func findPositionsWith(wantedOp string, instructions []instruction) []int {
	positions := make([]int, 0, len(instructions))

	for i, inst := range instructions {
		if inst.operation == wantedOp {
			positions = append(positions, i)
		}
	}

	return positions
}
