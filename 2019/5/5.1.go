package main

const (
	opAdd      = 1
	opMultiply = 2
	opInput    = 3
	opOutput   = 4
	opHalt     = 99

	paramModePosition  = 0
	paramModeImmediate = 1
)

type input interface {
	read() int
}

type output interface {
	write(int)
}

func CalculateFirstOpcode(ints []int, in input, out output) []int {
	for i := 0; ints[i]%100 != opHalt; {
		args := getArguments(ints, i)
		switch ints[i] % 100 {
		case opAdd, opMultiply:
			ints[ints[i+3]] = calculateArithmeticOp(ints[i]%100, args[0], args[1])
		case opInput:
			ints[ints[i+1]] = in.read()
		case opOutput:
			out.write(args[0])
		}
		i += len(args) + 1
	}
	return ints
}

func getArguments(instructions []int, position int) []int {
	opCode := instructions[position] % 100
	instruction := instructions[position] / 100

	size := 1
	if opCode == opAdd || opCode == opMultiply {
		size = 3
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

func calculateArithmeticOp(op, a, b int) int {
	switch op {
	case opAdd:
		return a + b
	case opMultiply:
		return a * b
	}
	return 0
}
