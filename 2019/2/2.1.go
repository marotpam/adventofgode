package main

const (
	opAdd      = 1
	opMultiply = 2
	opHalt     = 99
)

func CalculateFirstOpcode(ints []int) []int {
	for i := 0; ints[i] != opHalt; {
		switch ints[i] {
		case opAdd:
			ints[ints[i+3]] = ints[ints[i+1]] + ints[ints[i+2]]
		case opMultiply:
			ints[ints[i+3]] = ints[ints[i+1]] * ints[ints[i+2]]
		}
		i += 4
	}
	return ints
}
