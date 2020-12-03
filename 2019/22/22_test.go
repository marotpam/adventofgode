package _2019

import (
	"reflect"
	"testing"
)

func TestDealingIntoNewStack(t *testing.T) {
	d := newDeck(10)

	d.dealIntoNewStack()

	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	assertCardsMatch(t, expected, d.cards)
}

func TestCuttingN(t *testing.T) {
	t.Run("take n off top", func(t *testing.T) {
		d := newDeck(10)

		d.cut(3)

		expected := []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
		assertCardsMatch(t, expected, d.cards)
	})
	t.Run("take n off bottom", func(t *testing.T) {
		d := newDeck(10)

		d.cut(-4)

		expected := []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
		assertCardsMatch(t, expected, d.cards)
	})
}

func TestDealWithIncrement(t *testing.T) {
	d := newDeck(10)

	d.dealWithIncrement(3)

	expected := []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3}
	assertCardsMatch(t, expected, d.cards)
}

func assertCardsMatch(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf(
			"expectation failed, expecting %v got %v",
			expected, actual,
		)
	}
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		name         string
		instructions []string
		card         int
		want         int
	}{
		{
			name: "input for first part",
			card: 2019,
			instructions: []string{
				"cut 1650",
				"deal with increment 24",
				"cut 7388",
				"deal with increment 27",
				"cut -4764",
				"deal with increment 31",
				"cut 7890",
				"deal into new stack",
				"cut -1564",
				"deal with increment 46",
				"cut 940",
				"deal with increment 29",
				"cut -2698",
				"deal with increment 49",
				"cut 7279",
				"deal with increment 72",
				"cut -8095",
				"deal into new stack",
				"deal with increment 33",
				"cut 6936",
				"deal with increment 58",
				"cut -5224",
				"deal with increment 18",
				"deal into new stack",
				"deal with increment 73",
				"deal into new stack",
				"deal with increment 46",
				"cut 3746",
				"deal with increment 56",
				"cut 2168",
				"deal with increment 55",
				"cut 4353",
				"deal into new stack",
				"cut -1545",
				"deal into new stack",
				"deal with increment 9",
				"cut 7797",
				"deal into new stack",
				"deal with increment 36",
				"cut 5488",
				"deal with increment 59",
				"cut 2417",
				"deal with increment 60",
				"cut 7692",
				"deal into new stack",
				"cut 3590",
				"deal with increment 51",
				"cut -4782",
				"deal with increment 62",
				"cut -8990",
				"deal with increment 44",
				"deal into new stack",
				"deal with increment 27",
				"cut 1810",
				"deal with increment 38",
				"cut -7252",
				"deal with increment 48",
				"deal into new stack",
				"deal with increment 13",
				"cut 1515",
				"deal with increment 25",
				"cut 4664",
				"deal into new stack",
				"deal with increment 37",
				"cut 2738",
				"deal into new stack",
				"deal with increment 48",
				"cut 1958",
				"deal into new stack",
				"deal with increment 33",
				"cut 7938",
				"deal with increment 19",
				"cut 6615",
				"deal with increment 3",
				"cut 5269",
				"deal with increment 45",
				"cut 980",
				"deal with increment 23",
				"cut -4607",
				"deal with increment 5",
				"deal into new stack",
				"deal with increment 27",
				"cut 6567",
				"deal into new stack",
				"cut 8801",
				"deal with increment 26",
				"cut 4446",
				"deal into new stack",
				"deal with increment 72",
				"cut 7990",
				"deal with increment 5",
				"cut -6196",
				"deal with increment 14",
				"deal into new stack",
				"deal with increment 42",
				"deal into new stack",
				"deal with increment 30",
				"deal into new stack",
				"deal with increment 65",
				"cut 7733",
			},
			want: 7614,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindCardPositionAfterShuffling(tt.instructions, tt.card); got != tt.want {
				t.Errorf("FindCardPositionAfterShuffling() = %v, want %v", got, tt.want)
			}
		})
	}
}
