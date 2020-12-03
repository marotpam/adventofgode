package _2018

import (
	"reflect"
	"strings"
	"testing"
)

func TestStepDurationCanBeCalculated(t *testing.T) {
	a := newStep("A")
	if a.duration != 61 {
		t.Errorf("A should have a duration of 61, got %d instead", a.duration)
	}

	z := newStep("Z")
	if z.duration != 86 {
		t.Errorf("Z should have a duration of 86, got %d instead", z.duration)
	}
}

func TestSchedulerCanFindTheNextAvailableWorkerWithOneAvailable(t *testing.T) {
	s := newScheduler(1)
	if w := s.idleWorkerIds(); len(w) != 1 {
		t.Errorf("A free worker should be found, none returned")
	}
}

func TestSchedulerCantFindTheNextAvailableWorkerWhenAllAreBusy(t *testing.T) {

	steps := []requirement{
		{"C", "A"},
		{"A", "E"},
	}
	i := newInstructions(steps)

	s := newScheduler(1)
	o, d := s.workOn(i)

	expectedOrder := []string{"C", "A", "E"}
	if !reflect.DeepEqual(expectedOrder, o) {
		t.Errorf("Expecting order to be %+v, got %+v in %d\n", expectedOrder, o, d)
	}
}

func TestAWorkerWillWorkOnAStepForTheDurationOfIt(t *testing.T) {
	w := worker{}
	s := newStepWithDuration("A", 1)

	w.workOn(s)

	if !w.isBusy() {
		t.Errorf("worker should be busy after starting its work\n")
	}

	w.work()

	if w.isBusy() {
		t.Errorf("worker should not be busy after a second has elapsed\n")
	}
}

func TestAWorkerWillUnlockABlockedStepAfterWorkingOnItsBlocker(t *testing.T) {
	w := worker{}
	blocker := newStepWithDuration("A", 1)
	blocked := newStep("B")

	blocker.blocking(blocked)
	if !blocked.isBlocked() {
		t.Errorf("blocked step should still be blocked because no worker started working on it\n")
	}

	w.workOn(blocker)

	if !blocked.isBlocked() {
		t.Errorf("blocked step should still be blocked because the worker didnt finish working on it\n")
	}

	w.work()

	if blocked.isBlocked() {
		t.Errorf("blocked step should no longer be blocked because the worker has finished working on it\n")
	}
}

func TestExample(t *testing.T) {
	steps := []requirement{
		{"C", "A"},
		{"C", "F"},
		{"A", "B"},
		{"A", "D"},
		{"B", "E"},
		{"D", "E"},
		{"F", "E"},
	}

	i := newInstructions(steps)

	expectedOrder := "CABFDE"

	if o := strings.Join(i.order(), ""); expectedOrder != o {
		t.Errorf("Failed asserting expected order %+v matches %+v\n", expectedOrder, o)
	}
}

func TestGivenExample(t *testing.T) {
	steps := []requirement{
		{"C", "A"},
		{"C", "F"},
		{"A", "B"},
		{"A", "D"},
		{"B", "E"},
		{"D", "E"},
		{"F", "E"},
	}

	i := newInstructions(steps)

	s := newScheduler(2)
	o, d := s.workOn(i)

	expectedOrder := []string{"C", "A", "B", "F", "D", "E"}

	if !reflect.DeepEqual(expectedOrder, o) {
		t.Errorf("Failed asserting expected order %+v matches %+v\n", expectedOrder, o)
	}

	if d != 15 {
		t.Errorf("It should take 15 units, got %d instead\n", d)
	}
}
