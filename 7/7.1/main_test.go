package main

import (
	"reflect"
	"testing"
)

func TestStepsCanBeBlockedByOthers(t *testing.T) {
	steps := []requirement{
		{"C", "A"},
		{"A", "E"},
	}

	i := newInstructions(steps)

	expected := newStep("C")

	s := i.getStepById("C")

	if s.id != expected.id {
		t.Errorf("Failed asserting expected %+v matches %+v\n", expected, s)
	}

	if !reflect.DeepEqual(s.blocks, expected.blocks) {
		t.Errorf("Failed comparing blocks expected %+v matches %+v\n", expected, s)
	}
}

func TestBlockedStepsAreKeptInOrder(t *testing.T) {
	steps := []requirement{
		{"C", "B"},
		{"C", "A"},
		{"A", "E"},
	}

	i := newInstructions(steps)

	s := i.getStepById("C")

	if !reflect.DeepEqual([]string{"A", "B"}, s.blockedIds()) {
		t.Errorf("Failed comparing blocks expected %+v matches %+v\n", []string{"A", "B"}, s.blockedIds())
	}
}

func TestStepsCanBeUnlocked(t *testing.T) {
	steps := []requirement{
		{"C", "B"},
		{"C", "A"},
		{"A", "E"},
		{"D", "A"},
	}

	i := newInstructions(steps)

	s := i.getStepById("C")
	s.done()

	a := i.getStepById("A")

	if !reflect.DeepEqual([]string{}, s.blockedIds()) {
		t.Errorf("Failed comparing blocked ids expected %+v matches %+v\n", []string{}, s.blockedIds())
	}

	if !reflect.DeepEqual([]string{"D"}, a.blockedByIds()) {
		t.Errorf("Failed comparing bloced by ids expected %+v matches %+v\n", []string{"D"}, a.blockedByIds())
	}
}

func TestItCanReturnTheStepsInOrder(t *testing.T) {
	steps := []requirement{
		{"C", "A"},
		{"A", "E"},
	}

	i := newInstructions(steps)

	expectedOrder := []string{"C", "A", "E"}

	if o := i.order(); !reflect.DeepEqual(expectedOrder, o) {
		t.Errorf("Failed asserting expected order %+v matches %+v\n", expectedOrder, o)
	}
}
