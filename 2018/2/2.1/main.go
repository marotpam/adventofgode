package main

import (
	"bufio"
	"fmt"
	"os"
)

func getOccurrences(boxID string) (exactlyTwice int, exactlyThrice int) {
	occurrences := map[rune]int{}
	for _, r := range boxID {
		occurrences[r]++
	}

	for _, o := range occurrences {
		if o == 2 {
			exactlyTwice = 1
		} else if o == 3 {
			exactlyThrice = 1
		}
	}

	return exactlyTwice, exactlyThrice
}

func main() {
	totalTwice, totalThrice := 0, 0

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		tw, tr := getOccurrences(fileScanner.Text())
		totalTwice += tw
		totalThrice += tr
	}

	fmt.Println(totalTwice * totalThrice) // 248*29=7192
}
