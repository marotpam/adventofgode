package _2018

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

func entryFromLine(line string) entry {
	var year, month, day, hour, minute, guardID int

	_, err := fmt.Sscanf(line, `[%d-%d-%d %d:%d] wakes up`, &year, &month, &day, &hour, &minute)
	if err == nil {
		return entry{year, month, day, hour, minute, "wakes_up", 0}
	}

	_, err = fmt.Sscanf(line, `[%d-%d-%d %d:%d] falls asleep`, &year, &month, &day, &hour, &minute)
	if err == nil {
		return entry{year, month, day, hour, minute, "falls_asleep", 0}
	}

	fmt.Sscanf(line, `[%d-%d-%d %d:%d] Guard #%d begins shift`, &year, &month, &day, &hour, &minute, &guardID)
	return entry{year, month, day, hour, minute, "begin_shift", guardID}
}

func readEntries() []entry {
	es := []entry{}

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		es = append(es, entryFromLine(fileScanner.Text()))
	}

	return es
}

type guard struct {
	id              int
	sleepingMinutes []int
}

func (g guard) countMinutesSleeping() int {
	c := 0

	for _, m := range g.sleepingMinutes {
		c += m
	}

	return c
}

func (g guard) sleptBetween(start, end int) {
	for i := start; i < end; i++ {
		g.sleepingMinutes[i]++
	}
}

func (g guard) mostLikelyMinuteToBeAsleep() int {
	maxMinutes, minute := 0, 0
	for i := 0; i < len(g.sleepingMinutes); i++ {
		if g.sleepingMinutes[i] > maxMinutes {
			maxMinutes = g.sleepingMinutes[i]
			minute = i
		}
	}
	return minute
}

func newGuard(id int) guard {
	return guard{id, make([]int, 60)}
}

type entry struct {
	year, month, day, hour, minute int
	action                         string
	guardID                        int
}

func (e *entry) isForSameDay(other entry) bool {
	return e.year == other.year && e.month == other.month && e.day == other.day
}

func (e *entry) happenedBefore(other entry) bool {
	d := time.Date(
		e.year, time.Month(e.month), e.day, e.hour, e.minute, 0, 0, time.UTC)
	o := time.Date(
		other.year, time.Month(other.month), other.day, other.hour, other.minute, 0, 0, time.UTC)
	return d.Before(o)
}

type scheduler struct {
	entries []entry
	guards  map[int]guard
}

func (s scheduler) calculate() int {
	asleepAt, guardID := 0, 0

	for _, e := range s.entries {
		if e.action == "wakes_up" {
			_, ok := s.guards[guardID]
			if !ok {
				s.guards[guardID] = newGuard(guardID)
			}
			s.guards[guardID].sleptBetween(asleepAt, e.minute)
		} else if e.action == "begin_shift" {
			guardID = e.guardID
		} else {
			asleepAt = e.minute
		}
	}

	maxMinutes := 0
	var guardWithMostMinutes guard
	for _, g := range s.guards {
		if m := g.countMinutesSleeping(); m > maxMinutes {
			maxMinutes = m
			guardWithMostMinutes = g
		}
	}

	return guardWithMostMinutes.id * guardWithMostMinutes.mostLikelyMinuteToBeAsleep()
}

func newScheduler(entries []entry) scheduler {
	return scheduler{entries, map[int]guard{}}
}

func main() {
	entries := readEntries()
	sort.Slice(entries, func(i, j int) bool { return entries[i].happenedBefore(entries[j]) })

	s := newScheduler(entries)
	fmt.Println(s.calculate()) // 2657*33=87681
}
