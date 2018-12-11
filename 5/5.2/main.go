package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type unit struct {
	letter string
}

func (u unit) reactsWith(other unit) bool {
	return strings.ToLower(u.letter) == strings.ToLower(other.letter) &&
		u.letter != other.letter
}

func (u unit) hasType(t string) bool {
	return strings.ToLower(u.letter) == strings.ToLower(t)
}

type Polymer struct {
	units []unit
}

func NewPolymer(units []unit) *Polymer {
	return &Polymer{initUnits(units)}
}

func initUnits(units []unit) []unit {
	if len(units) <= 1 {
		return units
	}

	for i := 0; i < len(units) - 1; i++ {
		if units[i].reactsWith(units[i+1]) {
			return initUnits(append(units[:i], units[i+2:]...))
		}
	}

	return units
}

func (p Polymer) resultingPolymers() string {
	s := ""

	for _, u := range p.units {
		s += u.letter
	}

	return s
}

func (p Polymer) getTypes() []string {
	m := map[string]bool{}

	for _, u := range p.units {
		m[strings.ToLower(u.letter)] = true
	}

	ts := []string{}
	for t := range m {
		ts = append(ts, t)
	}

	return ts
}

func (p Polymer) unitsWithoutType(t string) []unit {
	r := []unit{}

	for _, u := range p.units{
		if !u.hasType(t) {
			r = append(r, u)
		}
	}

	return r
}

func main() {
	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	fileScanner.Scan()

	us := []unit{}
	for _, c := range fileScanner.Text() {
		us = append(us, unit{string(c)})
	}
	p := NewPolymerWithoutReaction(us)
	ts := p.getTypes()

	shortestLength := math.MaxInt64
	for _, t := range ts {
		wo := NewPolymer(p.unitsWithoutType(t))

		if l := len(wo.resultingPolymers()); l < shortestLength {
			shortestLength = l
		}
	}

	fmt.Println(shortestLength) // 5310
}

func NewPolymerWithoutReaction(units []unit) *Polymer {
	return &Polymer{units}
}