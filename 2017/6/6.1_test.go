package _2017

import (
	"testing"
)

func TestCountJumpsSecond(t *testing.T) {
	type testCase struct {
		name string
		in   []int
		out  int
	}

	tcs := []testCase{
		{
			name: "given example",
			in:   []int{0, 2, 7, 0},
			out:  5,
		},
		{
			name: "given input",
			in:   []int{2, 8, 8, 5, 4, 2, 3, 1, 5, 5, 1, 2, 15, 13, 5, 14},
			out:  5,
		},
	}

	for _, tc := range tcs {
		if out := CountRedistributionCycles(tc.in); out != tc.out {
			t.Errorf("%s: expecting %d got %d for %+v", tc.name, tc.out, out, tc.in)
		}
	}
}
