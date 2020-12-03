package _2019

const (
	opAdd                = 1
	opMultiply           = 2
	opInput              = 3
	opOutput             = 4
	opJumpIfTrue         = 5
	opJumpIfFalse        = 6
	opLessThan           = 7
	opEquals             = 8
	opAdjustRelativeBase = 9
	opHalt               = 99

	paramModePosition     = 0
	paramModeImmediate    = 1
	paramModeRelativeBase = 2
)

type input interface {
	read() int
}

type output interface {
	write(int)
}

func CalculateSecondOpcode(relativeBase int, ints []int, in input, out output) map[int]int {
	memory := map[int]int{}
	for i, v := range ints {
		memory[i] = v
	}

	opCode := 0
	for i := 0; opCode != opHalt; {
		opCode = memory[i] % 100
		instruction := memory[i] / 100

		operandCount := 1
		switch opCode {
		case opAdd, opMultiply, opEquals, opLessThan:
			operandCount = 3
		case opJumpIfTrue, opJumpIfFalse:
			operandCount = 2
		}

		operands := make([]int, operandCount)
		offset := memory[i+operandCount]
		for o := 0; o < operandCount; o++ {
			p := memory[i+o+1]
			switch instruction % 10 {
			case paramModePosition:
				p = memory[p]
			case paramModeRelativeBase:
				p = memory[p+relativeBase]
				if o == operandCount-1 {
					offset += relativeBase
				}
			}

			operands[o] = p
			instruction = instruction / 10
		}
		newPosition := -1

		switch opCode {
		case opAdd, opMultiply, opEquals, opLessThan:
			memory[offset] = calculateSecondArithmeticOp(opCode, operands[0], operands[1])
		case opJumpIfTrue:
			if operands[0] != 0 {
				newPosition = operands[1]
			}
		case opJumpIfFalse:
			if operands[0] == 0 {
				newPosition = operands[1]
			}
		case opInput:
			memory[offset] = in.read()
		case opOutput:
			out.write(operands[0])
		case opAdjustRelativeBase:
			relativeBase += operands[0]
		}
		i += len(operands) + 1

		if newPosition >= 0 {
			i = newPosition
		}
	}
	return memory
}

func calculateSecondArithmeticOp(op, a, b int) int {
	switch op {
	case opAdd:
		return a + b
	case opMultiply:
		return a * b
	case opEquals:
		if a == b {
			return 1
		}
	case opLessThan:
		if a < b {
			return 1
		}
	}
	return 0
}
