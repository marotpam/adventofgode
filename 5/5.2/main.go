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
	p := &Polymer{}

	for _, u := range units {
		p.AddUnit(u)
	}

	return p
}

func (p *Polymer) AddUnit(u unit) {
	if len(p.units) == 0 {
		p.units = []unit{u}
	} else if p.units[len(p.units)-1].reactsWith(u) {
		p.units = p.units[:len(p.units)-1]
	} else {
		p.units = append(p.units, u)
	}
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

	for _, u := range p.units {
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

	ch := make(chan int, 1)

	shortestLength := math.MaxInt64
	for _, t := range ts {
		go func(t string) {
			wo := NewPolymer(p.unitsWithoutType(t))
			ch <- len(wo.resultingPolymers())
		}(t)
	}

	for i := 0; i < len(ts); i++ {
		l := <-ch
		if l < shortestLength {
			shortestLength = l
		}
	}

	fmt.Println(shortestLength) // 5310
}

func NewPolymerWithoutReaction(units []unit) *Polymer {
	return &Polymer{units}
}
