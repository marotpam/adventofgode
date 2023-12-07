package _2023

import "testing"

func TestCountNumberOfWaysForRecord(t *testing.T) {
	type testCase struct {
		race     race
		expected int
	}

	tcs := []testCase{
		{ // part 1, example race 1
			race:     race{time: 7, distance: 9},
			expected: 4,
		},
		{ // part 1, example race 2
			race:     race{time: 15, distance: 40},
			expected: 8,
		},
		{ // part 1, example race 3
			race:     race{time: 30, distance: 200},
			expected: 9,
		},
		{ // part 2, example
			race:     race{time: 71530, distance: 940200},
			expected: 71503,
		},
		{ // part 2, input
			race:     race{time: 40929790, distance: 215106415051100},
			expected: 28545089,
		},
	}

	for _, tc := range tcs {
		if got := tc.race.countNumberOfWaysForRecord(); got != tc.expected {
			t.Errorf("got %d want %d", got, tc.expected)
		}
	}
}

func TestMultiplyNumberOfWaysToBeatAllRecords(t *testing.T) {
	type testCase struct {
		rawInput string
		expected int
	}

	tcs := []testCase{
		{ // example
			rawInput: `Time:      7  15   30
			Distance:  9  40  200`,
			expected: 288,
		},
		{ // part 1
			rawInput: `Time:        40     92     97     90
			Distance:   215   1064   1505   1100`,
			expected: 6209190,
		},
	}

	for _, tc := range tcs {
		if got := MultiplyNumberOfWaysToBeatAllRecords(tc.rawInput); got != tc.expected {
			t.Errorf("got %d want %d", got, tc.expected)
		}
	}
}
