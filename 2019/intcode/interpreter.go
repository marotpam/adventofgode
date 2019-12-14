package main

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

type Reader interface {
	read() int
}

type Writer interface {
	write(int)
}

type Interpreter struct {
	relativeBase int
	memory map[int]int
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		relativeBase: 0,
		memory:       nil,
	}
}

func (in *Interpreter) Run(instructions []int, reader Reader, writer Writer) {
	in.memory = map[int]int{}
	for i, v := range instructions {
		in.memory[i] = v
	}

	opCode := 0
	for i := 0; opCode != opHalt; {
		opCode = in.memory[i] % 100
		instruction := in.memory[i] / 100

		operandCount := 1
		switch opCode {
		case opAdd, opMultiply, opEquals, opLessThan:
			operandCount = 3
		case opJumpIfTrue, opJumpIfFalse:
			operandCount = 2
		}

		operands := make([]int, operandCount)
		offset := in.memory[i+operandCount]
		for o := 0; o < operandCount; o++ {
			p := in.memory[i+o+1]
			switch instruction % 10 {
			case paramModePosition:
				p = in.memory[p]
			case paramModeRelativeBase:
				p = in.memory[p+in.relativeBase]
				if o == operandCount-1 {
					offset += in.relativeBase
				}
			}

			operands[o] = p
			instruction = instruction / 10
		}
		newPosition := -1

		switch opCode {
		case opAdd, opMultiply, opEquals, opLessThan:
			in.memory[offset] = calculateArithmeticOp(opCode, operands[0], operands[1])
		case opJumpIfTrue:
			if operands[0] != 0 {
				newPosition = operands[1]
			}
		case opJumpIfFalse:
			if operands[0] == 0 {
				newPosition = operands[1]
			}
		case opInput:
			in.memory[offset] = reader.read()
		case opOutput:
			writer.write(operands[0])
		case opAdjustRelativeBase:
			in.relativeBase += operands[0]
		}
		i += len(operands) + 1

		if newPosition >= 0 {
			i = newPosition
		}
	}
}

func calculateArithmeticOp(op, a, b int) int {
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
