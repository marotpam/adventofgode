package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readBoxIDs() []string {
	ids := []string{}
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		ids = append(ids, fileScanner.Text())
	}
	return ids
}

func onlyMatching(first, second string) string {
	difs := []string{}
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			difs = append(difs, string(second[i]))
		}
	}
	return strings.Join(difs, "")
}

func findAlmostIdentical() string {
	ids := readBoxIDs()

	for i := 0; i < len(ids); i++ {
		for j := i + 1; j < len(ids); j++ {
			m := onlyMatching(ids[i], ids[j])
			if len(ids[i])-len(m) == 1 {
				return m
			}
		}
	}

	return ""
}

func main() {
	fmt.Println(findAlmostIdentical()) // mbruvapghxlzycbhmfqjonsie
}
