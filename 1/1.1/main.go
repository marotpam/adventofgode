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

func main() {
	total := 0

	for _, n := range readNumbers() {
		total += n
	}

	fmt.Println(total) // 505
}