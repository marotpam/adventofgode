package main

const (
	opAdd      = 1
	opMultiply = 2
	opInput    = 3
	opOutput   = 4
	opHalt     = 99

	paramModePosition = 0
	paramModeImmediate    = 1
)

type input interface {
	read() int
}

type output interface {
	write(int)
}

func CalculateFirstOpcode(ints []int, in input, out output) []int {
	for i := 0; ints[i]%100 != opHalt; {
		opCode := ints[i] % 100
		parameters := getParameters(opCode, ints[i]/100)
		switch opCode {
		case opAdd, opMultiply:
			a := ints[i+1]
			if parameters[0] == paramModePosition {
				a = ints[a]
			}

			b := ints[i+2]
			if parameters[1] == paramModePosition {
				b = ints[b]
			}

			ints[ints[i+3]] = calculateArithmeticOp(opCode, a, b)
		case opInput:
			ints[ints[i+1]] = in.read()
		case opOutput:
			v := ints[i+1]
			if parameters[0] == paramModePosition {
				v = ints[v]
			}
			out.write(v)
		}
		i += len(parameters) + 1
	}
	return ints
}

func getParameters(opCode int, instruction int) []int {
	size := 1
	if opCode == opAdd || opCode == opMultiply {
		size = 3
	}
	p := make([]int, size)
	for i := 0; i < size; i++ {
		p[i] = instruction % 10
		instruction = instruction / 10
	}
	return p
}

func calculateArithmeticOp(op, a, b int) int {
	switch op {
	case opAdd:
		return a + b
	case opMultiply:
		return a * b
	}
	return 0
}
