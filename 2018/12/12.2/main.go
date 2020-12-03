package _2018

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type pot struct {
	index          int
	hasPlant       bool
	willHavePlant  bool
	previous, next *pot
}

func (p *pot) willGrow() bool {
	return !p.hasPlant && p.willHavePlant
}

func (p *pot) willDie() bool {
	return p.hasPlant && !p.willHavePlant
}

func (p *pot) getRule() string {
	buf := bytes.Buffer{}

	for po := p.previous.previous; po != p.next.next.next; po = po.next {
		r := '.'
		if po.hasPlant {
			r = '#'
		}
		buf.WriteRune(r)
	}

	return buf.String()
}

type garden struct {
	first, last *pot
	rules       map[string]string
}

func newGarden(initialState string, rules map[string]string) *garden {
	minusThree := &pot{
		index: -3,
	}
	minusTwo := &pot{
		index:    -2,
		previous: minusThree,
	}
	minusThree.next = minusTwo

	minusOne := &pot{
		index:    -1,
		previous: minusTwo,
	}
	minusTwo.next = minusOne

	var previous *pot

	for i := range initialState {
		pot := &pot{
			index:    i,
			hasPlant: initialState[i] == '#',
		}
		if previous == nil {
			pot.previous = minusOne
			minusOne.next = pot
		} else {
			pot.previous = previous
			previous.next = pot
		}
		previous = pot
	}

	plusOne := &pot{
		index:    len(initialState),
		previous: previous,
	}
	previous.next = plusOne

	plusTwo := &pot{
		index:    plusOne.index + 1,
		previous: plusOne,
	}
	plusOne.next = plusTwo

	plusThree := &pot{
		index:    plusTwo.index + 1,
		previous: plusTwo,
	}
	plusTwo.next = plusThree

	return &garden{minusOne.next, plusOne, rules}
}

func (g *garden) growGeneration() {
	first := g.first

	previous := g.first.previous
	g.grow(previous)
	if previous.willGrow() {
		newHead := &pot{
			index: previous.previous.index-1,
			next:  previous.previous,
		}
		previous.previous.previous = newHead
		g.first = previous
	} else if previous.willDie() {
		newNext := g.first.next
		newNext.previous = nil
		g.first = newNext
	}

	var p *pot
	for p = first; p != g.last; p = p.next {
		g.grow(p)
	}

	afterLast := p
	g.grow(afterLast)

	if afterLast.willGrow() {
		g.last = afterLast.next
		newTail := &pot{
			index:    afterLast.next.next.index+1,
			previous: afterLast.next.next,
		}
		afterLast.next.next.next = newTail
	} else if afterLast.willDie() {
		g.last = afterLast.previous
	}

	for p := g.first; p != g.last; p = p.next {
		p.hasPlant = p.willHavePlant
	}
}

func (g *garden) grow(p *pot) {
	buf := bytes.Buffer{}

	for po := p.previous.previous; po != p.next.next.next; po = po.next {
		if po == nil {
			return
		}
		r := '.'
		if po.hasPlant {
			r = '#'
		}
		buf.WriteRune(r)
	}

	applicableRule := buf.String()

	if r, ok := g.rules[applicableRule]; ok {
		if r == "#" {
			p.willHavePlant = true
		} else {
			p.willHavePlant = false
		}
	} else {
		p.willHavePlant = false
	}
}

func (g *garden) potsAsString() string {
	var buf bytes.Buffer
	i := 0
	for p := g.first; p != g.last; p = p.next {
		r := '.'
		if p.hasPlant {
			r = '#'
		}
		buf.WriteRune(r)
		i++
	}

	return buf.String()
}

func (g *garden) growFor(nGenerations int) {
	prev := 0
	for i := 0; i < nGenerations; i++ {
		g.growGeneration()
		sum := g.sumPotIndexesWithPlants()
		fmt.Println(i+1, g.first.index, g.last.index, sum, sum - prev)
		prev = sum
	}
}

func (g *garden) sumPotIndexesWithPlants() int {
	i := 0

	for p := g.first; p != g.last; p = p.next {
		if p.hasPlant {
			i += p.index
		}
	}

	return i
}

func main() {
	fh, _ := os.Open("input.txt")
	defer fh.Close()

	fs := bufio.NewScanner(fh)
	fs.Scan()
	var initialState string
	fmt.Sscanf(fs.Text(), `initial state: %s`, &initialState)

	fs.Scan()

	rules := map[string]string{}
	for fs.Scan() {
		var rule, result string
		fmt.Sscanf(fs.Text(), `%s => %s`, &rule, &result)
		rules[rule] = result
	}

	g := newGarden(initialState, rules)

	g.growFor(50000000000)

	// after round 103 (where the counter is 9375), each round
	// the counter increases by 69 units consistently
	// with this, we can calculate the amount after 50B with
	// the formula 9375+(50000000000-103)*69
	fmt.Println(g.sumPotIndexesWithPlants()) // 3450000002268
}
