package _2021

import (
	"math"
	"strings"
)

func SubtractCommonQuantities(rawInput string, stepCount int) int {
	polymer := parse(rawInput)

	for i := 0; i < stepCount; i++ {
		polymer.step()
	}

	return polymer.subtractCommonQuantities()
}

type polymer struct {
	template       string
	insertionRules map[string]rune
	counts         map[rune]int
}

func (p *polymer) step() {
	templateBuilder := strings.Builder{}
	for i := 0; i < len(p.template)-1; i++ {
		pair := p.template[i : i+2]
		templateBuilder.WriteRune(rune(pair[0]))
		element, ok := p.insertionRules[pair]
		if ok {
			templateBuilder.WriteRune(element)
		}
	}
	templateBuilder.WriteRune(rune(p.template[len(p.template)-1]))
	p.template = templateBuilder.String()
}

func (p *polymer) subtractCommonQuantities() int {
	for _, c := range p.template {
		_, ok := p.counts[c]
		if !ok {
			p.counts[c] = 0
		}

		p.counts[c]++
	}

	min, max := math.MaxInt64, 0
	for _, c := range p.counts {
		if c < min {
			min = c
		}

		if c > max {
			max = c
		}
	}

	return max - min
}

func parse(rawInput string) polymer {
	lines := strings.Split(rawInput, "\n")
	polymerTemplate := lines[0]

	insertionRules := make(map[string]rune, len(lines)-2)
	for _, l := range lines[2:] {
		parts := strings.Split(l, " -> ")
		insertionRules[parts[0]] = rune(parts[1][0])
	}

	return polymer{
		template:       polymerTemplate,
		insertionRules: insertionRules,
		counts:         map[rune]int{},
	}
}
