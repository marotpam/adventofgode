package _2018

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type step struct {
	id                string
	blockedBy, blocks []*step
	visited           bool
	duration          uint8
}

func newStep(id string) *step {
	return &step{
		id:       id,
		duration: 60 + id[0] - 'A' + 1,
		visited:  false,
	}
}

func newStepWithDuration(id string, duration uint8) *step {
	return &step{id: id, duration: duration, visited: false}
}

func (s *step) blocking(other *step) {
	s.blocks = append(s.blocks, other)
	sort.Slice(s.blocks, func(i, j int) bool {
		return s.blocks[i].id < s.blocks[j].id
	})

	other.addBlocker(s)
}

func (s *step) addBlocker(other *step) {
	s.blockedBy = append(s.blockedBy, other)
	sort.Slice(s.blockedBy, func(i, j int) bool {
		return s.blockedBy[i].id < s.blockedBy[j].id
	})
}

func (s *step) isBlocked() bool {
	return len(s.blockedBy) > 0
}

func (s *step) done() {
	for _, b := range s.blocks {
		b.noLongerBlockedBy(s)
	}
	s.blocks = nil
	s.visited = true
}

func (s *step) noLongerBlockedBy(blocker *step) {
	for i, b := range s.blockedBy {
		if b.id == blocker.id {
			s.blockedBy = append(s.blockedBy[:i], s.blockedBy[i+1:]...)
			return
		}
	}
}

type requirement struct {
	blocker, blocked string
}

type instructions struct {
	steps          map[string]*step
	pendingStepIds []string
}

func newInstructions(rs []requirement) *instructions {
	steps := map[string]*step{}
	pendingStepIds := []string{}

	for _, r := range rs {
		blocker, ok := steps[r.blocker]
		if !ok {
			blocker = newStep(r.blocker)
			steps[r.blocker] = blocker
			pendingStepIds = append(pendingStepIds, r.blocker)
		}
		blocked, ok := steps[r.blocked]
		if !ok {
			blocked = newStep(r.blocked)
			steps[r.blocked] = blocked
			pendingStepIds = append(pendingStepIds, r.blocked)
		}

		blocker.blocking(blocked)
	}

	sort.Strings(pendingStepIds)

	return &instructions{steps, pendingStepIds}
}

func (i *instructions) availableSteps() []*step {
	a := []*step{}

	for _, s := range i.steps {
		if !s.isBlocked() && !s.visited {
			a = append(a, s)
		}
	}

	return a
}

type worker struct {
	id        int
	busyUntil uint8
	step      *step
}

func (w *worker) workOn(s *step) {
	s.visited = true
	w.step = s
	w.busyUntil = s.duration
}

func (w *worker) isBusy() bool {
	return w.step != nil
}

func (w *worker) work() string {
	if w.busyUntil > 0 {
		w.busyUntil--
		if w.busyUntil == 0 {
			stepId := w.step.id
			w.step.done()
			w.step = nil
			return stepId
		}
	}
	return ""
}

type scheduler struct {
	workers map[int]*worker
}

func (sch *scheduler) idleWorkerIds() []int {
	iws := []int{}

	for id, w := range sch.workers {
		if !w.isBusy() {
			iws = append(iws, id)
		}
	}

	return iws
}

func (sch *scheduler) workOn(i *instructions) ([]string, int) {
	t := 0
	r := []string{}

	for len(r) < len(i.pendingStepIds) {
		idleWorkerIds := sch.idleWorkerIds()
		availableSteps := i.availableSteps()
		for w := 0; w < len(idleWorkerIds) && w < len(availableSteps); w++ {
			idleWorkerID := idleWorkerIds[w]
			sch.workers[idleWorkerID].workOn(availableSteps[w])
		}

		finishedSteps := sch.tick()
		if len(finishedSteps) > 0 {
			r = append(r, finishedSteps...)
		}
		t++
	}

	return r, t
}

func (sch *scheduler) tick() []string {
	finishedStepIds := []string{}
	for _, w := range sch.workers {
		if finishedStepId := w.work(); finishedStepId != "" {
			finishedStepIds = append(finishedStepIds, finishedStepId)
		}
	}
	sort.Strings(finishedStepIds)

	return finishedStepIds
}

func (sch *scheduler) isDone() bool {
	return len(sch.idleWorkerIds()) == len(sch.workers)
}

func newScheduler(workersCount int) *scheduler {
	ws := map[int]*worker{}

	for i := 0; i < workersCount; i++ {
		ws[i] = &worker{id: i}
	}

	return &scheduler{ws}
}

func readRequirements() []requirement {
	rs := []requirement{}

	fileHandle, _ := os.Open("input.txt")
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		var blocker, blocked string
		fmt.Sscanf(fileScanner.Text(), `Step %s must be finished before step %s can begin.`, &blocker, &blocked)
		rs = append(rs, requirement{blocker, blocked})
	}

	return rs
}

func main() {
	i := newInstructions(readRequirements())

	s := newScheduler(5)
	_, d := s.workOn(i)

	fmt.Println(d) // 877
}
