package main

import (
	"bufio"
	"fmt"
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

func newUnit(l string) unit {
	return unit{l}
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
	for _, t := range ts {
		ts = append(ts, t)
	}

	return ts
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
	p := NewPolymer(us)

	fmt.Println(len(p.resultingPolymers())) // 11894
}