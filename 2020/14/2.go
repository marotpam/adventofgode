package _2020

import (
	"fmt"
	"strconv"
	"strings"
)

func (m mask) decode(register int) []int {
	binary := strconv.FormatInt(int64(register), 2)
	res := []rune(fmt.Sprintf("%036s", binary))
	for i := 0; i < len(res); i++ {
		if m[len(m)-1-i] != '0' {
			res[len(res)-1-i] = rune(m[len(m)-1-i])
		}
	}
	return decodeAll(string(res))
}

func decodeAll(s string) []int {
	r, err := strconv.ParseInt(s, 2, 64)
	if err == nil {
		return []int{int(r)}
	}

	return append(
		decodeAll(strings.Replace(s, "X", "0", 1)),
		decodeAll(strings.Replace(s, "X", "1", 1))...,
	)
}

func SumValuesInMemory2(rawInput string) int {
	registers := make(map[int]int)
	var m mask

	for _, l := range strings.Split(rawInput, "\n") {
		parts := strings.Split(l, " = ")
		if parts[0] == "mask" {
			m = mask(parts[1])
			continue
		}

		o := parseOp(parts)

		affectedRegisters := m.decode(o.register)
		for _, r := range affectedRegisters {
			registers[r] = o.value
		}
	}

	c := 0
	for _, v := range registers {
		c += v
	}

	return c
}
