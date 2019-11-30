package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestANewGuardHasntSlept(t *testing.T) {
	g := newGuard(1)
	if m := g.countMinutesSleeping(); m != 0 {
		t.Errorf("New guards should have slept 0 minutes, got %d", m)
	}
}

func TestCountingMinutesSleepingWhenFallingAsleepOnce(t *testing.T) {
	g := newGuard(1)

	g.sleptBetween(5, 10)

	if m := g.countMinutesSleeping(); m != 5 {
		t.Errorf("Guards should have slept 5 minutes, got %d", m)
	}
}

func TestCountingMinutesSleepingWhenFallingAsleepTwice(t *testing.T) {
	g := newGuard(1)

	g.sleptBetween(5, 10)
	g.sleptBetween(25, 35)

	if m := g.countMinutesSleeping(); m != 15 {
		t.Errorf("Guard should have slept 15 minutes, got %d", m)
	}
}

func TestGetMostLikelyHourToBeAsleep(t *testing.T) {
	g := newGuard(1)

	g.sleptBetween(5, 10)
	g.sleptBetween(9, 11)

	if m := g.mostLikelyMinuteToBeAsleep(); m != 9 {
		t.Errorf("Guard should be more like to be asleep in minute 9, got %d", m)
	}
}

func TestEntriesCanBeComparedByTheTimeTheyHappened(t *testing.T) {
	firstEntry := entry{2018, 12, 06, 05, 00, "action", 0}
	secondEntry := entry{2018, 12, 06, 05, 01, "action", 0}

	if !firstEntry.happenedBefore(secondEntry) {
		t.Error("First entry should have happened before")
	}
}

func TestASliceOfUnorderedEntriesCanBeSortedByOccurrence(t *testing.T) {
	unorderedEntries := []entry{
		entry{2018, 12, 06, 05, 01, "action", 0},
		entry{2018, 12, 01, 01, 00, "action", 0},
		entry{2018, 11, 06, 05, 00, "action", 0},
	}

	orderedEntries := []entry{
		entry{2018, 11, 06, 05, 00, "action", 0},
		entry{2018, 12, 01, 01, 00, "action", 0},
		entry{2018, 12, 06, 05, 01, "action", 0},
	}

	sort.Slice(unorderedEntries, func(i, j int) bool { return unorderedEntries[i].happenedBefore(unorderedEntries[j]) })

	if !reflect.DeepEqual(unorderedEntries, orderedEntries) {
		t.Errorf("Failed asserting that \n%+v is \n%+v\n", unorderedEntries, orderedEntries)
	}
}

func TestASliceOfUnorderedEntriesCanBeSorted(t *testing.T) {
	unorderedEntries := []entry{
		{year: 1518, month: 4, day: 15, hour: 0, minute: 17, action: ""},
		{year: 1518, month: 5, day: 8, hour: 0, minute: 52, action: ""},
		{year: 1518, month: 10, day: 3, hour: 0, minute: 32, action: ""},
		{year: 1518, month: 7, day: 5, hour: 0, minute: 51, action: ""},
		{year: 1518, month: 6, day: 9, hour: 0, minute: 16, action: ""},
		{year: 1518, month: 4, day: 28, hour: 0, minute: 43, action: ""},
		{year: 1518, month: 7, day: 9, hour: 0, minute: 38, action: ""},
	}

	orderedEntries := []entry{
		{year: 1518, month: 4, day: 15, hour: 0, minute: 17, action: ""},
		{year: 1518, month: 4, day: 28, hour: 0, minute: 43, action: ""},
		{year: 1518, month: 5, day: 8, hour: 0, minute: 52, action: ""},
		{year: 1518, month: 6, day: 9, hour: 0, minute: 16, action: ""},
		{year: 1518, month: 7, day: 5, hour: 0, minute: 51, action: ""},
		{year: 1518, month: 7, day: 9, hour: 0, minute: 38, action: ""},
		{year: 1518, month: 10, day: 3, hour: 0, minute: 32, action: ""},
	}

	sort.Slice(unorderedEntries, func(i, j int) bool {
		return unorderedEntries[i].happenedBefore(unorderedEntries[j])
	})

	if !reflect.DeepEqual(unorderedEntries, orderedEntries) {
		t.Errorf("Failed asserting that \n%+v is \n%+v", unorderedEntries, orderedEntries)
	}
}

func TestASchedulerCanCalculateResultWithOnlyOneGuardAndOneDay(t *testing.T) {
	s := newScheduler([]entry{
		{2018, 1, 1, 23, 58, "begin_shift", 1},
		{2018, 1, 2, 00, 10, "falls_asleep", 0},
		{2018, 1, 2, 00, 27, "wakes_up", 0},
	})

	if c := s.calculate(); c != 10 {
		t.Errorf("Should return 10, got %d instead", c)
	}
}
