package _2017

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	actionDecrease = "dec"
	actionIncrease = "inc"

	opGreaterThan        = ">"
	opGreaterOrEqualThan = ">="
	opLowerThan          = "<"
	opLowerOrEqualThan   = "<="
	opEqualTo            = "=="
	opNotEqualTo         = "!="
)

type computer struct {
	registers         map[string]int
	maxAllocatedValue int
}

func newComputer() *computer {
	return &computer{
		registers:         make(map[string]int),
		maxAllocatedValue: 0,
	}
}

func (c *computer) process(instructions []instruction) {
	for _, i := range instructions {
		if i.condition.isMetBy(*c) {
			switch i.action {
			case actionIncrease:
				c.registers[i.targetRegister] += i.value
			case actionDecrease:
				c.registers[i.targetRegister] -= i.value
			}

			if c.registers[i.targetRegister] > c.maxAllocatedValue {
				c.maxAllocatedValue = c.registers[i.targetRegister]
			}
		}
	}
}

func (c *computer) getMaxValue() int {
	max := 0

	for _, v := range c.registers {
		if v > max {
			max = v
		}
	}

	return max
}

type condition struct {
	register string
	op       string
	value    int
}

func (c condition) isMetBy(computer computer) bool {
	v := computer.registers[c.register]
	switch c.op {
	case opGreaterThan:
		return v > c.value
	case opGreaterOrEqualThan:
		return v >= c.value
	case opLowerThan:
		return v < c.value
	case opLowerOrEqualThan:
		return v <= c.value
	case opEqualTo:
		return v == c.value
	case opNotEqualTo:
		return v != c.value
	}
	return false
}

type instruction struct {
	targetRegister string
	action         string
	value          int
	condition      condition
}

func GetLargestValue(rawInput string) int {
	instructions := parse(rawInput)

	computer := newComputer()
	computer.process(instructions)

	return computer.getMaxValue()
}

func parse(rawInput string) []instruction {
	instructions := []instruction{}

	scanner := bufio.NewScanner(strings.NewReader(rawInput))
	for scanner.Scan() {

		var targetRegister, action, conditionRegister, op string
		var actionValue, conditionValue int
		fmt.Sscanf(scanner.Text(), "%s %s %d if %s %s %d",
			&targetRegister, &action, &actionValue, &conditionRegister, &op, &conditionValue,
		)

		instructions = append(instructions, instruction{
			targetRegister: targetRegister,
			action:         action,
			value:          actionValue,
			condition: condition{
				register: conditionRegister,
				op:       op,
				value:    conditionValue,
			},
		})
	}

	return instructions
}
