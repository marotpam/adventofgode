package main

const (
	opJumpIfTrue  = 5
	opJumpIfFalse = 6
	opLessThan    = 7
	opEquals      = 8
)

func CalculateSecondOpcode(ints []int, in input, out output) []int {
	for i := 0; ints[i]%100 != opHalt; {
		opCode := ints[i] % 100
		parameters := getSecondParameters(ints, i)
		newPosition := -1

		switch opCode {
		case opAdd, opMultiply, opEquals, opLessThan:
			ints[ints[i+3]] = calculateSecondArithmeticOp(opCode, parameters[0], parameters[1])
		case opJumpIfTrue:
			if parameters[0] != 0 {
				newPosition = parameters[1]
			}
		case opJumpIfFalse:
			if parameters[0] == 0 {
				newPosition = parameters[1]
			}
		case opInput:
			ints[ints[i+1]] = in.read()
		case opOutput:
			out.write(parameters[0])
		}
		i += len(parameters) + 1

		if newPosition >= 0 {
			i = newPosition
		}
	}
	return ints
}

func getSecondParameters(instructions []int, position int) []int {
	opCode := instructions[position] % 100
	instruction := instructions[position] / 100

	var size int
	switch opCode {
	case opAdd, opMultiply, opEquals, opLessThan:
		size = 3
	case opJumpIfTrue, opJumpIfFalse:
		size = 2
	default:
		size = 1
	}

	p := make([]int, size)
	for i := 0; i < size; i++ {
		v := instructions[position+i+1]
		if instruction%10 == paramModePosition {
			v = instructions[v]
		}

		p[i] = v
		instruction = instruction / 10
	}
	return p
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
