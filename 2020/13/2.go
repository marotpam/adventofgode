package _2020

import (
	"strconv"
	"strings"
)

func SolvePart2(rawInput string) int {
	res, step := 0, 1
	for i, s := range strings.Split(rawInput, ",") {
		id, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		for (res+i)%id != 0 {
			res += step
		}
		step *= id
	}
	return res
}
