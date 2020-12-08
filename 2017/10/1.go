package _2017

import (
	"strconv"
	"strings"
)

const listLength = 256

type list struct {
	numbers         []int
	currentPosition int
	skipSize        int
}

func newList() *list {
	numbers := make([]int, 0, listLength)
	for i := 0; i < listLength; i++ {
		numbers = append(numbers, i)
	}

	return &list{
		numbers:         numbers,
		currentPosition: 0,
		skipSize:        0,
	}
}

func (l *list) process(length int) {
	for i := 0; i < length/2; i++ {
		front, back := (l.currentPosition+i)%len(l.numbers), (l.currentPosition+length-i-1)%len(l.numbers)
		l.numbers[front], l.numbers[back] = l.numbers[back], l.numbers[front]
	}

	l.currentPosition += (l.skipSize + length) % len(l.numbers)
	l.skipSize++
}

func MultiplyFirstTwoInList(rawInput string) int {
	list := newList()

	lengths := parse(rawInput)
	for _, l := range lengths {
		list.process(l)
	}

	return list.numbers[0] * list.numbers[1]
}

func parse(input string) []int {
	numbers := strings.Split(input, ",")
	ints := make([]int, 0, len(numbers))

	for _, rawNumber := range numbers {
		n, _ := strconv.Atoi(rawNumber)
		ints = append(ints, n)
	}
	return ints
}
