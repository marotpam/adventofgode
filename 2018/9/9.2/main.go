package _2018

import (
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	marbleCount   int
	currentMarble  *marble
	firstMarble  *marble
	playersCount  int
	currentPlayer int
	scores        map[int]int
}

type marble struct {
	previous, next *marble
	value int
}

func newGame() *game {
	firstMarble := &marble{nil, nil, 0}
	firstMarble.previous = firstMarble
	firstMarble.next = firstMarble

	return &game{0, firstMarble, firstMarble,  0, 0, map[int]int{}}
}

func newGameWithPlayers(nPlayers int) *game {
	scores := make(map[int]int, nPlayers)
	firstMarble := &marble{nil, nil, 0}
	firstMarble.previous = firstMarble
	firstMarble.next = firstMarble

	return &game{0, firstMarble, firstMarble, nPlayers, 0, scores}
}

func (g *game) addMarble() {
	g.marbleCount++
	value := g.marbleCount

	if value%23 != 0 {
		curr := g.currentMarble.next

		m := &marble{curr, curr.next, value}
		m.previous.next = m
		m.next.previous = m

		g.currentMarble = m
	} else {
		curr := g.currentMarble
		for i := 0; i < 7; i++ {
			curr = curr.previous
		}
		g.addScore(curr.value + g.marbleCount)

		curr.previous.next = curr.next
		curr.next.previous = curr.previous

		g.currentMarble = curr.next
	}

	if g.playersCount != 0 {
		g.currentPlayer = (g.currentPlayer+1)%g.playersCount
	}
}

func (g *game) addScore(points int) {
	g.scores[g.currentPlayer] += points
}

func (g *game) playUntil(lastMarble int) {
	for i := 0; i < lastMarble; i++ {
		g.addMarble()
	}
}

func (g *game) highScore() int {
	hs := 0
	for _, s := range g.scores {
		if s > hs {
			hs = s
		}
	}
	return hs
}

func (g *game) printMarbles() string {
	vs := []string{strconv.Itoa(g.firstMarble.value)}

	for n := g.firstMarble.next; n != g.firstMarble; n = n.next {
		vs = append(vs, strconv.Itoa(n.value))
	}

	return strings.Join(vs, ", ")
}

func main() {
	g := newGameWithPlayers(426)

	g.playUntil(72058)

	fmt.Println(g.highScore())
}