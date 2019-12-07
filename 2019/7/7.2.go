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

type readWriter struct {
	elements []int
}

func (r *readWriter) write(n int) {
	r.elements = append(r.elements, n)
}

func (r *readWriter) read() int {
	a, rest := r.elements[0], r.elements[1:]
	r.elements = rest

	return a
}

func calculateSecondOptimalThrusterSequenceWithInputs(instructions []int, sequence []int) int {
	readWriters :=make([]*readWriter, len(sequence))
	positions := make(map[int]int, len(sequence))
	ampInstructions := make(map[int][]int, len(sequence))
	for i, s := range sequence {
		readWriters[i] = &readWriter{elements: []int{s}}
		positions[i] = 0
		ampInstructions[i] = append([]int{}, instructions...)
	}
	readWriters[0].elements = append(readWriters[0].elements, 0)

	for i := 0; ;  {
		pos := CalculateSecondOpcode(ampInstructions[i], readWriters[i], readWriters[(i+1)%len(sequence)], positions[i])
		positions[i] = pos
		if pos == -1 {
			delete(positions, i)
			if len(positions) == 0 {
				return readWriters[0].elements[len(readWriters[0].elements)-1]
			}
		}
		i = (i+1)%len(sequence)
	}

	return 0
}
