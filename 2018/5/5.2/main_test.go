package _2018

import (
	"reflect"
	"testing"
)

func TestGetPolymerTypes(t *testing.T) {
	type testCase struct {
		input          string
		expectedOutput []string
		description    string
	}

	tcs := []testCase{
		{
			"bAB",
			[]string{"b", "a"},
			"with two types",
		},
	}

	for _, tc := range tcs {
		us := []unit{}
		for _, l := range tc.input {
			us = append(us, unit{string(l)})
		}

		p := NewPolymer(us)

		if res := p.getTypes(); !reflect.DeepEqual(res, tc.expectedOutput) {
			t.Errorf("%s: Expecting %s, got %s", tc.description, tc.expectedOutput, res)
		}
	}
}

func TestPolymerUnitsWithoutType(t *testing.T) {
	p := NewPolymerWithoutReaction([]unit{{"a"}, {"A"}, {"B"}})

	if res := p.unitsWithoutType("a"); !reflect.DeepEqual([]unit{{"B"}}, res) {
		t.Errorf("Should only contain B, got %s", res)
	}
}
