package main

func CalculateSecondOptimalThrusterSequence(inputs []int) int {
	perms := permutations([]int{5, 6, 7, 8, 9})
	max := 0
	for _, p := range perms {
		inputsCopy := append([]int{}, inputs...)
		c := calculateSecondOptimalThrusterSequenceWithInputs(inputsCopy, p)
		if c > max {
			max = c
		}
	}
	return max
}

type io struct {
	elements []int
}

func (i *io) write(n int) {
	i.elements = append(i.elements, n)
}

func (i *io) read() int {
	a, rest := i.elements[0], i.elements[1:]
	i.elements = rest

	return a
}

type amplifier struct {
	*io
	currentPosition int
	instructions    []int
}

func newAmplifier(phaseSetting int, instructions []int) *amplifier {
	return &amplifier{
		instructions:    instructions,
		io:              &io{elements: []int{phaseSetting}},
		currentPosition: 0,
	}
}

func (a *amplifier) isDone() bool {
	return a.currentPosition == -1
}

func (a *amplifier) input(i int) {
	a.io.elements = append(a.io.elements, i)
}

func (a *amplifier) output() int {
	return a.io.elements[len(a.io.elements)-1]
}

func calculateSecondOptimalThrusterSequenceWithInputs(instructions []int, sequence []int) int {
	amplifiers := make([]*amplifier, len(sequence))
	for i, s := range sequence {
		amplifiers[i] = newAmplifier(s, append([]int{}, instructions...))
	}
	amplifiers[0].input(0)

	for i := 0; ; i = (i + 1) % len(sequence) {
		amp := amplifiers[i]
		amp.currentPosition = CalculateSecondOpcode(
			amp.instructions,
			amp,
			amplifiers[(i+1)%len(sequence)],
			amp.currentPosition,
		)

		if allDone(amplifiers) {
			return amplifiers[0].output()
		}
	}

	return 0
}

func allDone(as []*amplifier) bool {
	for _, a := range as {
		if !a.isDone() {
			return false
		}
	}
	return true
}
