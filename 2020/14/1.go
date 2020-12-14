package _2020

import (
	"fmt"
	"strconv"
	"strings"
)

type mask string

func (m mask) apply(value int) int {
	binary := strconv.FormatInt(int64(value), 2)
	res := []rune(fmt.Sprintf("%036s", binary))
	for i := 0; i < len(res); i++ {
		if m[len(m)-1-i] != 'X' {
			res[len(res)-1-i] = rune(m[len(m)-1-i])
		}
	}
	r, _ := strconv.ParseInt(string(res), 2, 64)
	return int(r)
}

type op struct {
	register, value int
}

func SumValuesInMemory1(rawInput string) int {
	registers := make(map[int]int)
	var m mask

	for _, l := range strings.Split(rawInput, "\n") {
		parts := strings.Split(l, " = ")
		if parts[0] == "mask" {
			m = mask(parts[1])
			continue
		}

		o := parseOp(parts)

		registers[o.register] = m.apply(o.value)
	}

	c := 0
	for _, v := range registers {
		c += v
	}

	return c
}

func parseOp(parts []string) op {
	start := strings.Index(parts[0], "[") + 1
	end := strings.Index(parts[0], "]")
	if start < 0 || end < 0 {
		panic(parts)
	}
	r, _ := strconv.Atoi(parts[0][start:end])
	v, _ := strconv.Atoi(parts[1])

	return op{
		register: r,
		value:    v,
	}
}
