package main

const (
	opJumpIfTrue  = 5
	opJumpIfFalse = 6
	opLessThan    = 7
	opEquals      = 8
	opAdd         = 1
	opMultiply    = 2
	opInput       = 3
	opOutput      = 4
	opHalt        = 99

	paramModePosition  = 0
	paramModeImmediate = 1
)

func CalculateFirstOptimalThrusterSequence(inputs []int) int {
	perms := permutations([]int{0, 1, 2, 3, 4})
	max := 0
	for _, p := range perms {
		inputsCopy := append([]int{}, inputs...)
		c := calculateFirstOptimalThrusterSequenceWithInputs(inputsCopy, p)
		if c > max {
			max = c
		}
	}
	return max
}

func calculateFirstOptimalThrusterSequenceWithInputs(inputs []int, sequence []int) int {
	inputSignal, outputSignal := 0, 0

	for _, s := range sequence {
		input := &fakeInput{inputs: []int{s, inputSignal}}
		output := &fakeOutput{}

		CalculateSecondOpcode(inputs, input, output)

		outputSignal = output.outputs[len(output.outputs)-1]
		inputSignal = outputSignal
	}

	return outputSignal
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
			return
		}

		for i := 0; i < n; i++ {
			helper(arr, n-1)
			if n%2 == 1 {
				tmp := arr[i]
				arr[i] = arr[n-1]
				arr[n-1] = tmp
			} else {
				tmp := arr[0]
				arr[0] = arr[n-1]
				arr[n-1] = tmp
			}
		}
	}
	helper(arr, len(arr))
	return res
}

type input interface {
	read() int
}

type output interface {
	write(int)
}

type fakeInput struct {
	inputs []int
}

func (i *fakeInput) read() int {
	a, rest := i.inputs[0], i.inputs[1:]
	i.inputs = rest

	return a
}

type fakeOutput struct {
	outputs []int
}

func (o *fakeOutput) write(n int) {
	o.outputs = append(o.outputs, n)
}

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
