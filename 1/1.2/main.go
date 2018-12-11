package main

import (
	"bufio"
	"fmt"
	"os"
)

func readNumbers() []int{
	n := []int{}
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		var number int
		fmt.Sscanf(fileScanner.Text(), "%d", &number)

		n = append(n, number)
	}
	return n
}

func findRepeatedFrequency() int {
	frequencies := map[int]bool{}
	total := 0
	frequencies[total] = true
	numbers := readNumbers()
	for ;; {
		for _, n := range numbers {
			total += n
			if frequencies[total] {
				return total
			} else {
				frequencies[total] = true
			}
		}
	}
}

func main() {
	fmt.Println(findRepeatedFrequency()) // 72330
}