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
		parameters := getSecondParameters(opCode, ints[i]/100)
		newPosition := -1

		switch opCode {
		case opAdd, opMultiply, opEquals, opLessThan:
			a := ints[i+1]
			if parameters[0] == paramModePosition {
				a = ints[a]
			}

			b := ints[i+2]
			if parameters[1] == paramModePosition {
				b = ints[b]
			}

			ints[ints[i+3]] = calculateSecondArithmeticOp(opCode, a, b)
		case opInput:
			ints[ints[i+1]] = in.read()
		case opOutput:
			v := ints[i+1]
			if parameters[0] == paramModePosition {
				v = ints[v]
			}
			out.write(v)
		case opJumpIfTrue:
			a := ints[i+1]
			if parameters[0] == paramModePosition {
				a = ints[a]
			}

			b := ints[i+2]
			if parameters[1] == paramModePosition {
				b = ints[b]
			}

			if a != 0 {
				newPosition = b
			}
		case opJumpIfFalse:
			a := ints[i+1]
			if parameters[0] == paramModePosition {
				a = ints[a]
			}

			b := ints[i+2]
			if parameters[1] == paramModePosition {
				b = ints[b]
			}

			if a == 0 {
				newPosition = b
			}
		}
		i += len(parameters) + 1

		if newPosition >= 0 {
			i = newPosition
		}
	}
	return ints
}

func getSecondParameters(opCode int, instruction int) []int {
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
		p[i] = instruction % 10
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
