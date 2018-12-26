package main

import (
	"reflect"
	"testing"
)

func TestAddrAddsTwoRegistersIntoAThirdOne(t *testing.T) {
	input := []int{1, 2, 3, 4}

	output := addr(0, 1, 0, input)

	expectedOutput := []int{3, 2, 3, 4}
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("registers should be %+v, got %+v", expectedOutput, output)
	}
}

func TestAddiAddsARegistersAndAValueIntoARegister(t *testing.T) {
	input := []int{1, 2, 3, 4}

	output := addi(0, 1, 0, input)

	expectedOutput := []int{2, 2, 3, 4}
	if !reflect.DeepEqual(expectedOutput, output) {
		t.Errorf("registers should be %+v, got %+v", expectedOutput, output)
	}
}

func TestCountingPossibleOpcodes(t *testing.T) {
	s := sample{
		input:  []int{3, 2, 1, 1},
		opcode: 9,
		a:      2,
		b:      1,
		c:      2,
		output: []int{3, 2, 2, 1},
	}

	if c := countPossibleOpcodes(s); c != 3 {
		t.Errorf("3 operations could apply to the sample, got %d\n", c)
	}
}