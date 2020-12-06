package _2017

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

type line struct {
	name       string
	weight     int
	discsAbove []string
}

type node struct {
	name     string
	weight   int
	parent   *node
	children []*node
}

func FindBottom(input string) string {
	lines := parseLines(input)
	nodeNames := make(map[string]bool)

	for _, l := range lines {
		nodeNames[l.name] = true
	}

	for _, l := range lines {
		for _, d := range l.discsAbove {
			delete(nodeNames, d)
		}
	}

	for k := range nodeNames {
		return k
	}

	return ""
}

func parseLines(input string) []line {
	lines := []line{}
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var name, discs string
		var weight int
		text := scanner.Text()
		_, err := fmt.Sscanf(text, `%s (%d)`, &name, &weight)
		if err != nil {
			log.Fatal("cannot read", err)
		}

		parts := strings.Split(text, " -> ")
		if len(parts) > 1 {
			discs = parts[1]
		}

		l := line{
			name:       name,
			weight:     weight,
			discsAbove: strings.Split(discs, ", "),
		}

		lines = append(lines, l)
	}

	return lines
}
