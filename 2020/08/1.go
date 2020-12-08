package _2020

import (
	"fmt"
	"strings"
)

const (
	opJump       = "jmp"
	opAccumulate = "acc"
	opNoOp       = "nop"
)

type console struct {
	accumulator int
	pointer     int
}

func (c *console) process(instructions []instruction) bool {
	c.reset()

	visitedInstructions := make(map[int]bool, len(instructions))

	for i := 0; i < len(instructions); {
		if visitedInstructions[i] {
			return false
		}
		visitedInstructions[i] = true

		inst := instructions[i]
		switch inst.operation {
		case opJump:
			i += inst.argument
		case opAccumulate:
			c.accumulator += inst.argument
			i++
		case opNoOp:
			i++
		}
	}

	return true
}

func GetAccumulator(rawInput string) int {
	instructions := parse(rawInput)

	c := console{}
	c.process(instructions)
	return c.accumulator
}

type instruction struct {
	operation string
	argument  int
}

func parse(input string) []instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, 0, len(lines))

	for _, l := range lines {
		var operation string
		var argument int
		_, err := fmt.Sscanf(l, "%s %d", &operation, &argument)
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, instruction{
			operation: operation,
			argument:  argument,
		})
	}

	return instructions
}
