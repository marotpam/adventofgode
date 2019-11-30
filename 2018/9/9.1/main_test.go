package main

import (
	"reflect"
	"testing"
)

func TestANewGameStartsWithOneMarble(t *testing.T) {
	g := newGame()

	if !reflect.DeepEqual("0", g.printMarbles()) {
		t.Errorf("game should have the first marble after start, got %+v", g.printMarbles())
	}
}

func TestAddingMarbles(t *testing.T) {
	type testCase struct {
		description string
		nMarbles int
		expectedMarbles string
	}

	tcs := []testCase{
		{
			"One marble",
			1,
			"0, 1",
		},{
			"Two marbles",
			2,
			"0, 2, 1",
		},{
			"Three marbles",
			3,
			"0, 2, 1, 3",
		},{
			"Twenty-two marbles",
			22,
			"0, 16, 8, 17, 4, 18, 9, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15",
		},{
			"Twenty-three marbles",
			23,
			"0, 16, 8, 17, 4, 18, 19, 2, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15",
		},{
			"Twenty-four marbles",
			24,
			"0, 16, 8, 17, 4, 18, 19, 2, 24, 20, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15",
		},{
			"Twenty-five marbles",
			25,
			"0, 16, 8, 17, 4, 18, 19, 2, 24, 20, 25, 10, 21, 5, 22, 11, 1, 12, 6, 13, 3, 14, 7, 15",
		},
	}

	for _, tc := range tcs {
		g := newGame()
		for i := 0; i < tc.nMarbles; i++ {
			g.addMarble()
		}

		if m := g.printMarbles(); m != tc.expectedMarbles {
			t.Errorf("%s failed, expecting %s but got %s\n", tc.description, tc.expectedMarbles, m)
		}
	}
}

func TestHighScores(t *testing.T) {
	type testCase struct {
		nPlayers, playUntil, expectedScore int
	}

	tcs := []testCase{
		{
			9, 25, 32,
		},{
			10, 1618, 8317,
		},{
			13, 7999, 146373,
		},{
			17, 1104, 2764,
		},{
			21, 6111, 54718,
		},{
			30, 5807, 37305,
		},{
			426, 72058, 424112,
		},{
			426, 72058, 424112,
		},{
			426, 7205800, 3487352628,
		},
	}

	for _, tc := range tcs {
		g := newGameWithPlayers(tc.nPlayers)
		g.playUntil(tc.playUntil)

		if hs := g.highScore(); hs != tc.expectedScore {
			t.Errorf("High score should be %d, found %d instead\n", tc.expectedScore, hs)
		}
	}
}