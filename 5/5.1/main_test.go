package main

import "testing"

func TestUnitsWithDifferentTypesDontReact(t *testing.T) {
	u := newUnit("a")
	r := newUnit("b")

	if u.reactsWith(r) {
		t.Errorf("a and b have different types, so they should not react")
	}
}

func TestUnitsWithSameTypeAndSamePolarityDontReact(t *testing.T) {
	u := newUnit("a")
	r := newUnit("a")

	if u.reactsWith(r) {
		t.Errorf("a and a have same types and polarity, so they should not react")
	}
}

func TestUnitsWithSameTypeAndDifferentPolarityDontReact(t *testing.T) {
	u := newUnit("a")
	r := newUnit("A")

	if !u.reactsWith(r) {
		t.Errorf("a and a have same type and different polarity, so they should react")
	}
}

func TestResultingPolymer(t *testing.T) {
	type testCase struct {
		input          string
		expectedOutput string
		description    string
	}

	tcs := []testCase{
		{
			"dabAcCaCBAcCcaDA",
			"dabCBAcaDA",
			"complicated",
		}, {
			"bAaB",
			"",
			"reacting ends",
		}, {
			"bAa",
			"b",
			"reacting suffix",
		}, {
			"Aa",
			"",
			"only reacting units",
		}, {
			"Ab",
			"Ab",
			"non reacting units",
		},
	}

	for _, tc := range tcs {
		us := []unit{}
		for _, l := range tc.input {
			us = append(us, unit{string(l)})
		}

		p := NewPolymer(us)

		if res := p.resultingPolymers(); res != tc.expectedOutput {
			t.Errorf("%s: Expecting %s, got %s", tc.description, tc.expectedOutput, res)
		}
	}

}
