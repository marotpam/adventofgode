package _2020

import (
	"strings"
)

const (
	add        = '+'
	multiply   = '*'
	max        = '>'
	parenOpen  = '('
	parenClose = ')'
	whiteSpace = ' '
)

type operator rune

func (o operator) apply(a, b int) int {
	switch o {
	case add:
		return a + b
	case multiply:
		return a * b
	case max:
		if a > b {
			return a
		}
		return b
	}
	return 0
}

func SumLines(rawInput string) int {
	c := 0
	for _, l := range strings.Split(rawInput, "\n") {
		c += evaluate(l)
	}
	return c
}

func evaluate(line string) int {
	ints := newIntsStack()
	ops := newOpsStack()
	res := 0
	for _, part := range line {
		switch part {
		case whiteSpace:
			continue
		case add, multiply:
			ops.push(operator(part))
		case parenOpen:
			ints.push(res)
			ops.push(operator(max))
			res = 0
		case parenClose:
			op := ops.pop()
			res = op.apply(res, ints.pop())
		default:
			n := int(part - '0')
			op := ops.pop()
			res = op.apply(res, n)
		}
	}
	return res
}

type intStack struct {
	ints []int
}

func (s *intStack) push(n int) {
	s.ints = append(s.ints, n)
}

func (s *intStack) pop() int {
	if len(s.ints) == 0 {
		return 0
	}
	var n int
	last := len(s.ints) - 1
	n, s.ints = s.ints[last], s.ints[:last]
	return n
}

func newIntsStack() intStack {
	return intStack{ints: []int{}}
}

type opsStack struct {
	ops []operator
}

func (s *opsStack) push(n operator) {
	s.ops = append(s.ops, n)
}

func (s *opsStack) pop() operator {
	if len(s.ops) == 0 {
		return operator(max)
	}
	var o operator
	last := len(s.ops) - 1
	o, s.ops = s.ops[last], s.ops[:last]
	return o
}

func newOpsStack() opsStack {
	return opsStack{ops: []operator{}}
}
