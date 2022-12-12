package _2022

import (
	"sort"
	"strconv"
	"strings"
)

func CountMaxCalories(rawInput string, topElfCount int) int {
	caloriesPerElf := make([]int, 0, 0)
	calorieCount := 0

	for _, rawLine := range strings.Split(rawInput, "\n") {
		line := strings.TrimSpace(rawLine)
		c, err := strconv.Atoi(line)
		if err != nil {
			caloriesPerElf = append(caloriesPerElf, calorieCount)
			calorieCount = 0
			continue
		}

		calorieCount += c
	}

	sort.Ints(caloriesPerElf)
	elfCount := len(caloriesPerElf) - 1

	maxCalories := 0
	for i := 0; i < topElfCount && i < elfCount; i++ {
		maxCalories += caloriesPerElf[elfCount-i]
	}

	return maxCalories
}
