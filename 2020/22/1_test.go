package _2020

import "testing"

func Test_deck_score(t *testing.T) {
	tests := []struct {
		name string
		d    deck
		want int
	}{
		{
			name: "given example",
			d:    []int{3, 2, 10, 6, 8, 5, 9, 4, 7, 1},
			want: 306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.score(); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeckOperations(t *testing.T) {
	t.Run("adding cards to an empty deck", func(t *testing.T) {
		d := deck([]int{})

		d = d.add(1, 2)

		if d.isEmpty() {
			t.Error("deck should not be empty after adding a couple of cards")
		}

		if d[0] != 1 {
			t.Errorf("Unexpected card when popping the deck: %d", d[0])
		}
	})
	t.Run("adding cards and popping them afterwards", func(t *testing.T) {
		d := deck([]int{})

		d = d.add(2)

		var card int
		card, d = d.pop()

		if card != 2 {
			t.Errorf("Unexpected card when popping the deck: %d", d[0])
		}

		if !d.isEmpty() {
			t.Error("deck should not be empty after adding a couple of cards")
		}
	})
}

func Test_Play(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "given example",
			input: `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`,
			want: 306,
		},
		{
			name: "given input",
			input: `Player 1:
47
19
22
31
24
6
10
5
1
48
46
27
8
45
16
28
33
41
42
36
50
39
30
11
17

Player 2:
4
18
21
37
34
15
35
38
20
23
9
25
32
13
26
2
12
44
14
49
3
40
7
43
29`,
			want: 33434,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Play(tt.input); got != tt.want {
				t.Errorf("score() = %v, want %v", got, tt.want)
			}
		})
	}
}
